package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 合同任务 Model
type Task struct {
	gorm.Model
	ContractID      uint   `gorm:"type:int;comment:合同ID;default:(-)" json:"contractID"`
	ProductID       uint   `gorm:"type:int;comment:产品ID;default:(-)" json:"productID"`
	Number          int    `gorm:"type:int;comment:数量" json:"number"`
	Unit            string `gorm:"type:varchar(20);comment:单位" json:"unit"`
	Price           int    `gorm:"type:int;comment:单价(元)" json:"rrice"`
	TotalPrice      int    `gorm:"type:int;comment:总价(元)" json:"totalPrice"`
	Status          int    `gorm:"type:int;comment:状态" json:"status"`
	TechnicianManID uint   `gorm:"type:int;comment:技术负责人ID;default:(-)" json:"technicianManID"`
	PurchaseManID   uint   `gorm:"type:int;comment:采购负责人ID;default:(-)" json:"purchaseManID"`
	InventoryManID  uint   `gorm:"type:int;comment:库存负责人ID;default:(-)" json:"inventoryManID"`
	ShipmentManID   uint   `gorm:"type:int;comment:发货人员ID;default:(-)" json:"shipmentManID"`

	Contract      Contract `gorm:"foreignKey:ContractID" json:"contract"`
	Product       Product  `gorm:"foreignKey:ProductID" json:"product"`
	TechnicianMan Employee `gorm:"foreignKey:TechnicianManID" json:"technicianMan"`
	PurchaseMan   Employee `gorm:"foreignKey:PurchaseManID" json:"purchaseMan"`
	InventoryMan  Employee `gorm:"foreignKey:InventoryManID" json:"inventoryMan"`
	ShipmentMan   Employee `gorm:"foreignKey:ShipmentManID" json:"shipmentMan"`
}

func DeleteTask(id int) (code int) {
	err = db.Where("id = ?", id).Delete(&Task{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectTaskByContractID(contractID int) (tasks []Task, code int) {
	err = db.Preload("Contract").Preload("Product").Preload("TechnicianMan").Preload("PurchaseMan").Preload("InventoryMan").Preload("ShipmentMan").Where("contract_id = ?", contractID).Find(&tasks).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return tasks, msg.SUCCESS
}
