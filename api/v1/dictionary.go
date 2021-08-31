package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/rbac"

	"github.com/gin-gonic/gin"
)

func QueryDictionaryType(c *gin.Context) {
	var dictionaryType model.DictionaryType
	module := c.Query("module")
	name := c.Query("name")
	dictionaryType, code = model.SelectDictionaryType(module, name)
	msg.Message(c, code, dictionaryType)
}

func QueryDictionaryTypes(c *gin.Context) {
	module := c.Query("module")
	var dictionaryTypes []model.DictionaryType
	dictionaryTypes, code = model.SelectDictionaryTypes(module)
	msg.Message(c, code, dictionaryTypes)
}

func AddDictionary(c *gin.Context) {
	code = rbac.Check(c, "system", "all")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var dictionary model.Dictionary
		_ = c.ShouldBindJSON(&dictionary)
		code = model.InsertDictionary(&dictionary)
		msg.Message(c, code, dictionary)
	}
}

func DelDictionary(c *gin.Context) {
	code = rbac.Check(c, "system", "all")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		uid := c.Param("uid")
		code = model.DeleteDictionary(uid)
		msg.Message(c, code, nil)
	}
}

//字典表查询
func QueryDictionaries(c *gin.Context) {
	var dictionaries []model.Dictionary
	parentUID := c.Query("parentUID")
	DictionaryTypeUID := c.Query("dictionaryTypeUID")
	dictionaries, code = model.SelectDictionaries(parentUID, DictionaryTypeUID)

	msg.Message(c, code, dictionaries)
}
