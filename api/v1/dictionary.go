package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//字典表查询
func QueryDictionaries(c *gin.Context) {
	parentID, parentIDErr := strconv.Atoi(c.DefaultQuery("parentID", "0"))
	module := c.DefaultQuery("module", "")
	name := c.DefaultQuery("name", "")
	var dictionaries []model.Dictionary

	if parentIDErr == nil {
		dictionaries, code = model.SelectDictionaries(parentID, module, name)
	}

	msg.Message(c, code, dictionaries)
}

func QueryDictionarieTextGroup(c *gin.Context) {
	module := c.DefaultQuery("module", "")
	var dictionaries []model.Dictionary
	dictionaries, code = model.SelectDictionarieTextGroup(module)
	msg.Message(c, code, dictionaries)
}
