package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

//区域 Model
type Area struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
}

func SelectAreas() (areas []Area, code int) {
	err = db.Find(&areas).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return areas, msg.SUCCESS
}
