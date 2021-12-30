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

	Task Task `gorm:"foreignKey:TaskUID;references:UID" json:"task"`
}

func InsertPayment(payment *Payment) (code int) {

	var contract Contract

	db.First(&contract, "uid = ?", payment.ContractUID)

	if contract.UID == "" {
		return msg.ERROR
	}

	db.Preload("Product.Type").Find(&contract.Tasks, "UID = ?", payment.TaskUID)
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

		//合同回款金额更新
		if tErr := tdb.Exec("UPDATE contract SET payment_total_amount = payment_total_amount + ? WHERE uid = ?", payment.Money, payment.ContractUID).Error; tErr != nil {
			return tErr
		}

		//产品任务回款金额更新
		if tErr := tdb.Exec("UPDATE task SET payment_total_price = payment_total_price + ? WHERE uid = ?", payment.Money, payment.TaskUID).Error; tErr != nil {
			return tErr
		}

		tempPushMoney1 := payment.PushMoney * 0.5
		tempPushMoney2 := payment.PushMoney - tempPushMoney1
		if contract.IsPreDeposit {
			//办事处业务费，提成 UP（预存款合同）
			if tErr := tdb.Exec("UPDATE office SET money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE uid = ?", tempPushMoney1, tempPushMoney2, payment.BusinessMoney, contract.OfficeUID).Error; tErr != nil {
				return tErr
			}
		} else {
			//办事处业务费，提成，任务量 UP
			if tErr := tdb.Exec("UPDATE office SET target_load = target_load + ?, money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE uid = ?", payment.Money, tempPushMoney1, tempPushMoney2, payment.BusinessMoney, contract.OfficeUID).Error; tErr != nil {
				return tErr
			}
		}
		return nil
	})
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdatePayment(payment *Payment) (code int) {
	//TODO
	return msg.SUCCESS
}

func SelectPayments(contractUID string) (payments []Payment, code int) {
	err = db.Preload("Task.Product").Where("contract_uid = ?", contractUID).Find(&payments).Error
	if err != nil {
		return payments, msg.ERROR
	}
	return payments, msg.SUCCESS
}
