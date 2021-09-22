package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/rbac"

	"github.com/gin-gonic/gin"
)

func DelTask(c *gin.Context) {
	code = rbac.Check(c, "contract", "delete")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		uid := c.Param("uid")
		code = model.DeleteTask(uid)
		msg.Message(c, code, nil)
	}
}

func QueryTasks(c *gin.Context) {
	var tasks []model.Task
	var task model.Task
	_ = c.ShouldBindJSON(&task)
	tasks, code = model.SelectTasks(&task)
	msg.Message(c, code, tasks)
}

func ApproveTask(c *gin.Context) {
	var taskFlowQuery model.TaskFlowQuery
	_ = c.ShouldBindJSON(&taskFlowQuery)
	code = model.ApproveTask(taskFlowQuery.UID, taskFlowQuery.Status, taskFlowQuery.EmployeeUID)
	msg.Message(c, code, nil)
}
