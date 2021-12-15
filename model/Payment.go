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
	InvoiceUID           string  `gorm:"type:varchar(32);comment:发票UID;default:(-)" json:"invoiceUID"`
	TaskUID              string  `gorm:"type:varchar(32);comment:任务UID;default:(-)" json:"taskUID"`
	PaymentDate          XDate   `gorm:"type:date;comment:回款日期" json:"paymentDate"`
	EmployeeUID          string  `gorm:"type:varchar(32);comment:录入人员ID;default:(-)" json:"employeeUID"`
	Money                float64 `gorm:"type:decimal(20,6);comment:回款金额(元)" json:"money"`
	TheoreticalPushMoney float64 `gorm:"type:decimal(20,6);comment:理论提成(元)" json:"theoreticalPushMoney"`
	Fine                 float64 `gorm:"type:decimal(20,6);comment:回款延迟扣除(元)" json:"fine"`
	PushMoney            float64 `gorm:"type:decimal(20,6);comment:实际提成(元)" json:"pushMoney"`
	BusinessMoney        float64 `gorm:"type:decimal(20,6);comment:业务费用(元)" json:"businessMoney"`
	Parities             float64 `gorm:"type:decimal(20,6);comment:汇率" json:"parities"`
}

func InsertPayment(payment *Payment) (code int) {

	var contract Contract

	db.First(&contract, "uid = ?", payment.ContractUID)

	if contract.UID == "" {
		return msg.ERROR
	}

	if contract.IsSpecial {
		db.Find(&contract.Tasks, "contract_uid = ?", contract.UID)
	} else {
		db.Preload("Product.Type").Find(&contract.Tasks, "contract_uid = ?", contract.UID)
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

		//任务回款金额更新
		if payment.TaskUID != "" {
			if tErr := tdb.Exec("UPDATE task SET payment_total_price = payment_total_price + ? WHERE uid = ?", payment.Money, payment.TaskUID).Error; tErr != nil {
				return tErr
			}
		}

		tempPushMoney := payment.PushMoney * 0.5
		if !contract.IsPreDeposit {
			//办事处业务费，提成，提升（预存款合同）
			if tErr := tdb.Exec("UPDATE office SET money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE uid = ?", tempPushMoney, tempPushMoney, payment.BusinessMoney, contract.OfficeUID).Error; tErr != nil {
				return tErr
			}
		} else {
			//办事处业务费，提成，任务量提升
			if tErr := tdb.Exec("UPDATE office SET target_load = target_load + ?, money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE uid = ?", payment.Money, tempPushMoney, tempPushMoney, payment.BusinessMoney, contract.OfficeUID).Error; tErr != nil {
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

func SelectPaymentsByContractUID(contractUID string) (payments []Payment, code int) {
	err = db.Where("contract_uid = ?", contractUID).Find(&payments).Error
	if err != nil {
		return payments, msg.ERROR
	}
	return payments, msg.SUCCESS
}

func SelectContractAndPayments(uid string) (contract Contract, code int) {
	err = db.Preload("ContractUnit").
		Preload("Employee").Preload("Employee.Office").
		Preload("Customer").Preload("Customer.Company").
		Preload("Payments").
		First(&contract, "uid = ?", uid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return contract, msg.ERROR_CONTRACT_NOT_EXIST
		} else {
			return contract, msg.ERROR_CONTRACT_SELECT
		}
	}
	return contract, msg.SUCCESS
}
