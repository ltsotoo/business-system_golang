package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//通过KeyId查询某一类别下的所有字典数据
func QuerySystemDictionaryValuesByKeyId(c *gin.Context) {
	keyId, keyIdErr := strconv.Atoi(c.Query("keyId"))
	var systemDictionaryValues []model.SystemDictionaryValue
	var code int
	if keyIdErr == nil {
		systemDictionaryValues, code = model.SelectValuesBykeyId(keyId)
	}
	msg.Message(c, code, systemDictionaryValues)
}

//通过ParentId查询某一类别下的所有字典数据
func QuerySystemDictionaryValuesByParentId(c *gin.Context) {
	parentId, parentIdErr := strconv.Atoi(c.Query("parentId"))
	var systemDictionaryValues []model.SystemDictionaryValue
	var code int
	if parentIdErr == nil {
		systemDictionaryValues, code = model.SelectValuesByParentId(parentId)
	}
	msg.Message(c, code, systemDictionaryValues)
}
