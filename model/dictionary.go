package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

//字典类型表 Model
type DictionaryType struct {
	BaseModel
	UID  string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Text string `gorm:"type:varchar(20);comment:文本;not null" json:"text"`

	Dictionaries []Dictionary `gorm:"foreignKey:DictionaryTypeUID;references:UID" json:"dictionaries"`
}

//字典表 Model
type Dictionary struct {
	BaseModel
	UID               string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	DictionaryTypeUID string `gorm:"type:varchar(32);comment:类型ID;default:(-)" json:"dictionaryTypeUID"`
	Text              string `gorm:"type:varchar(20);comment:文本" json:"text"`

	DictionaryType DictionaryType `gorm:"foreignKey:DictionaryTypeUID;references:UID" json:"dictionaryType"`
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

func SelectDictionaryType(name string) (dictionaryType DictionaryType, code int) {
	err = db.Preload("Dictionaries").Where("name = ?", name).First(&dictionaryType).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return dictionaryType, msg.ERROR_SYSTE_DIC_TYPE_SELECT
	}
	return dictionaryType, msg.SUCCESS
}

func InsertDictionary(dictionary *Dictionary) (code int) {
	dictionary.UID = uid.Generate()
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

func SelectDictionaries(name string, text string) (dictionaries []Dictionary, code int) {
	err = db.Joins("DictionaryType").Where("DictionaryType.name = ? AND dictionary.text LIKE ?", name, "%"+text+"%").Find(&dictionaries).Error
	if err != nil {
		return dictionaries, msg.ERROR
	}
	return dictionaries, msg.SUCCESS
}
