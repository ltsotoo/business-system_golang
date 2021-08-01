package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DelTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteTask(id)
	msg.Message(c, code, nil)
}

func QueryTasksByContractID(c *gin.Context) {
	var tasks []model.Task
	conetractID, conetractIDErr := strconv.Atoi(c.Query("contractID"))
	if conetractIDErr == nil {
		tasks, code = model.SelectTaskByContractID(conetractID)
	}
	msg.Message(c, code, tasks)
}
