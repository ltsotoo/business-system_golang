package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 合同 Model
type Contract struct {
	gorm.Model
	No                    string `gorm:"type:varchar(20);comment:合同编号;not null" json:"no"`
	AreaID                uint   `gorm:"type:int;comment:所属区域ID;not null" json:"areaID"`
	EmployeeID            uint   `gorm:"type:int;comment:业务员ID;not null" json:"employeeID"`
	IsEntryCustomer       bool   `gorm:"type:boolean;comment:客户是否录入;not null" json:"isEntryCustomer"`
	CustomerID            uint   `gorm:"type:int;comment:客户ID" json:"customerID"`
	ContractDate          string `gorm:"type:varchar(20);comment:签订日期;not null" json:"contractDate"`
	ContractUnitID        uint   `gorm:"type:int;comment:签订单位;not null" json:"contractUnitID"`
	EstimatedDeliveryDate string `gorm:"type:varchar(20);comment:合同交货日期;not null" json:"estimatedDeliveryDate"`
	EndDeliveryDate       string `gorm:"type:varchar(20);comment:实际交货日期" json:"endDeliveryDate"`
	InvoiceType           int    `gorm:"type:int;comment:开票类型;not null" json:"invoiceType"`
	InvoiceContent        string `gorm:"type:varchar(20);comment:开票内容" json:"invoiceContent"`
	IsSpecial             bool   `gorm:"type:boolean;comment:特殊合同?;not null" json:"isSpecial"`
	TotalAmount           int    `gorm:"type:int;comment:总金额(元);not null" json:"totalAmount"`
	Remarks               string `gorm:"type:varchar(20);comment:备注" json:"remarks"`
	Status                int    `gorm:"type:int;comment:状态;not null" json:"status"`

	Tasks    []Task   `gorm:"foreignKey:ContractID" json:"tasks"`
	Customer Customer `gorm:"foreignKey:CustomerID" json:"customer"`
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
	db.Preload("Tasks").Preload("Customer").First(&contract, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return contract, msg.ERROR_CONTRACT_NOT_EXIST
		} else {
			return contract, msg.ERROR
		}
	}
	return contract, msg.SUCCESS
}

func SelectContracts(pageSize int, pageNo int, contractQuery ContractQuery) (contracts []Contract, code int, total int64) {
	err = db.Limit(pageSize).Offset((pageNo-1)*pageSize).Where("no LIKE ?", "%"+contractQuery.No+"%").Find(&contracts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR, 0
	}
	db.Model(&contracts).Count(&total)
	return contracts, msg.SUCCESS, total
}
