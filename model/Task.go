package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 合同任务 Model
type Task struct {
	gorm.Model
	ContractID      uint   `gorm:"type:int;comment:合同ID;not null" json:"contractID"`
	ProductID       uint   `gorm:"type:int;comment:产品ID;not null" json:"productID"`
	Number          int    `gorm:"type:int;comment:数量;not null" json:"number"`
	Unit            string `gorm:"type:varchar(20);comment:单位;not null" json:"unit"`
	Status          int    `gorm:"type:int;comment:状态;not null" json:"status"`
	TechnicianManID uint   `gorm:"type:int;comment:技术负责人ID" json:"technicianManID"`
	PurchaseManID   uint   `gorm:"type:int;comment:采购负责人ID" json:"purchaseManID"`
	InventoryManID  uint   `gorm:"type:int;comment:库存负责人ID" json:"inventoryManID"`
	ShipmentManID   uint   `gorm:"type:int;comment:发货人员ID" json:"shipmentManID"`
}

func DeleteTask(id int) (code int) {
	err = db.Where("id = ?", id).Delete(&Task{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectTaskByContractID(contractID int) (tasks []Task, code int) {
	err = db.Where("contract_id = ?", contractID).Find(&tasks).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return tasks, msg.SUCCESS
}
