package model

import (
	"business-system_golang/utils/magic"
	"business-system_golang/utils/msg"
	uidUtils "business-system_golang/utils/uid"
	"time"

	"gorm.io/gorm"
)

// 合同任务 Model
type Task struct {
	BaseModel
	UID              string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractUID      string  `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	ProductUID       string  `gorm:"type:varchar(32);comment:产品ID" json:"productUID"`
	Number           int     `gorm:"type:int;comment:数量" json:"number"`
	Unit             string  `gorm:"type:varchar(9);comment:单位" json:"unit"`
	Price            float64 `gorm:"type:decimal(20,6);comment:单价(元)" json:"price"`
	TotalPrice       float64 `gorm:"type:decimal(20,6);comment:总价(元)" json:"totalPrice"`
	Status           int     `gorm:"type:int;comment:状态" json:"status"`
	Type             int     `gorm:"type:int;comment:状态(1:标准/第三方有库存 2:标准/第三方无库存 3:非标准定制)" json:"type"`
	TechnicianManUID string  `gorm:"type:varchar(32);comment:技术负责人ID;default:(-)" json:"technicianManUID"`
	PurchaseManUID   string  `gorm:"type:varchar(32);comment:采购负责人ID;default:(-)" json:"purchaseManUID"`
	InventoryManUID  string  `gorm:"type:varchar(32);comment:库存负责人ID;default:(-)" json:"inventoryManUID"`
	ShipmentManUID   string  `gorm:"type:varchar(32);comment:发货人员ID;default:(-)" json:"shipmentManUID"`
	CurrentRemarks   string  `gorm:"type:varchar(32);comment:上一级备注;default:(-)" json:"currentRemarks"`
	Remarks          string  `gorm:"type:varchar(600);comment:备注" json:"remarks"`

	TechnicianDays        int   `gorm:"type:int;comment:技术预计工作天数;default:(-)" json:"technicianDays"`
	PurchaseDays          int   `gorm:"type:int;comment:采购预计工作天数;default:(-)" json:"purchaseDays"`
	TechnicianStartDate   XTime `gorm:"type:datetime;comment:技术开始工作日期;default:(-)" json:"technicianStartDate"`
	TechnicianRealEndDate XTime `gorm:"type:datetime;comment:技术实际结束工作日期;default:(-)" json:"technicianRealEndDate"`
	PurchaseStartDate     XTime `gorm:"type:datetime;comment:采购开始工作日期";default:(-) json:"purchaseStartDate"`
	PurchaseRealEndDate   XTime `gorm:"type:datetime;comment:采购实际结束工作日期;default:(-)" json:"purchaseRealEndDate"`

	CurrentRemarksText string `gorm:"-" json:"currentRemarksText"`

	Contract      Contract    `gorm:"foreignKey:ContractUID;references:UID" json:"contract"`
	Product       Product     `gorm:"foreignKey:ProductUID;references:UID" json:"product"`
	TechnicianMan Employee    `gorm:"foreignKey:TechnicianManUID;references:UID" json:"technicianMan"`
	PurchaseMan   Employee    `gorm:"foreignKey:PurchaseManUID;references:UID" json:"purchaseMan"`
	InventoryMan  Employee    `gorm:"foreignKey:InventoryManUID;references:UID" json:"inventoryMan"`
	ShipmentMan   Employee    `gorm:"foreignKey:ShipmentManUID;references:UID" json:"shipmentMan"`
	TaskRemarks   TaskRemarks `gorm:"foreignKey:CurrentRemarks;references:UID" json:"taskRemarks"`
}

type TaskRemarks struct {
	BaseModel
	UID     string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	TaskUID string `gorm:"type:varchar(32);comment:合同ID" json:"taskUID"`
	Status  int    `gorm:"type:int;comment:合同状态" json:"status"`
	From    string `gorm:"type:varchar(32);comment:填写人" json:"from"`
	To      string `gorm:"type:varchar(32);comment:接受人" json:"to"`
	Text    string `gorm:"type:varchar(499);comment:备注文本" json:"text"`
}

type TaskFlowQuery struct {
	UID              string `json:"UID"`
	ContractUID      string `json:"contractUID"`
	Status           int    `json:"status"`
	Type             int    `json:"type"`
	TechnicianManUID string `json:"technicianManUID"`
	PurchaseManUID   string `json:"purchaseManUID"`
	InventoryManUID  string `json:"inventoryManUID"`
	ShipmentManUID   string `json:"shipmentManUID"`
	IsReset          bool   `json:"isReset"`

	TechnicianDays int `json:"technicianDays"`
	PurchaseDays   int `json:"purchaseDays"`
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
	err = db.Preload("Contract").Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).
		Preload("TechnicianMan").Preload("PurchaseMan").Preload("InventoryMan").Preload("ShipmentMan").
		Where(&task).Find(&tasks).Error
	if err != nil {
		return tasks, msg.ERROR
	}
	return tasks, msg.SUCCESS
}

func SelectMyTasks(pageSize int, pageNo int, task *Task, uid string) (tasks []Task, code int, total int64) {
	var maps = make(map[string]interface{})
	if task.Status != 0 {
		maps["task.status"] = task.Status
	}
	maps["Contract.status"] = 2

	err = db.Joins("Contract").Where(maps).
		Where(db.Where("technician_man_uid = ?", uid).
			Or("purchase_man_uid = ?", uid).
			Or("inventory_man_uid = ?", uid).
			Or("shipment_man_uid = ?", uid)).
		Find(&tasks).Count(&total).
		Preload("Product").Preload("TechnicianMan").
		Preload("PurchaseMan").Preload("InventoryMan").Preload("ShipmentMan").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&tasks).Error
	if err != nil {
		return tasks, msg.ERROR, total
	}
	return tasks, msg.SUCCESS, total
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
	maps["type"] = taskFlowQuery.Type
	switch taskFlowQuery.Type {
	case magic.TASK_TYPE_1:
		maps["Status"] = magic.TASK_STATUS_NOT_STORAGE

		maps["TechnicianManUID"] = nil
		maps["PurchaseManUID"] = nil
		maps["TechnicianDays"] = nil
		maps["PurchaseDays"] = nil
		maps["TechnicianStartDate"] = nil
		maps["PurchaseStartDate"] = nil

	case magic.TASK_TYPE_2:
		maps["Status"] = magic.TASK_STATUS_NOT_PURCHASE

		maps["TechnicianManUID"] = nil
		maps["PurchaseManUID"] = taskFlowQuery.PurchaseManUID
		maps["TechnicianDays"] = nil
		maps["PurchaseDays"] = taskFlowQuery.PurchaseDays
		t := time.Now().Format("2006-01-02 15:04:05")
		maps["TechnicianStartDate"] = nil
		maps["PurchaseStartDate"] = t

	case magic.TASK_TYPE_3:
		maps["Status"] = magic.TASK_STATUS_NOT_DESIGN

		maps["TechnicianManUID"] = taskFlowQuery.TechnicianManUID
		maps["PurchaseManUID"] = taskFlowQuery.PurchaseManUID
		maps["TechnicianDays"] = taskFlowQuery.TechnicianDays
		maps["PurchaseDays"] = taskFlowQuery.PurchaseDays
		t := time.Now().Format("2006-01-02 15:04:05")
		maps["TechnicianStartDate"] = t
		maps["PurchaseStartDate"] = nil

	}
	maps["InventoryManUID"] = taskFlowQuery.InventoryManUID
	maps["ShipmentManUID"] = taskFlowQuery.ShipmentManUID

	if taskFlowQuery.IsReset {
		err = db.Transaction(func(tdb *gorm.DB) error {
			if tErr := tdb.Model(&Task{}).Where("uid = ?", taskFlowQuery.UID).Updates(maps).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Where("task_uid = ?", taskFlowQuery.UID).Delete(&TaskRemarks{}).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Model(&Contract{}).Where("uid = ?", taskFlowQuery.ContractUID).Update("production_status", 1).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	} else {
		err = db.Model(&Task{}).Where("uid = ?", taskFlowQuery.UID).Updates(maps).Error
	}

	if err != nil {
		code = msg.ERROR_CONTRACT_UPDATE_STATUS
	} else {
		code = msg.SUCCESS
	}
	return
}

func NextTaskStatus(uid string, lastStatus int, from string, to string, status int, currentRemarksText string) (code int) {
	var maps = make(map[string]interface{})
	maps["status"] = status
	err = db.Transaction(func(tdb *gorm.DB) error {
		if currentRemarksText != "" {
			remarksUID := uidUtils.Generate()
			if tErr := tdb.Model(&TaskRemarks{}).Create(map[string]interface{}{
				"CreatedAt": time.Now().Format("2006-01-02 15:04:05"),
				"UID":       remarksUID,
				"TaskUID":   uid,
				"Status":    lastStatus,
				"From":      from,
				"To":        to,
				"Text":      currentRemarksText,
			}).Error; tErr != nil {
				return tErr
			}
		}
		if tErr := tdb.Model(&Task{}).Where("uid = ?", uid).Updates(maps).Error; tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		code = msg.ERROR_CONTRACT_UPDATE_STATUS
	} else {
		code = msg.SUCCESS
	}
	return
}

func SelectTaskRemarks(taskUID string, to string) (taskRemarksList []TaskRemarks, code int) {
	err = db.Where("task_uid = ? AND `to` = ?", taskUID, to).Find(&taskRemarksList).Error
	if err != nil {
		code = msg.ERROR
	} else {
		code = msg.SUCCESS
	}
	return
}
