package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 字典表
type SystemDictionaryValue struct {
	gorm.Model
	KeyID    uint   `gorm:"type:int;comment:编号;not null" json:"keyID"`
	ParentID uint   `gorm:"type:int;comment:父ID" json:"parentID"`
	Name     string `gorm:"type:varchar(20);comment:名称" json:"name"`
}

func SelectValuesBykeyID(keyID int) (systemDictionaryValues []SystemDictionaryValue, code int) {
	err = db.Where("key_id = ?", keyID).Find(&systemDictionaryValues).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return systemDictionaryValues, msg.SUCCESS
}

func SelectValuesByParentID(parentID int) (systemDictionaryValues []SystemDictionaryValue, code int) {
	err = db.Where("parent_id = ?", parentID).Find(&systemDictionaryValues).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return systemDictionaryValues, msg.SUCCESS
}
