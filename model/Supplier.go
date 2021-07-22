package model

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Address  string `gorm:"type:varchar(20);comment:地址;not null" json:"address"`
	Linkman  string `gorm:"type:varchar(20);comment:联系人名称;not null" json:"linkman"`
	Phone    string `gorm:"type:varchar(20);comment:联系电话;not null" json:"phone"`
	WechatID string `gorm:"type:varchar(20);comment:微信号" json:"wechatID"`
	Email    string `gorm:"type:varchar(20);comment:邮箱" json:"email"`
}
