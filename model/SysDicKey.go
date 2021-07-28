package model

import "gorm.io/gorm"

// 字典表目录
type SystemDictionaryKey struct {
	gorm.Model
	ModelName string `gorm:"type:varchar(20);comment:模块;not null" json:"modelName"`
	Name      string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
}
