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

func QueryTasks(c *gin.Context) {
	var tasks []model.Task
	var task model.Task
	_ = c.ShouldBindJSON(&task)
	tasks, code = model.SelectTasks(&task)
	msg.Message(c, code, tasks)
}
