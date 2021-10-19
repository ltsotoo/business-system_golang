package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//查询我的任务
func QueryMyTasks(c *gin.Context) {
	var tasks []model.Task
	var task model.Task
	var total int64

	_ = c.ShouldBindJSON(&task)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	tasks, code, total = model.SelectMyTasks(pageSize, pageNo, &task, c.MustGet("employeeUID").(string))
	msg.MessageForList(c, code, tasks, pageSize, pageNo, total)
}

//查询我的报销
func QueryMyExpenses(c *gin.Context) {
	var expenses []model.Expense
	var expense model.Expense
	var total int64

	_ = c.ShouldBindJSON(&expense)
	expense.EmployeeUID = c.MustGet("employeeUID").(string)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	expenses, code, total = model.SelectMyExpenses(pageSize, pageNo, &expense)
	msg.MessageForList(c, code, expenses, pageSize, pageNo, total)
}
