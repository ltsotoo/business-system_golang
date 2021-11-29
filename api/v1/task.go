package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/magic"
	"business-system_golang/utils/msg"
	"time"

	"github.com/gin-gonic/gin"
)

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
	taskFlowQuery.ContractUID = task.ContractUID
	code = model.ApproveTask(&taskFlowQuery)
	msg.Message(c, code, nil)
}

func NextTask(c *gin.Context) {
	var taskFlowQuery model.TaskFlowQuery
	var dbTask model.Task
	_ = c.ShouldBindJSON(&taskFlowQuery)
	dbTask, code = model.SelectTask(taskFlowQuery.UID)

	if code == msg.SUCCESS && taskFlowQuery.Status == dbTask.Status {
		var maps = make(map[string]interface{})
		t := time.Now().Format("2006-01-02 15:04:05")
		from, to := "", ""
		switch dbTask.Status {
		case magic.TASK_STATUS_NOT_DESIGN:
			maps["status"] = magic.TASK_STATUS_NOT_PURCHASE
			from = dbTask.TechnicianManUID
			to = dbTask.PurchaseManUID
			maps["technician_real_end_date"] = t
			maps["purchase_start_date"] = t
		case magic.TASK_STATUS_NOT_PURCHASE:
			maps["status"] = magic.TASK_STATUS_NOT_STORAGE
			from = dbTask.PurchaseManUID
			to = dbTask.InventoryManUID
			maps["purchase_real_end_date"] = t
			maps["inventory_start_date"] = t
		case magic.TASK_STATUS_NOT_STORAGE:
			from = dbTask.InventoryManUID
			if dbTask.Type == magic.TASK_TYPE_3 {
				maps["status"] = magic.TASK_STATUS_NOT_ASSEMBLY
				to = dbTask.TechnicianManUID
			} else {
				maps["status"] = magic.TASK_STATUS_NOT_SHIPMENT
				maps["shipment_start_date"] = t
				to = dbTask.ShipmentManUID
			}
		case magic.TASK_STATUS_NOT_ASSEMBLY:
			maps["status"] = magic.TASK_STATUS_NOT_SHIPMENT
			maps["shipment_start_date"] = t
			from = dbTask.TechnicianManUID
			to = dbTask.ShipmentManUID
		case magic.TASK_STATUS_NOT_SHIPMENT:
			maps["status"] = magic.TASK_STATUS_SHIPMENT
			from = dbTask.ShipmentManUID
			to = dbTask.ShipmentManUID
		}
		code = model.NextTaskStatus(dbTask.UID, dbTask.Status, from, to, maps, taskFlowQuery.CurrentRemarksText)
		if code == msg.SUCCESS {
			code = checkTasksUpdateContract(dbTask.ContractUID)
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
