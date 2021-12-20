package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EditMyPwd(c *gin.Context) {
	var employeeQuery model.EmployeeQuery
	_ = c.ShouldBindJSON(&employeeQuery)

	code = model.UpdatePwd(&employeeQuery)
	msg.Message(c, code, nil)
}

func QueryMy(c *gin.Context) {
	var employee model.Employee
	uid := c.MustGet("employeeUID").(string)
	employee, code = model.SelectEmployee(uid)
	msg.Message(c, code, employee)
}

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
	var expenseQuery model.ExpenseQuery
	var total int64

	_ = c.ShouldBindJSON(&expenseQuery)
	expenseQuery.EmployeeUID = c.MustGet("employeeUID").(string)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	expenses, code, total = model.SelectExpenses(pageSize, pageNo, &expenseQuery)
	msg.MessageForList(c, code, expenses, pageSize, pageNo, total)
}
