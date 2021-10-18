package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

type Payment struct {
	BaseModel
	UID         string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractUID string `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	EmployeeUID string `gorm:"type:varchar(32);comment:技术负责人ID" json:"employeeUID"`
	Money       int    `gorm:"type:int;comment:回款金额(元)" json:"money"`
	Remarks     string `gorm:"type:varchar(499);comment:备注" json:"remarks"`
}

func InsertPayment(payment *Payment) (code int) {
	payment.UID = uid.Generate()
	err = db.Create(&payment).Error
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
