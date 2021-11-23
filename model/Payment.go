package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	BaseModel
	UID         string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractUID string  `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	EmployeeUID string  `gorm:"type:varchar(32);comment:技术负责人ID" json:"employeeUID"`
	Money       float64 `gorm:"type:decimal(20,6);comment:回款金额(元)" json:"money"`
	Remarks     string  `gorm:"type:varchar(600);comment:备注" json:"remarks"`
}

type PaymentQuery struct {
	UID                string    `json:"UID"`
	EndPaymentDate     time.Time `json:"endPaymentDate"`
	PaymentTotalAmount float64   `json:"paymentTotalAmount"`
}

func InsertPayment(payment *Payment) (code int) {
	payment.UID = uid.Generate()
	// err = db.Create(&payment).Error
	err = db.Transaction(func(tdb *gorm.DB) error {
		if tErr := tdb.Create(&payment).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Exec("UPDATE contract SET payment_total_amount = payment_total_amount + ? WHERE uid = ?", payment.Money, payment.ContractUID).Error; tErr != nil {
			return tErr
		}
		return nil
	})
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeletePayment(uid string) (code int) {
	err = db.Where("uid = ?", uid).Delete(&Payment{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdatePayment(payment *Payment) (code int) {
	var maps = make(map[string]interface{})
	maps["money"] = payment.Money
	maps["remarks"] = payment.Remarks
	err = db.Model(&Payment{}).Where("uid = ?", payment.UID).Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
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
	err = db.Preload("Area").Preload("ContractUnit").
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
