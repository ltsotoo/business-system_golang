package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"
)

type TaskProcurement struct {
	BaseModel
	UID            string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractUID    string `gorm:"type:varchar(32);comment:合同UID" json:"contractUID"`
	TaskUID        string `gorm:"type:varchar(32);comment:任务UID" json:"taskUID"`
	EmployeeUID    string `gorm:"type:varchar(32);comment:采购发起人UID" json:"employeeUID"`
	PurchaseManUID string `gorm:"type:varchar(32);comment:采购负责人ID;default:(-)" json:"purchaseManUID"`
	Text           string `gorm:"type:varchar(499);comment:采购内容" json:"text"`
	Status         int    `gorm:"type:int;comment:状态(-1:采购取消，1:采购中，2:采购完成)" json:"status"`
}

func InsertTaskProcurements(taskProcurement *TaskProcurement) (code int) {
	taskProcurement.UID = uid.Generate()
	err = db.Create(&taskProcurement).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateTaskProcurement(uid string, status int) (code int) {
	var maps = make(map[string]interface{})
	maps["status"] = status

	err = db.Model(&TaskProcurement{}).Where("uid = ?", uid).Updates(maps).Error

	if err != nil {
		code = msg.ERROR
	} else {
		code = msg.SUCCESS
	}
	return
}

func SelectMyApplicationTaskProcurements(employeeUID string) (taskProcurements []TaskProcurement, code int) {
	err = db.Where("employee_uid = ?", employeeUID).Find(&taskProcurements).Error
	if err != nil {
		return taskProcurements, msg.ERROR
	}
	return taskProcurements, msg.SUCCESS
}

func SelectMyTaskProcurements(purchaseManUID string) (taskProcurements []TaskProcurement, code int) {
	err = db.Where("purchase_man_uid = ?", purchaseManUID).Find(&taskProcurements).Error
	if err != nil {
		return taskProcurements, msg.ERROR
	}
	return taskProcurements, msg.SUCCESS
}
