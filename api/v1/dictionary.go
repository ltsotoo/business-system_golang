package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func AddDictionary(c *gin.Context) {
	var dictionary model.Dictionary
	var dictionaryType model.DictionaryType
	_ = c.ShouldBindJSON(&dictionary)
	dictionaryType, _ = model.SelectDictionaryType(dictionary.DictionaryTypeUID)
	if dictionaryType.UID != "" {
		dictionary.DictionaryTypeUID = dictionaryType.UID
		code = model.InsertDictionary(&dictionary)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, dictionary)
}

func DelDictionary(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteDictionary(uid)
	msg.Message(c, code, nil)
}

//字典表查询
func QueryDictionaries(c *gin.Context) {
	var dictionaries []model.Dictionary
	text := c.Query("Text")
	typeName := c.Query("TypeName")
	dictionaries, code = model.SelectDictionaries(typeName, text)

	msg.Message(c, code, dictionaries)
}

func QueryDictionarieTypes(c *gin.Context) {
	var dictionarieTypes []model.DictionaryType
	category := c.Query("Category")
	dictionarieTypes, code = model.SelectDictionaryTypes(category)

	msg.Message(c, code, dictionarieTypes)
}
