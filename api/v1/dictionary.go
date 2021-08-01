package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//字典表查询
func QueryDictionaries(c *gin.Context) {
	name := c.Query("name")
	parentID, parentIDErr := strconv.Atoi(c.DefaultQuery("parentID", "0"))
	var dictionaries []model.Dictionary

	if parentIDErr == nil {
		dictionaries, code = model.SelectDictionaries(name, parentID)
	}

	msg.Message(c, code, dictionaries)
}
