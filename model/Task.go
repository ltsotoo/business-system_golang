package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 合同任务 Model
type Task struct {
	gorm.Model
	UID              string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractNo       string `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	ProductUID       string `gorm:"type:varchar(32);comment:产品ID" json:"productUID"`
	Number           int    `gorm:"type:int;comment:数量" json:"number"`
	Unit             string `gorm:"type:varchar(9);comment:单位" json:"unit"`
	Price            int    `gorm:"type:int;comment:单价(元)" json:"price"`
	TotalPrice       int    `gorm:"type:int;comment:总价(元)" json:"totalPrice"`
	Status           int    `gorm:"type:int;comment:状态" json:"status"`
	TechnicianManUID string `gorm:"type:varchar(32);comment:技术负责人ID" json:"technicianManUID"`
	PurchaseManUID   string `gorm:"type:varchar(32);comment:采购负责人ID" json:"purchaseManUID"`
	InventoryManUID  string `gorm:"type:varchar(32);comment:库存负责人ID" json:"inventoryManUID"`
	ShipmentManUID   string `gorm:"type:varchar(32);comment:发货人员ID" json:"shipmentManUID"`

	Contract      Contract `gorm:"foreignKey:ContractNo;references:No" json:"contract"`
	Product       Product  `gorm:"foreignKey:ProductUID;references:UID" json:"product"`
	TechnicianMan Employee `gorm:"foreignKey:TechnicianManUID;references:UID" json:"technicianMan"`
	PurchaseMan   Employee `gorm:"foreignKey:PurchaseManUID;references:UID" json:"purchaseMan"`
	InventoryMan  Employee `gorm:"foreignKey:InventoryManUID;references:UID" json:"inventoryMan"`
	ShipmentMan   Employee `gorm:"foreignKey:ShipmentManUID;references:UID" json:"shipmentMan"`
}

func DeleteTask(uid string) (code int) {
	err = db.Delete(&Task{}, "uid = ?", uid).Error
	if err != nil {
		return msg.ERROR_TASK_DELETE
	}
	return msg.SUCCESS
}

func SelectTasks(task *Task) (tasks []Task, code int) {
	err = db.Preload("Contract").Preload("Product").Preload("TechnicianMan").Preload("PurchaseMan").Preload("InventoryMan").Preload("ShipmentMan").
		Where(&task).Find(&tasks).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return tasks, msg.SUCCESS
}
