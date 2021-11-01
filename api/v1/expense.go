package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/magic"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//发起报销
func AddExpense(c *gin.Context) {
	var expense model.Expense
	_ = c.ShouldBindJSON(&expense)
	expense.EmployeeUID = c.MustGet("employeeUID").(string)
	expense.Status = magic.EXPENSE_STATUS_NOT_APPROVAL
	code = model.InsertExpense(&expense)
	msg.Message(c, code, expense)
}

//审核报销
func ApprovalExpense(c *gin.Context) {
	var expense, expenseDB model.Expense
	_ = c.ShouldBindJSON(&expense)
	expenseDB, code = model.SelectExpense(expense.UID)
	if code == msg.SUCCESS && expenseDB.Status == magic.EXPENSE_STATUS_NOT_APPROVAL &&
		(expense.Status == magic.EXPENSE_STATUS_FAIL || expense.Status == magic.EXPENSE_STATUS_PASS) {
		expense.ApproverUID = c.MustGet("employeeUID").(string)
		code = model.UpdateMoneyExpense(&expense)
	}
	msg.Message(c, code, expense)
}

//查看报销
func QueryExpense(c *gin.Context) {
	var expense model.Expense
	uid := c.Param("uid")
	expense, code = model.SelectExpense(uid)
	msg.Message(c, code, expense)
}

//查看报销列表
func QueryExpenses(c *gin.Context) {
	var expenses []model.Expense
	var total int64
	var expenseQuery model.ExpenseQuery

	_ = c.ShouldBindJSON(&expenseQuery)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	expenses, code, total = model.SelectExpenses(pageSize, pageNo, &expenseQuery)
	msg.MessageForList(c, msg.SUCCESS, expenses, pageSize, pageNo, total)
}
