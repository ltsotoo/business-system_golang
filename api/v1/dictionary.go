package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func AddDictionary(c *gin.Context) {
	var dictionary model.Dictionary
	_ = c.ShouldBindJSON(&dictionary)
	code = model.InsertDictionary(&dictionary)
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
	if typeName != "" {
		dictionaries, code = model.SelectDictionariesByTypeName(typeName, text)
	} else {
		parentUID := c.Query("ParentUID")
		if parentUID != "" {
			dictionaries, code = model.SelectDictionaries(parentUID, text)
		} else {
			code = msg.ERROR
		}
	}

	msg.Message(c, code, dictionaries)
}
