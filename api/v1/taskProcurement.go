package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

//发起采购
func AddTaskProcurement(c *gin.Context) {
	var taskProcurement model.TaskProcurement
	_ = c.ShouldBindJSON(&taskProcurement)
	taskProcurement.EmployeeUID = c.MustGet("employeeUID").(string)
	code = model.InsertTaskProcurements(&taskProcurement)
	msg.Message(c, code, taskProcurement)
}

func NextTaskProcurement(c *gin.Context) {
	var taskProcurement model.TaskProcurement
	_ = c.ShouldBindJSON(&taskProcurement)
	if taskProcurement.UID != "" && taskProcurement.Status != 0 {
		code = model.UpdateTaskProcurement(taskProcurement.UID, taskProcurement.Status)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

//查询我发起的采购任务
func QueryMyApplicationTaskProcurements(c *gin.Context) {
	var taskProcurements []model.TaskProcurement
	taskProcurements, code = model.SelectMyApplicationTaskProcurements(c.MustGet("employeeUID").(string))
	msg.Message(c, code, taskProcurements)
}

//查询我的采购任务
func QueryMyTaskProcurements(c *gin.Context) {
	var taskProcurements []model.TaskProcurement
	taskProcurements, code = model.SelectMyTaskProcurements(c.MustGet("employeeUID").(string))
	msg.Message(c, code, taskProcurements)
}
