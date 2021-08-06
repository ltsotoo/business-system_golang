package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 客户公司 Model
type Company struct {
	gorm.Model
	AreaID  uint   `gorm:"int;comment:地区ID;default:(-)" json:"areaID"`
	Name    string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Address string `gorm:"type:varchar(20);comment:地址" json:"address"`

	Area Area `gorm:"foreignKey:AreaID" json:"area"`
}

func SelectCompanys() (companys []Company, code int) {
	err = db.Find(&companys).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return companys, msg.SUCCESS
}

func SelectCompanysByAreaID(areaID int) (companys []Company, code int) {
	err = db.Where("area_id = ?", areaID).Find(&companys).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return companys, msg.SUCCESS
}
