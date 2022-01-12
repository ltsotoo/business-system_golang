package model

import "business-system_golang/utils/msg"

type YearPlan struct {
	ID            uint    `gorm:"primary_key" json:"ID"`
	Year          int     `gorm:"type:int;comment:年份" json:"year"`
	UID           string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name          string  `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	Number        string  `gorm:"type:varchar(50);comment:编号" json:"number"`
	BusinessMoney float64 `gorm:"type:decimal(20,6);comment:业务费用(元)" json:"businessMoney"`
	Money         float64 `gorm:"type:decimal(20,6);comment:办事处目前可报销额度(元)" json:"money"`
	MoneyCold     float64 `gorm:"type:decimal(20,6);comment:办事处今年冻结报销额度(元)" json:"moneyCold"`
	TaskLoad      float64 `gorm:"type:decimal(20,6);comment:今年目标量(元)" json:"taskLoad"`
	TargetLoad    float64 `gorm:"type:decimal(20,6);comment:今年完成量(元)" json:"targetLoad"`
}

func StartYearPlan() (code int) {
	//修改系统状态
	err = db.Model(&System{}).Where("text = ?", "SystemSettlement").Update("value", true).Error
	if err != nil {
		return msg.ERROR
	}
	SystemSettlement = true
	return msg.SUCCESS
}

func UpdateYearOffice(office *Office) (code int) {
	var maps = make(map[string]interface{})
	maps["new_money"] = office.NewMoney
	maps["new_business_money"] = office.NewBusinessMoney
	maps["new_money_cold"] = office.NewMoneyCold
	maps["new_task_load"] = office.NewTaskLoad
	maps["new_target_load"] = office.NewTargetLoad
	maps["is_submit"] = true

	err = db.Model(&Office{}).Where("uid = ?", office.UID).Updates(maps).Error

	if err != nil {
		return msg.ERROR
	} else {
		return msg.SUCCESS
	}
}
