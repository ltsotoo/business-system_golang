package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"
)

type Invoice struct {
	BaseModel
	UID         string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractUID string  `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	EmployeeUID string  `gorm:"type:varchar(32);comment:添加员工UID;default:(-)" json:"employeeUID"`
	Code        string  `gorm:"type:varchar(100);comment:发票号" json:"code"`
	Money       float64 `gorm:"type:decimal(20,6);comment:总金额" json:"money"`
	IsDelete    bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	Contract Contract `gorm:"foreignKey:ContractUID;references:UID" json:"contract"`
	Employee Employee `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
}

func InsertInvoice(invoice *Invoice) (code int) {
	invoice.UID = uid.Generate()
	err = db.Create(&invoice).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteInvoice(uid string) (code int) {
	err = db.Model(&Invoice{}).Where("uid = ?", uid).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateInvoice(invoice *Invoice) (code int) {
	var maps = make(map[string]interface{})
	maps["code"] = invoice.Code
	maps["money"] = invoice.Money

	err = db.Model(&Invoice{}).Where("uid = ?", invoice.UID).Updates(maps).Error

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectInvoices(contractUID string) (invoices []Invoice, code int) {
	var maps = make(map[string]interface{})
	maps["is_delete"] = false
	maps["contract_uid"] = contractUID

	err = db.Preload("Contract.Employee").Where(maps).Find(&invoices).Error
	if err != nil {
		return invoices, msg.ERROR
	}
	return invoices, msg.SUCCESS
}
