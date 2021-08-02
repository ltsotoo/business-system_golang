package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

//字典表 Model
type Dictionary struct {
	gorm.Model
	ParentID    uint   `gorm:"type:int;comment:编号" json:"parentID"`
	Name        string `gorm:"type:varchar(20);comment:名称" json:"name"`
	Value       string `gorm:"type:varchar(20);comment:value" json:"value"`
	Text        string `gorm:"type:varchar(20);comment:名称" json:"text"`
	Description string `gorm:"type:varchar(20);comment:描述" json:"description"`
}

func SelectDictionaries(name string, parentID int) (dictionaries []Dictionary, code int) {
	var maps = make(map[string]interface{})
	maps["name"] = name
	maps["parent_id"] = parentID
	err = db.Where(maps).Find(&dictionaries).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return dictionaries, msg.SUCCESS
}
