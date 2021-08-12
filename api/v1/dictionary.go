package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//字典表类型查询
func QueryDictionaryTypes(c *gin.Context) {
	module := c.Query("module")
	var dictionaryTypes []model.DictionaryType
	dictionaryTypes, code = model.SelectDictionaryTypes(module)
	msg.Message(c, code, dictionaryTypes)
}

func AddDictionary(c *gin.Context) {
	var dictionary model.Dictionary
	_ = c.ShouldBindJSON(&dictionary)
	code = model.CreateDictionary(&dictionary)
	msg.Message(c, code, dictionary)
}

func DelDictionary(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteDictionary(id)
	msg.Message(c, code, nil)
}

//字典表查询
func QueryDictionaries(c *gin.Context) {
	var dictionaries []model.Dictionary
	parentID, parentIDErr := strconv.Atoi(c.DefaultQuery("parentID", "0"))
	dictionaryTypeID, dictionaryTypeIDErr := strconv.Atoi(c.DefaultQuery("dictionaryTypeID", "0"))
	if parentIDErr == nil && dictionaryTypeIDErr == nil {
		dictionaries, code = model.SelectDictionaries(parentID, dictionaryTypeID)
	}
	msg.Message(c, code, dictionaries)
}

func QueryDictionariesByDictionaryType(c *gin.Context) {
	var dictionaries []model.Dictionary
	module := c.Query("module")
	name := c.Query("name")
	dictionaries, code = model.SelectDictionariesByDictionaryType(module, name)
	msg.Message(c, code, dictionaries)
}
