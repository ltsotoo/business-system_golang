package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//查询所有的区域
func QueryAreas(c *gin.Context) {
	var areas []model.Area
	areas, code = model.SelectAreas()
	msg.Message(c, code, areas)
}

//通过KeyId查询某一类别下的所有字典数据
func QuerySystemDictionaryValuesByKeyId(c *gin.Context) {
	keyID, keyIDErr := strconv.Atoi(c.Query("keyID"))
	var systemDictionaryValues []model.SystemDictionaryValue
	if keyIDErr == nil {
		systemDictionaryValues, code = model.SelectValuesBykeyID(keyID)
	}
	msg.Message(c, code, systemDictionaryValues)
}

//通过ParentId查询某一类别下的所有字典数据
func QuerySystemDictionaryValuesByParentId(c *gin.Context) {
	parentID, parentIDErr := strconv.Atoi(c.Query("parentID"))
	var systemDictionaryValues []model.SystemDictionaryValue
	if parentIDErr == nil {
		systemDictionaryValues, code = model.SelectValuesByParentID(parentID)
	}
	msg.Message(c, code, systemDictionaryValues)
}
