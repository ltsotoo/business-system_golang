package model

import "gorm.io/gorm"

// 客户 Model
type Customer struct {
	gorm.Model
	Name          string `gorm:"type:varchar(20);comment:姓名;not null" json:"name"`
	Address       string `gorm:"type:varchar(20);comment:地址;not null" json:"address"`
	Company       string `gorm:"type:varchar(20);comment:公司;not null" json:"company"`
	ResearchGroup string `gorm:"type:varchar(20);comment:课题组;not null" json:"researchGroup"`
	Phone         string `gorm:"type:varchar(20);comment:电话;not null" json:"phone"`
	WechatID      string `gorm:"type:varchar(20);comment:微信号" json:"wechatID"`
	Email         string `gorm:"type:varchar(20);comment:邮箱" json:"email"`
}
