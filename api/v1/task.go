package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/magic"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func DelTask(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteTask(uid)
	msg.Message(c, code, nil)
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

func NextTask(c *gin.Context) {
	var task, dbTask model.Task
	_ = c.ShouldBindJSON(&task)
	dbTask, code = model.SelectTask(task.UID)
	if code == msg.SUCCESS && task.Status == dbTask.Status {
		from := task.Status
		switch task.Status {
		case magic.TASK_STATUS_NOT_DESIGN:
			task.Status = magic.TASK_STATUS_NOT_PURCHASE
		case magic.TASK_STATUS_NOT_PURCHASE:
			task.Status = magic.TASK_STATUS_NOT_STORAGE
		case magic.TASK_STATUS_NOT_STORAGE:
			if task.Type == magic.TASK_TYPE_3 {
				task.Status = magic.TASK_STATUS_NOT_ASSEMBLY
			} else {
				task.Status = magic.TASK_STATUS_NOT_SHIPMENT
			}
		case magic.TASK_STATUS_NOT_ASSEMBLY:
			task.Status = magic.TASK_STATUS_NOT_SHIPMENT
		case magic.TASK_STATUS_NOT_SHIPMENT:
			task.Status = magic.TASK_STATUS_SHIPMENT
		}
		code = model.NextTaskStatus(task.UID, from, task.Status, task.CurrentRemarksText)
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
			if task.Status != magic.TASK_STATUS_SHIPMENT {
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

func QueryTaskRemarks(c *gin.Context) {
	var taskRemarksList []model.TaskRemarks
	taskUID := c.Query("taskUID")
	taskRemarksList, code = model.SelectTaskRemarks(taskUID)
	msg.Message(c, code, taskRemarksList)
}
