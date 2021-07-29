package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 客户公司 Model
type Company struct {
	gorm.Model
	Name    string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Address string `gorm:"type:varchar(20);comment:地址;not null" json:"address"`
}

func SelectCompanys() (companys []Company, code int) {
	err = db.Find(&companys).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return companys, msg.SUCCESS
}
