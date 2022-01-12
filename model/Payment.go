package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

type Payment struct {
	BaseModel
	UID                  string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractUID          string  `gorm:"type:varchar(32);comment:合同UID;default:(-)" json:"contractUID"`
	TaskUID              string  `gorm:"type:varchar(32);comment:任务UID;default:(-)" json:"taskUID"`
	PaymentDate          XDate   `gorm:"type:date;comment:回款日期" json:"paymentDate"`
	EmployeeUID          string  `gorm:"type:varchar(32);comment:录入人员ID;default:(-)" json:"employeeUID"`
	Money                float64 `gorm:"type:decimal(20,6);comment:回款金额(人民币)" json:"money"`
	TheoreticalPushMoney float64 `gorm:"type:decimal(20,6);comment:理论提成(元)" json:"theoreticalPushMoney"`
	Fine                 float64 `gorm:"type:decimal(20,6);comment:回款延迟扣除(元)" json:"fine"`
	PushMoney            float64 `gorm:"type:decimal(20,6);comment:实际提成(元)" json:"pushMoney"`
	BusinessMoney        float64 `gorm:"type:decimal(20,6);comment:业务费用(元)" json:"businessMoney"`

	Task     Task     `gorm:"foreignKey:TaskUID;references:UID" json:"task"`
	Employee Employee `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
}

func InsertPayment(payment *Payment) (code int) {

	var contract Contract
	db.First(&contract, "uid = ?", payment.ContractUID)
	if contract.UID == "" {
		return msg.ERROR
	}

	//是否预存款合同
	if contract.IsPreDeposit {
		payment.UID = uid.Generate()
		err = db.Transaction(func(tdb *gorm.DB) error {
			//添加记录
			if tErr := tdb.Create(&payment).Error; tErr != nil {
				return tErr
			}
			//修改合同可用预存款
			if tErr := tdb.Exec("UPDATE contract SET payment_total_amount = payment_total_amount + ?,pre_deposit = pre_deposit + ? WHERE uid = ?", payment.Money, payment.Money, payment.ContractUID).Error; tErr != nil {
				return tErr
			}
			//办事处任务量 UP
			if tErr := tdb.Exec("UPDATE office SET target_load = target_load + ? WHERE uid = ?", payment.Money, contract.OfficeUID).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	} else {
		db.Preload("Product.Type").Find(&contract.Tasks, "UID = ? AND contract_uid = ?", payment.TaskUID, payment.ContractUID)
		if len(contract.Tasks) != 1 {
			return msg.ERROR
		}

		payment.UID = uid.Generate()
		err = db.Transaction(func(tdb *gorm.DB) error {

			//计算提成并创建提成记录
			payment.TheoreticalPushMoney, payment.Fine, payment.PushMoney, payment.BusinessMoney = calculate(&contract, payment)
			if tErr := tdb.Create(&payment).Error; tErr != nil {
				return tErr
			}

			//合同回款金额更新，总付款金额更新
			if tErr := tdb.Exec("UPDATE contract SET payment_total_amount = payment_total_amount + ? WHERE uid = ?", payment.Money, payment.ContractUID).Error; tErr != nil {
				return tErr
			}

			//产品任务回款金额更新
			if tErr := tdb.Exec("UPDATE task SET payment_total_price = payment_total_price + ? WHERE uid = ?", payment.Money, payment.TaskUID).Error; tErr != nil {
				return tErr
			}

			tempPushMoney1 := payment.PushMoney * 0.5
			tempPushMoney2 := payment.PushMoney - tempPushMoney1

			//办事处业务费，提成，任务量 UP
			if contract.Tasks[0].Product.Type.IsTaskLoad {
				if tErr := tdb.Exec("UPDATE office SET target_load = target_load + ?, money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE uid = ?", payment.Money, tempPushMoney1, tempPushMoney2, payment.BusinessMoney, contract.OfficeUID).Error; tErr != nil {
					return tErr
				}
			} else {
				if tErr := tdb.Exec("UPDATE office SET money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE uid = ?", tempPushMoney1, tempPushMoney2, payment.BusinessMoney, contract.OfficeUID).Error; tErr != nil {
					return tErr
				}
			}

			return nil
		})
	}

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdatePayment(payment *Payment) (code int) {
	//查询数据库记录
	var oldPayment Payment
	db.First(&oldPayment, "uid = ?", payment.UID)
	if oldPayment.ID == 0 {
		return msg.ERROR
	}
	//查询合同
	var contract Contract
	db.First(&contract, "uid = ?", payment.ContractUID)
	if contract.ID == 0 {
		return msg.ERROR
	}

	if contract.IsPreDeposit {
		err = db.Transaction(func(tdb *gorm.DB) error {
			//合同回款金额更新，总付款金额更新
			tempPaymentTotalAmount := payment.Money - oldPayment.Money
			if tErr := tdb.Exec("UPDATE contract SET payment_total_amount = payment_total_amount + ?,pre_deposit = pre_deposit + ? WHERE uid = ?", tempPaymentTotalAmount, tempPaymentTotalAmount, payment.ContractUID).Error; tErr != nil {
				return tErr
			}
			//办事处任务量跟新
			if tErr := tdb.Exec("UPDATE office SET target_load = target_load + ? WHERE uid = ?", tempPaymentTotalAmount, contract.OfficeUID).Error; tErr != nil {
				return tErr
			}
			//更新记录
			var maps = make(map[string]interface{})
			maps["employee_uid"] = payment.EmployeeUID
			maps["payment_date"] = payment.PaymentDate.Time
			maps["money"] = payment.Money
			if tErr := tdb.Model(&Payment{}).Where("uid = ?", payment.UID).Updates(maps).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	} else {
		//计算提成
		db.Preload("Product.Type").Find(&contract.Tasks, "UID = ? AND contract_uid = ?", payment.TaskUID, payment.ContractUID)
		if len(contract.Tasks) != 1 {
			return msg.ERROR
		}
		payment.TheoreticalPushMoney, payment.Fine, payment.PushMoney, payment.BusinessMoney = calculate(&contract, payment)
		//更新数据库
		err = db.Transaction(func(tdb *gorm.DB) error {
			tempMoney := payment.Money - oldPayment.Money
			//合同回款金额更新
			if tErr := tdb.Exec("UPDATE contract SET payment_total_amount = payment_total_amount + ? WHERE uid = ?", tempMoney, payment.ContractUID).Error; tErr != nil {
				return tErr
			}
			//产品任务回款金额更新
			if tErr := tdb.Exec("UPDATE task SET payment_total_price = payment_total_price + ? WHERE uid = ?", tempMoney, payment.TaskUID).Error; tErr != nil {
				return tErr
			}
			tempOldPushMoney1 := oldPayment.PushMoney * 0.5
			tempOldPushMoney2 := oldPayment.PushMoney - tempOldPushMoney1
			tempPushMoney1 := payment.PushMoney*0.5 - tempOldPushMoney1
			tempPushMoney2 := payment.PushMoney - tempPushMoney1 - tempOldPushMoney2
			tempBusinessMoney := payment.BusinessMoney - oldPayment.BusinessMoney

			//办事处业务费，提成，任务量 UP
			if contract.Tasks[0].Product.Type.IsTaskLoad {
				if tErr := tdb.Exec("UPDATE office SET target_load = target_load + ?, money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE uid = ?", tempMoney, tempPushMoney1, tempPushMoney2, tempBusinessMoney, contract.OfficeUID).Error; tErr != nil {
					return tErr
				}
			} else {
				if tErr := tdb.Exec("UPDATE office SET money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE uid = ?", tempPushMoney1, tempPushMoney2, tempBusinessMoney, contract.OfficeUID).Error; tErr != nil {
					return tErr
				}
			}

			//更新回款记录
			var maps = make(map[string]interface{})
			maps["employee_uid"] = payment.EmployeeUID
			maps["payment_date"] = payment.PaymentDate.Time
			maps["money"] = payment.Money
			maps["theoretical_push_money"] = payment.TheoreticalPushMoney
			maps["fine"] = payment.Fine
			maps["push_money"] = payment.PushMoney
			maps["business_money"] = payment.BusinessMoney
			if tErr := tdb.Model(&Payment{}).Where("uid = ?", payment.UID).Updates(maps).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	}

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectPayments(contractUID string) (payments []Payment, code int) {
	err = db.Preload("Task.Product").Preload("Employee").Where("contract_uid = ? AND task_uid is not null", contractUID).Find(&payments).Error
	if err != nil {
		return payments, msg.ERROR
	}
	return payments, msg.SUCCESS
}

func SelectPrePayments(contractUID string) (payments []Payment, code int) {
	err = db.Preload("Employee").Where("contract_uid = ? AND task_uid is null", contractUID).Find(&payments).Error
	if err != nil {
		return payments, msg.ERROR
	}
	return payments, msg.SUCCESS
}
