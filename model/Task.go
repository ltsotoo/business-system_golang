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
	UID                  string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractUID          string  `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	ProductUID           string  `gorm:"type:varchar(32);comment:产品ID" json:"productUID"`
	Number               int     `gorm:"type:int;comment:数量" json:"number"`
	Unit                 string  `gorm:"type:varchar(50);comment:单位" json:"unit"`
	StandardPrice        float64 `gorm:"type:decimal(20,6);comment:下单时标准价格" json:"standardPrice"`
	StandardPriceUSD     float64 `gorm:"type:decimal(20,6);comment:下单时标准价格(美元)" json:"standardPriceUSD"`
	Price                float64 `gorm:"type:decimal(20,6);comment:单价" json:"price"`
	TotalPrice           float64 `gorm:"type:decimal(20,6);comment:总价" json:"totalPrice"`
	PaymentTotalPriceUSD float64 `gorm:"type:decimal(20,6);comment:回款总金额(USD)" json:"paymentTotalPriceUSD"`
	PaymentTotalPrice    float64 `gorm:"type:decimal(20,6);comment:回款总金额(CNY)" json:"paymentTotalPrice"`
	Status               int     `gorm:"type:int;comment:状态(1:待设计 2:待采购 3:待入/出库 4:待装配 5:待发货 6:已发货)" json:"status"`
	Type                 int     `gorm:"type:int;comment:类型(1:标准/第三方有库存 2:标准/第三方无库存 3:非标准定制)" json:"type"`
	TechnicianManUID     string  `gorm:"type:varchar(32);comment:技术负责人ID;default:(-)" json:"technicianManUID"`
	PurchaseManUID       string  `gorm:"type:varchar(32);comment:采购负责人ID;default:(-)" json:"purchaseManUID"`
	InventoryManUID      string  `gorm:"type:varchar(32);comment:仓库负责人ID;default:(-)" json:"inventoryManUID"`
	ShipmentManUID       string  `gorm:"type:varchar(32);comment:物流人员ID;default:(-)" json:"shipmentManUID"`
	Remarks              string  `gorm:"type:varchar(600);comment:业务备注" json:"remarks"`
	ARemarks             string  `gorm:"type:varchar(600);comment:审批备注" json:"aRemarks"`
	PushMoney            float64 `gorm:"type:decimal(20,6);comment:提成(元)" json:"pushMoney"`
	PushMoneyPercentages float64 `gorm:"type:decimal(20,6);comment:特殊合同提成百分比" json:"pushMoneyPercentages"`

	TechnicianDays        int   `gorm:"type:int;comment:技术预计工作天数;default:(-)" json:"technicianDays"`
	PurchaseDays          int   `gorm:"type:int;comment:采购预计工作天数;default:(-)" json:"purchaseDays"`
	TechnicianStartDate   XTime `gorm:"type:datetime;comment:技术接到工作日期;default:(-)" json:"technicianStartDate"`
	TechnicianRealEndDate XTime `gorm:"type:datetime;comment:技术实际提交工作日期;default:(-)" json:"technicianRealEndDate"`
	PurchaseStartDate     XTime `gorm:"type:datetime;comment:采购接到工作日期;default:(-)" json:"purchaseStartDate"`
	PurchaseRealEndDate   XTime `gorm:"type:datetime;comment:采购实际提交工作日期;default:(-)" json:"purchaseRealEndDate"`
	InventoryStartDate    XTime `gorm:"type:datetime;comment:仓库接到工作日期;default:(-)" json:"inventoryStartDate"`
	InventoryRealEndDate  XTime `gorm:"type:datetime;comment:仓库实际提交工作日期;default:(-)" json:"inventoryRealEndDate"`
	ShipmentStartDate     XTime `gorm:"type:datetime;comment:物流接到工作日期;default:(-)" json:"shipmentStartDate"`
	ShipmentRealEndDate   XTime `gorm:"type:datetime;comment:物流实际提交工作日期;default:(-)" json:"shipmentRealEndDate"`

	Contract      Contract `gorm:"foreignKey:ContractUID;references:UID" json:"contract"`
	Product       Product  `gorm:"foreignKey:ProductUID;references:UID" json:"product"`
	TechnicianMan Employee `gorm:"foreignKey:TechnicianManUID;references:UID" json:"technicianMan"`
	PurchaseMan   Employee `gorm:"foreignKey:PurchaseManUID;references:UID" json:"purchaseMan"`
	InventoryMan  Employee `gorm:"foreignKey:InventoryManUID;references:UID" json:"inventoryMan"`
	ShipmentMan   Employee `gorm:"foreignKey:ShipmentManUID;references:UID" json:"shipmentMan"`
}

type TaskRemarks struct {
	UID     string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	TaskUID string `gorm:"type:varchar(32);comment:合同ID" json:"taskUID"`
	Status  int    `gorm:"type:int;comment:合同状态" json:"status"`
	From    string `gorm:"type:varchar(32);comment:填写人" json:"from"`
	To      string `gorm:"type:varchar(32);comment:接受人" json:"to"`
	Text    string `gorm:"type:varchar(600);comment:备注文本" json:"text"`
}

type TaskFlowQuery struct {
	UID                string `json:"UID"`
	ContractUID        string `json:"contractUID"`
	Status             int    `json:"status"`
	Type               int    `json:"type"`
	TechnicianManUID   string `json:"technicianManUID"`
	PurchaseManUID     string `json:"purchaseManUID"`
	InventoryManUID    string `json:"inventoryManUID"`
	ShipmentManUID     string `json:"shipmentManUID"`
	IsReset            bool   `json:"isReset"`
	ARemarks           string `json:"aRemarks"`
	CurrentRemarksText string ` json:"currentRemarksText"`

	TechnicianDays int `json:"technicianDays"`
	PurchaseDays   int `json:"purchaseDays"`
}

func SelectTask(uid string) (task Task, code int) {
	err = db.Where("uid = ?", uid).First(&task).Error
	if err != nil {
		return task, msg.ERROR
	}
	return task, msg.SUCCESS
}

func SelectTasks(task *Task) (tasks []Task, code int) {
	var maps = make(map[string]interface{})

	if task.ContractUID != "" {
		maps["contract_uid"] = task.ContractUID
	}

	err = db.Preload("Contract").Preload("Product").
		Preload("TechnicianMan").Preload("PurchaseMan").
		Preload("InventoryMan").Preload("ShipmentMan").
		Where(maps).Find(&tasks).Error
	if err != nil {
		code = msg.ERROR
	} else {
		code = msg.SUCCESS
	}
	return
}

func SelectMyTasks(pageSize int, pageNo int, task *Task, uid string) (tasks []Task, code int, total int64) {
	var maps = make(map[string]interface{})
	if task.Status != 0 {
		maps["task.status"] = task.Status
	}
	maps["Contract.status"] = 2
	maps["Contract.is_delete"] = false

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
	maps["a_remarks"] = taskFlowQuery.ARemarks
	switch taskFlowQuery.Type {
	case magic.TASK_TYPE_1:
		maps["status"] = magic.TASK_STATUS_NOT_STORAGE

		maps["technician_man_uid"] = nil
		maps["purchase_man_uid"] = nil
		maps["technician_days"] = nil
		maps["purchase_days"] = nil

		if taskFlowQuery.IsReset {
			t := time.Now().Format("2006-01-02 15:04:05")
			maps["inventory_start_date"] = t
		} else {
			maps["inventory_start_date"] = nil
		}

		maps["technician_start_date"] = nil
		maps["purchase_start_date"] = nil

	case magic.TASK_TYPE_2:
		maps["status"] = magic.TASK_STATUS_NOT_PURCHASE

		maps["technician_man_uid"] = nil
		maps["purchase_man_uid"] = taskFlowQuery.PurchaseManUID
		maps["technician_days"] = nil
		maps["purchase_days"] = taskFlowQuery.PurchaseDays

		if taskFlowQuery.IsReset {
			t := time.Now().Format("2006-01-02 15:04:05")
			maps["purchase_start_date"] = t
		} else {
			maps["purchase_start_date"] = nil
		}

		maps["technician_start_date"] = nil
		maps["inventory_start_date"] = nil

	case magic.TASK_TYPE_3:
		maps["status"] = magic.TASK_STATUS_NOT_DESIGN

		maps["technician_man_uid"] = taskFlowQuery.TechnicianManUID
		maps["purchase_man_uid"] = taskFlowQuery.PurchaseManUID
		maps["technician_days"] = taskFlowQuery.TechnicianDays
		maps["purchase_days"] = taskFlowQuery.PurchaseDays

		if taskFlowQuery.IsReset {
			t := time.Now().Format("2006-01-02 15:04:05")
			maps["technician_start_date"] = t
		} else {
			maps["technician_start_date"] = nil
		}

		maps["purchase_start_date"] = nil
		maps["inventory_start_date"] = nil

	}
	maps["shipment_start_date"] = nil
	maps["technician_real_end_date"] = nil
	maps["purchase_real_end_date"] = nil
	maps["inventory_real_end_date"] = nil
	maps["shipment_real_end_date"] = nil
	maps["inventory_man_uid"] = taskFlowQuery.InventoryManUID
	maps["shipment_man_uid"] = taskFlowQuery.ShipmentManUID

	if taskFlowQuery.IsReset {
		var contractMaps = make(map[string]interface{})
		contractMaps["production_status"] = 1
		contractMaps["end_delivery_date"] = nil
		err = db.Transaction(func(tdb *gorm.DB) error {
			if tErr := tdb.Model(&Task{}).Where("uid = ?", taskFlowQuery.UID).Updates(maps).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Where("task_uid = ?", taskFlowQuery.UID).Delete(&TaskRemarks{}).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Model(&Contract{}).Where("uid = ?", taskFlowQuery.ContractUID).Statement.Updates(contractMaps).Error; tErr != nil {
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

func NextTaskStatus(uid string, lastStatus int, from string, to string, maps map[string]interface{}, currentRemarksText string) (code int) {
	err = db.Transaction(func(tdb *gorm.DB) error {
		if currentRemarksText != "" {
			// remarksUID := uidUtils.Generate()
			// if tErr := tdb.Model(&TaskRemarks{}).Create(map[string]interface{}{
			// 	"UID":       remarksUID,
			// 	"TaskUID":   uid,
			// 	"Status":    lastStatus,
			// 	"From":      from,
			// 	"To":        to,
			// 	"Text":      currentRemarksText,
			// }).Error; tErr != nil {
			// 	return tErr
			// }

			if tErr := tdb.Create(&TaskRemarks{
				UID:     uidUtils.Generate(),
				TaskUID: uid,
				Status:  lastStatus,
				From:    from,
				To:      to,
				Text:    currentRemarksText,
			}).Error; tErr != nil {
				return tErr
			}
		}
		if tErr := tdb.Model(&Task{}).Where("uid = ?", uid).Updates(maps).Error; tErr != nil {
			return tErr
		}

		if maps["status"] == magic.TASK_STATUS_SHIPMENT {

			var tempTask Task

			if tErr := tdb.First(&tempTask, "uid = ?", uid).Error; tErr != nil {
				return tErr
			}

			if tErr := tdb.Exec("UPDATE product SET number_count = number_count - ? WHERE uid = ?", tempTask.Number, tempTask.ProductUID).Error; tErr != nil {
				return tErr
			}
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
	err = db.Where("task_uid = ? AND `from` = ?", taskUID, to).
		Or("task_uid = ? AND `to` = ?", taskUID, to).
		Find(&taskRemarksList).Error
	if err != nil {
		code = msg.ERROR
	} else {
		code = msg.SUCCESS
	}
	return
}
