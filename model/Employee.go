package model

import "gorm.io/gorm"

// 员工 Model
type Employee struct {
	gorm.Model
	Phone    string `gorm:"type:varchar(20);comment:电话;not null" json:"phone"`
	Name     string `gorm:"type:varchar(20);comment:姓名;not null" json:"name"`
	Password string `gorm:"type:varchar(20);comment:密码;not null" json:"password"`
	AreaId   int    `gorm:"type:varchar(20);comment:所属区域ID;not null" json:"aread"`
	WechatID string `gorm:"type:varchar(20);comment:微信号" json:"wechatId"`
	Email    string `gorm:"type:varchar(20);comment:邮箱" json:"email"`
}
