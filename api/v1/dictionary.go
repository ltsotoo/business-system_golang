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

//字典表查询
func QueryDictionaries(c *gin.Context) {
	parentID, parentIDErr := strconv.Atoi(c.DefaultQuery("parentID", "0"))
	module := c.Query("module")
	name := c.Query("name")
	var dictionaries []model.Dictionary

	if parentIDErr == nil {
		dictionaries, code = model.SelectDictionaries(parentID, module, name)
	}

	msg.Message(c, code, dictionaries)
}
