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
	Name string `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	Text string `gorm:"type:varchar(50);comment:文本;not null" json:"text"`

	Dictionaries []Dictionary `gorm:"foreignKey:DictionaryTypeUID;references:UID" json:"dictionaries"`
}

//字典表 Model
type Dictionary struct {
	BaseModel
	UID               string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	DictionaryTypeUID string `gorm:"type:varchar(32);comment:类型ID;default:(-)" json:"dictionaryTypeUID"`
	Text              string `gorm:"type:varchar(50);comment:文本" json:"text"`
	IsDelete          bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

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

func SelectDictionaryType(name string) (dictionaryType DictionaryType, code int) {
	err = db.Preload("Dictionaries").Where("name = ?", name).First(&dictionaryType).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return dictionaryType, msg.ERROR_SYSTE_DIC_TYPE_SELECT
	}
	return dictionaryType, msg.SUCCESS
}

func SelectDictionaryTypes(category string) (dictionaryTypes []DictionaryType, code int) {
	err = db.Where("category = ?", category).Find(&dictionaryTypes).Error
	if err != nil {
		return dictionaryTypes, msg.ERROR_SYSTE_DIC_TYPE_SELECT
	}
	return dictionaryTypes, msg.SUCCESS
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
	// err = db.Where("uid = ?", uid).Delete(&Dictionary{}).Error
	err = db.Model(&Dictionary{}).Where("uid = ?", uid).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR_SYSTE_DIC_DELETE
	}
	return msg.SUCCESS
}

func SelectDictionaries(name string, text string) (dictionaries []Dictionary, code int) {
	err = db.Joins("DictionaryType").Where("DictionaryType.name = ? AND is_delete = ? AND dictionary.text LIKE ?", name, false, "%"+text+"%").Find(&dictionaries).Error
	if err != nil {
		return dictionaries, msg.ERROR
	}
	return dictionaries, msg.SUCCESS
}
