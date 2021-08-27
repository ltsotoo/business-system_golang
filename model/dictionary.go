package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

//字典类型表 Model
type DictionaryType struct {
	gorm.Model
	UID       string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ParentUID string `gorm:"type:varchar(32);comment:父UID" json:"parentUID"`
	Module    string `gorm:"type:varchar(20);comment:模块;not null" json:"module"`
	Name      string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Text      string `gorm:"type:varchar(20);comment:文本;not null" json:"text"`

	Dictionaries []Dictionary `gorm:"foreignKey:DictionaryTypeUID;references:UID" json:"dictionaries"`
}

//字典表 Model
type Dictionary struct {
	gorm.Model
	UID               string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ParentUID         string `gorm:"type:varchar(32);comment:父ID" json:"parentUID"`
	DictionaryTypeUID string `gorm:"type:varchar(32);comment:类型ID;default:(-)" json:"dictionaryTypeUID"`
	Text              string `gorm:"type:varchar(20);comment:文本" json:"text"`
}

func InsertDictionaryType(dictionaryType *DictionaryType) (code int) {
	dictionaryType.UID = uid.Generate()
	err = db.Create(&dictionaryType).Error
	if err != nil {
		return msg.ERROR_SYSTE_DIC_TYPE_INSERT
	}
	return msg.SUCCESS
}

func DeleteDictionaryType(uid string) (code int) {
	err = db.Where("uid = ?", uid).Delete(&DictionaryType{}).Error
	if err != nil {
		return msg.ERROR_SYSTE_DIC_TYPE_DELETE
	}
	return msg.SUCCESS
}

func SelectDictionaryType(module string, name string) (dictionaryType DictionaryType, code int) {
	err = db.Preload("Dictionaries").Where("module = ? AND name = ?", module, name).First(&dictionaryType).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return dictionaryType, msg.ERROR_SYSTE_DIC_TYPE_SELECT
	}
	return dictionaryType, msg.SUCCESS
}

func SelectDictionaryTypes(module string) (dictionaryTypes []DictionaryType, code int) {
	err = db.Where("module = ?", module).Find(&dictionaryTypes).Error
	if err != nil {
		return nil, msg.ERROR_SYSTE_DIC_TYPE_SELECT
	}
	return dictionaryTypes, msg.SUCCESS
}

func InsertDictionary(dictionary *Dictionary) (code int) {
	err = db.Create(&dictionary).Error
	if err != nil {
		return msg.ERROR_SYSTE_DIC_INSERT
	}
	return msg.SUCCESS
}

func DeleteDictionary(uid string) (code int) {
	err = db.Where("uid = ?", uid).Delete(&Dictionary{}).Error
	if err != nil {
		return msg.ERROR_SYSTE_DIC_DELETE
	}
	return msg.SUCCESS
}

func SelectDictionaries(parentUID string, dictionaryTypeUID string) (dictionaries []Dictionary, code int) {
	var maps = make(map[string]interface{})
	if parentUID != "" {
		maps["parent_uid"] = parentUID
	}
	if dictionaryTypeUID != "" {
		maps["dictionary_type_uid"] = dictionaryTypeUID
	}
	err = db.Where(maps).Find(&dictionaries).Error
	if err != nil {
		return dictionaries, msg.ERROR
	}
	return dictionaries, msg.SUCCESS
}

func (dictionaryType *DictionaryType) BeforeCreate(tx *gorm.DB) (err error) {
	dictionaryType.UID = uid.Generate()
	return err
}
