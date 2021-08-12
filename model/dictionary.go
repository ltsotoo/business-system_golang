package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

type DictionaryType struct {
	gorm.Model
	ParentID uint   `gorm:"type:int;comment:父ID" json:"parentID"`
	Module   string `gorm:"type:varchar(20);comment:模块;not null" json:"module"`
	Name     string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Text     string `gorm:"type:varchar(20);comment:文本;not null" json:"text"`
}

//字典表 Model
type Dictionary struct {
	gorm.Model
	ParentID         uint   `gorm:"type:int;comment:父ID" json:"parentID"`
	DictionaryTypeID uint   `gorm:"type:int;comment:类型ID;default:(-)" json:"dictionaryTypeID"`
	Text             string `gorm:"type:varchar(20);comment:文本" json:"text"`

	ParentText string `gorm:"<-:false" json:"parentText"`
}

func SelectDictionaryTypes(module string) (dictionaryTypes []DictionaryType, code int) {
	err = db.Where("module = ?", module).Find(&dictionaryTypes).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return dictionaryTypes, msg.SUCCESS
}

func CreateDictionary(dictionary *Dictionary) (code int) {
	err = db.Create(&dictionary).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteDictionary(id int) (code int) {
	err = db.Where("id = ?", id).Delete(&Dictionary{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectDictionaries(parentID int, dictionaryTypeID int) (dictionaries []Dictionary, code int) {
	var maps = make(map[string]interface{})
	if parentID > 0 {
		maps["parent_id"] = parentID
	}
	if dictionaryTypeID > 0 {
		maps["dictionary_type_id"] = dictionaryTypeID
	}
	err = db.Where(maps).Find(&dictionaries).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return dictionaries, msg.SUCCESS
}

func SelectDictionariesByDictionaryType(module string, name string) (dictionaries []Dictionary, code int) {
	err = db.Joins("left join dictionary_type on dictionary_type.id = dictionary.dictionary_type_id").Where("dictionary_type.module = ? AND dictionary_type.name = ?", module, name).Find(&dictionaries).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return dictionaries, msg.SUCCESS
}
