package model

import (
	"gorm.io/gorm"
)

// 合同任务 Model
type Task struct {
	gorm.Model
	ContractId      int    `gorm:"type:int;comment:合同ID;not null" json:"contractId"`
	ProductId       int    `gorm:"type:int;comment:产品ID;not null" json:"productId"`
	Number          int    `gorm:"type:int;comment:数量;not null" json:"number"`
	Unit            string `gorm:"type:varchar(20);comment:单位;not null" json:"unit"`
	Status          int    `gorm:"type:int;comment:状态;not null" json:"status"`
	TechnicianManId int    `gorm:"type:int;comment:技术负责人ID" json:"technicianManId"`
	PurchaseManId   int    `gorm:"type:int;comment:采购负责人ID" json:"purchaseManId"`
	InventoryManId  int    `gorm:"type:int;comment:库存负责人ID" json:"inventoryManId"`
	ShipmentManId   int    `gorm:"type:int;comment:发货人员ID" json:"shipmentManId"`
}
