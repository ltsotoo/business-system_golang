package model

import "gorm.io/gorm"

// 合同 Model
type Contract struct {
	gorm.Model
	No                    string `gorm:"type:varchar(20);comment:合同编号;not null" json:"no"`
	AreaId                int    `gorm:"type:int;comment:所属区域ID;not null" json:"areaId"`
	EmployeeId            int    `gorm:"type:int;comment:业务员ID;not null" json:"employeeId"`
	IsEntryCustomer       bool   `gorm:"type:boolean;comment:客户是否录入;not null" json:"isEntryCustomer"`
	CustomerId            int    `gorm:"type:int;comment:客户ID" json:"customerId"`
	EstimatedDeliveryDate string `gorm:"type:varchar(20);comment:合同交货日期;not null" json:"estimatedDeliveryDate"`
	EndDeliveryDate       string `gorm:"type:varchar(20);comment:实际交货日期" json:"endDeliveryDate"`
	InvoiceType           int    `gorm:"type:int;comment:开票类型;not null" json:"invoiceType"`
	InvoiceContent        string `gorm:"type:varchar(20);comment:开票内容" json:"invoiceContent"`
	IsSpecial             bool   `gorm:"type:boolean;comment:是否是特殊合同;not null" json:"isSpecial"`
	Remarks               string `gorm:"type:varchar(20);comment:备注" json:"remarks"`
	Status                int    `gorm:"type:int;comment:状态;not null" json:"status"`

	Tasks           []Product `gorm:"many2many:contract_task" json:"tasks"`
	NoEntryCustomer Customer  `gorm:"-"`
}
