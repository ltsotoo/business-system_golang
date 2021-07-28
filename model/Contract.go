package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 合同 Model
type Contract struct {
	gorm.Model
	No                    string `gorm:"type:varchar(20);comment:合同编号;not null" json:"no"`
	AreaId                uint   `gorm:"type:int;comment:所属区域ID;not null" json:"areaId"`
	EmployeeId            uint   `gorm:"type:int;comment:业务员ID;not null" json:"employeeId"`
	IsEntryCustomer       bool   `gorm:"type:boolean;comment:客户是否录入;not null" json:"isEntryCustomer"`
	CustomerId            int    `gorm:"type:int;comment:客户ID" json:"customerId"`
	EstimatedDeliveryDate string `gorm:"type:varchar(20);comment:合同交货日期;not null" json:"estimatedDeliveryDate"`
	EndDeliveryDate       string `gorm:"type:varchar(20);comment:实际交货日期" json:"endDeliveryDate"`
	InvoiceType           int    `gorm:"type:int;comment:开票类型;not null" json:"invoiceType"`
	InvoiceContent        string `gorm:"type:varchar(20);comment:开票内容" json:"invoiceContent"`
	IsSpecial             bool   `gorm:"type:boolean;comment:是否是特殊合同;not null" json:"isSpecial"`
	Remarks               string `gorm:"type:varchar(20);comment:备注" json:"remarks"`
	Status                int    `gorm:"type:int;comment:状态;not null" json:"status"`

	Tasks           []Task   `gorm:"foreignKey:ContractId" json:"tasks"`
	NoEntryCustomer Customer `gorm:"-"`
}

func CreateContract(contract *Contract) (code int) {
	err = db.Create(&contract).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteContract(id int) (code int) {
	var contract Contract
	contract.ID = uint(id)
	err = db.Select("Tasks").Delete(&contract).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateContract(contract *Contract) (code int) {
	var maps = make(map[string]interface{})
	maps["Remarks"] = contract.Remarks
	err = db.Model(&contract).Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectContract(id int) (contract Contract, code int) {
	contract.ID = uint(id)
	err = db.Model(&contract).Association("Tasks").Find(&contract.Tasks)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return contract, msg.ERROR_CONTRACT_NOT_EXIST
		} else {
			return contract, msg.ERROR
		}
	}
	return contract, msg.SUCCESS
}

func SelectContracts(pageSize int, pageNo int) (contracts []Contract, code int, total int64) {
	err = db.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&contracts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR, 0
	}
	db.Model(&contracts).Count(&total)
	return contracts, msg.SUCCESS, total
}
