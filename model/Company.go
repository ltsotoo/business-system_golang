package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 客户公司 Model
type Company struct {
	gorm.Model
	UID     string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	AreaUID uint   `gorm:"int;comment:地区ID;default:(-)" json:"areaUID"`
	Name    string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Address string `gorm:"type:varchar(20);comment:地址" json:"address"`

	Area Area `gorm:"foreignKey:AreaUID;references:UID" json:"area"`
}

func SelectCompany(id int) (company Company, code int) {
	err = db.Preload("Area").First(&company, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return company, msg.ERROR_CUSTOMER_COMPANY_NOT_EXIST
		} else {
			return company, msg.ERROR
		}
	}
	return company, msg.SUCCESS
}

func SelectCompanys(company *Company) (companys []Company, code int) {
	err = db.Where(&company).Find(&companys).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return companys, msg.SUCCESS
}
