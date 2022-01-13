package model

import (
	"business-system_golang/utils/msg"
	"time"

	"gorm.io/gorm"
)

type YearPlan struct {
	ID            uint    `gorm:"primary_key" json:"ID"`
	Year          int     `gorm:"type:int;comment:年份" json:"year"`
	OfficeUID     string  `gorm:"type:varchar(32);comment:唯一标识;not null" json:"officeUID"`
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

func EndYearPlan() (code int) {
	var offices []Office
	err = db.Where("is_delete is false").Find(&offices).Error
	if err != nil {
		return msg.ERROR
	}
	var can = true
	for i := range offices {
		if !offices[i].IsSubmit {
			can = false
			break
		}
	}
	if can {
		year := time.Now().Year()
		err = db.Transaction(func(tdb *gorm.DB) error {
			for i := range offices {
				//创建旧纪录
				var yearPlan YearPlan
				tdb.First(&yearPlan, "office_uid = ? AND year = ?", offices[i].UID, year)
				if yearPlan.ID == 0 {
					yearPlan.Year = year
					yearPlan.OfficeUID = offices[i].UID
					yearPlan.BusinessMoney = offices[i].BusinessMoney
					yearPlan.Money = offices[i].Money
					yearPlan.MoneyCold = offices[i].MoneyCold
					yearPlan.TaskLoad = offices[i].TaskLoad
					yearPlan.TargetLoad = offices[i].TargetLoad
					if tErr := tdb.Create(&yearPlan).Error; tErr != nil {
						return tErr
					}
				}
				//更新新纪录
				var maps = make(map[string]interface{})
				maps["business_money"] = offices[i].NewBusinessMoney
				maps["money"] = offices[i].NewMoney
				maps["money_cold"] = offices[i].NewMoneyCold
				maps["task_load"] = offices[i].NewTaskLoad
				maps["target_load"] = offices[i].NewTargetLoad
				maps["new_money"] = 0
				maps["new_business_money"] = 0
				maps["new_money_cold"] = 0
				maps["new_task_load"] = 0
				maps["new_target_load"] = 0
				maps["is_submit"] = false
				if tErr := tdb.Model(&Office{}).Where("uid = ?", offices[i].UID).Updates(maps).Error; tErr != nil {
					return tErr
				}
				//重置员工的合同累积次数
				if tErr := tdb.Model(&Employee{}).Where("1 = 1").Update("contract_count", 0).Error; tErr != nil {
					return tErr
				}
				//修改系统状态
				if tErr := tdb.Model(&System{}).Where("text = ?", "SystemSettlement").Update("value", false).Error; tErr != nil {
					return tErr
				}
			}
			return nil
		})
		if err != nil {
			return msg.ERROR
		}
		SystemSettlement = true
		return msg.SUCCESS
	} else {
		return msg.ERROR
	}
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
