package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/magic"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

//查询我的技术任务
func QueryMyTasks1(c *gin.Context) {
	var tasks []model.Task
	tasks, code = model.SelectMyTasks(c.MustGet("employeeUID").(string), magic.TASK_STATUS_NOT_DESIGN)
	msg.Message(c, code, tasks)
}

//查询我的采购任务
func QueryMyTasks2(c *gin.Context) {
	var tasks []model.Task
	tasks, code = model.SelectMyTasks(c.MustGet("employeeUID").(string), magic.TASK_STATUS_NOT_PURCHASE)
	msg.Message(c, code, tasks)
}

//查询我的仓库任务
func QueryMyTasks3(c *gin.Context) {
	var tasks []model.Task
	tasks, code = model.SelectMyTasks(c.MustGet("employeeUID").(string), magic.TASK_STATUS_NOT_STORAGE)
	msg.Message(c, code, tasks)
}

//查询我的发货任务
func QueryMyTasks4(c *gin.Context) {
	var tasks []model.Task
	tasks, code = model.SelectMyTasks(c.MustGet("employeeUID").(string), magic.TASK_STATUS_NOT_SHIPMENT)
	msg.Message(c, code, tasks)
}

//查询我的确认收货任务
func QueryMyTasks5(c *gin.Context) {
	var tasks []model.Task
	tasks, code = model.SelectMyTasks(c.MustGet("employeeUID").(string), magic.TASK_STATUS_NOT_CONFIRM)
	msg.Message(c, code, tasks)
}

func QueryMyExpense(c *gin.Context) {

}
