package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/magic"
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
	code = model.ApproveTask(&taskFlowQuery)
	msg.Message(c, code, nil)
}

func LastTask(c *gin.Context) {
	var task, dbTask model.Task
	_ = c.ShouldBindJSON(&task)
	dbTask, code = model.SelectTask(task.UID)
	if code == msg.SUCCESS && task.Status == dbTask.Status && task.Status >= 2 && task.Status <= 4 {
		task.Status = dbTask.Status - 1
		code = model.NextTaskStatus(task.UID, task.Status, task.NextRemarks)
		msg.Message(c, code, nil)
	} else {
		msg.Message(c, code, nil)
	}
}

func NextTask(c *gin.Context) {
	var task, dbTask model.Task
	_ = c.ShouldBindJSON(&task)
	dbTask, code = model.SelectTask(task.UID)
	if code == msg.SUCCESS && task.Status == dbTask.Status {
		task.Status = dbTask.Status + 1
		code = model.NextTaskStatus(task.UID, task.Status, task.NextRemarks)
		if code == msg.SUCCESS {
			code = checkTasksUpdateContract(task.ContractUID)
		}
		msg.Message(c, code, nil)
	} else {
		msg.Message(c, code, nil)
	}
}

func checkTasksUpdateContract(contractUID string) int {
	tasks, _ := model.SelectTasksByContractUID(contractUID)
	var finsh = true
	if len(tasks) != 0 {
		for _, task := range tasks {
			if task.Status != magic.TASK_STATUS_FINISH {
				finsh = false
				break
			}
		}
	}
	if finsh {
		code = model.UpdateContractProductionStatusToFinish(contractUID)
	}
	return code
}
