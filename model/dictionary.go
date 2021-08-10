package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

//字典表 Model
type Dictionary struct {
	gorm.Model
	ParentID    uint   `gorm:"type:int;comment:父ID" json:"parentID"`
	Module      string `gorm:"type:varchar(20);comment:模块" json:"module"`
	Name        string `gorm:"type:varchar(20);comment:名称" json:"name"`
	Text        string `gorm:"type:varchar(20);comment:文本" json:"text"`
	Description string `gorm:"type:varchar(20);comment:描述" json:"description"`
}

func SelectDictionarieTextGroup(module string) (dictionaries []Dictionary, code int) {
	err = db.Select("name", "description").Distinct("name").Find(&dictionaries).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return dictionaries, msg.SUCCESS
}

func SelectDictionaries(parentID int, module string, name string) (dictionaries []Dictionary, code int) {
	var maps = make(map[string]interface{})
	if parentID != -1 {
		maps["parent_id"] = parentID
	}
	maps["module"] = module
	maps["name"] = name
	err = db.Where(maps).Find(&dictionaries).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return dictionaries, msg.SUCCESS
}
