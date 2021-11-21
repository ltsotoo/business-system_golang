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
	task, _ := model.SelectTask(taskFlowQuery.UID)
	if task.UID != "" {
		taskFlowQuery.ContractUID = task.ContractUID
		code = model.ApproveTask(&taskFlowQuery)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func NextTask(c *gin.Context) {
	var task, dbTask model.Task
	_ = c.ShouldBindJSON(&task)
	dbTask, code = model.SelectTask(task.UID)
	if code == msg.SUCCESS && task.Status == dbTask.Status {
		lastStatus := task.Status
		from, to := "", ""
		switch task.Status {
		case magic.TASK_STATUS_NOT_DESIGN:
			task.Status = magic.TASK_STATUS_NOT_PURCHASE
			from = dbTask.TechnicianManUID
			to = dbTask.PurchaseManUID
		case magic.TASK_STATUS_NOT_PURCHASE:
			task.Status = magic.TASK_STATUS_NOT_STORAGE
			from = dbTask.PurchaseManUID
			to = dbTask.InventoryManUID
		case magic.TASK_STATUS_NOT_STORAGE:
			from = dbTask.InventoryManUID
			if task.Type == magic.TASK_TYPE_3 {
				task.Status = magic.TASK_STATUS_NOT_ASSEMBLY
				to = dbTask.TechnicianManUID
			} else {
				task.Status = magic.TASK_STATUS_NOT_SHIPMENT
				to = dbTask.ShipmentManUID
			}
		case magic.TASK_STATUS_NOT_ASSEMBLY:
			task.Status = magic.TASK_STATUS_NOT_SHIPMENT
			from = dbTask.TechnicianManUID
			to = dbTask.ShipmentManUID
		case magic.TASK_STATUS_NOT_SHIPMENT:
			task.Status = magic.TASK_STATUS_SHIPMENT
			from = dbTask.ShipmentManUID
			to = dbTask.ShipmentManUID
		}
		code = model.NextTaskStatus(task.UID, lastStatus, from, to, task.Status, task.CurrentRemarksText)
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
	employeeUID := c.MustGet("employeeUID").(string)
	taskRemarksList, code = model.SelectTaskRemarks(taskUID, employeeUID)
	msg.Message(c, code, taskRemarksList)
}
