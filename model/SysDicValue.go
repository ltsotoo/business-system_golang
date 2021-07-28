package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 字典表
type SystemDictionaryValue struct {
	gorm.Model
	KeyId    uint   `gorm:"type:int;comment:编号;not null" json:"keyId"`
	ParentId uint   `gorm:"type:int;comment:父ID" json:"parentId"`
	Name     string `gorm:"type:varchar(20);comment:名称" json:"name"`
}

func SelectValuesBykeyId(keyId int) (systemDictionaryValues []SystemDictionaryValue, code int) {
	err = db.Where("key_id = ?", keyId).Find(&systemDictionaryValues).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return systemDictionaryValues, msg.SUCCESS
}

func SelectValuesByParentId(parentId int) (systemDictionaryValues []SystemDictionaryValue, code int) {
	err = db.Where("parent_id = ?", parentId).Find(&systemDictionaryValues).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return systemDictionaryValues, msg.SUCCESS
}
