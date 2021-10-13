package model

import (
	"business-system_golang/utils/magic"
	"business-system_golang/utils/msg"
)

// 合同任务 Model
type Task struct {
	BaseModel
	UID              string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractUID      string `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	ProductUID       string `gorm:"type:varchar(32);comment:产品ID" json:"productUID"`
	Number           int    `gorm:"type:int;comment:数量" json:"number"`
	Unit             string `gorm:"type:varchar(9);comment:单位" json:"unit"`
	Price            int    `gorm:"type:int;comment:单价(元)" json:"price"`
	TotalPrice       int    `gorm:"type:int;comment:总价(元)" json:"totalPrice"`
	Status           int    `gorm:"type:int;comment:状态" json:"status"`
	TechnicianManUID string `gorm:"type:varchar(32);comment:技术负责人ID;default:(-)" json:"technicianManUID"`
	PurchaseManUID   string `gorm:"type:varchar(32);comment:采购负责人ID;default:(-)" json:"purchaseManUID"`
	InventoryManUID  string `gorm:"type:varchar(32);comment:库存负责人ID;default:(-)" json:"inventoryManUID"`
	ShipmentManUID   string `gorm:"type:varchar(32);comment:发货人员ID;default:(-)" json:"shipmentManUID"`
	Remarks          string `gorm:"type:varchar(200);comment:备注" json:"remarks"`
	NextRemarks      string `gorm:"type:varchar(200);comment:流程备注" json:"nextRemarks"`

	Contract      Contract `gorm:"foreignKey:ContractUID;references:UID" json:"contract"`
	Product       Product  `gorm:"foreignKey:ProductUID;references:UID" json:"product"`
	TechnicianMan Employee `gorm:"foreignKey:TechnicianManUID;references:UID" json:"technicianMan"`
	PurchaseMan   Employee `gorm:"foreignKey:PurchaseManUID;references:UID" json:"purchaseMan"`
	InventoryMan  Employee `gorm:"foreignKey:InventoryManUID;references:UID" json:"inventoryMan"`
	ShipmentMan   Employee `gorm:"foreignKey:ShipmentManUID;references:UID" json:"shipmentMan"`
}

type TaskFlowQuery struct {
	UID              string `json:"UID"`
	Status           int    `json:"status"`
	TechnicianManUID string `json:"technicianManUID"`
	PurchaseManUID   string `json:"purchaseManUID"`
	InventoryManUID  string `json:"inventoryManUID"`
	ShipmentManUID   string `json:"shipmentManUID"`
}

func DeleteTask(uid string) (code int) {
	err = db.Delete(&Task{}, "uid = ?", uid).Error
	if err != nil {
		return msg.ERROR_TASK_DELETE
	}
	return msg.SUCCESS
}

func SelectTask(uid string) (task Task, code int) {
	err = db.Where("uid = ?", uid).First(&task).Error
	if err != nil {
		return task, msg.ERROR
	}
	return task, msg.SUCCESS
}

func SelectTasks(task *Task) (tasks []Task, code int) {
	err = db.Preload("Contract").Preload("Product").Preload("TechnicianMan").
		Preload("PurchaseMan").Preload("InventoryMan").Preload("ShipmentMan").
		Where(&task).Find(&tasks).Error
	if err != nil {
		return tasks, msg.ERROR
	}
	return tasks, msg.SUCCESS
}

func SelectMyTasks(uid string, status int) (tasks []Task, code int) {
	var maps = make(map[string]interface{})
	maps["status"] = status
	switch status {
	case magic.TASK_STATUS_NOT_DESIGN:
		maps["technician_man_uid"] = uid
	case magic.TASK_STATUS_NOT_PURCHASE:
		maps["purchase_man_uid"] = uid
	case magic.TASK_STATUS_NOT_STORAGE:
		maps["inventory_man_uid"] = uid
	case magic.TASK_STATUS_NOT_SHIPMENT:
		maps["shipment_man_uid"] = uid
	case magic.TASK_STATUS_NOT_CONFIRM:
		maps["shipment_man_uid"] = uid
	}
	err = db.Preload("Product").Preload("TechnicianMan").
		Preload("PurchaseMan").Preload("InventoryMan").Preload("ShipmentMan").
		Where(maps).Find(&tasks).Error
	if err != nil {
		return tasks, msg.ERROR
	}
	return tasks, msg.SUCCESS
}

func SelectTasksByContractUID(contractUID string) (tasks []Task, code int) {
	err = db.Where("contract_uid = ?", contractUID).Find(&tasks).Error
	if err != nil {
		return tasks, msg.ERROR
	}
	return tasks, msg.SUCCESS
}

func ApproveTask(taskFlowQuery *TaskFlowQuery) (code int) {
	var maps = make(map[string]interface{})
	maps["status"] = taskFlowQuery.Status
	if taskFlowQuery.TechnicianManUID != "" {
		maps["TechnicianManUID"] = taskFlowQuery.TechnicianManUID
	} else {
		maps["TechnicianManUID"] = nil
	}
	if taskFlowQuery.PurchaseManUID != "" {
		maps["PurchaseManUID"] = taskFlowQuery.PurchaseManUID
	} else {
		maps["PurchaseManUID"] = nil
	}
	if taskFlowQuery.InventoryManUID != "" {
		maps["InventoryManUID"] = taskFlowQuery.InventoryManUID
	} else {
		maps["InventoryManUID"] = nil
	}
	if taskFlowQuery.ShipmentManUID != "" {
		maps["ShipmentManUID"] = taskFlowQuery.ShipmentManUID
	} else {
		maps["ShipmentManUID"] = nil
	}

	err = db.Model(&Task{}).Where("uid = ?", taskFlowQuery.UID).Updates(maps).Error

	if err != nil {
		code = msg.ERROR_CONTRACT_UPDATE_STATUS
	} else {
		code = msg.SUCCESS
	}
	return
}

func NextTaskStatus(uid string, status int, nextRemarks string) (code int) {
	var maps = make(map[string]interface{})
	maps["status"] = status
	maps["next_remarks"] = nextRemarks

	err = db.Model(&Task{}).Where("uid = ?", uid).Updates(maps).Error

	if err != nil {
		code = msg.ERROR_CONTRACT_UPDATE_STATUS
	} else {
		code = msg.SUCCESS
	}
	return
}
