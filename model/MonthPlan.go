package model

import "business-system_golang/utils/msg"

type MonthPlan struct {
	ID    uint    `gorm:"primary_key" json:"ID"`
	Text  string  `gorm:"type:varchar(9);comment:名称" json:"text"`
	No    int     `gorm:"type:int;comment:月份" json:"no"`
	Value float64 `gorm:"type:decimal(20,6);comment:百分比" json:"value"`
}

func UpdateMonthPlan(monthPlan *MonthPlan) (code int) {
	err = db.Model(&MonthPlan{}).Where("id = ?", monthPlan.ID).Update("value", monthPlan.Value).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectMonthPlans() (monthPlans []MonthPlan, code int) {
	err = db.Find(&monthPlans).Error
	if err != nil {
		code = msg.ERROR
	} else {
		code = msg.SUCCESS
	}
	return
}
