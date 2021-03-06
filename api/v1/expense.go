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

	canAdd := true

	if expense.Type == magic.EXPENSE_TYPE_1 {
		var employee model.Employee
		employee, code = model.SelectEmployee(expense.EmployeeUID)
		if employee.ID == 0 || employee.Money < expense.Amount {
			canAdd = false
			code = msg.ERROR_EXPENSE_MONEY_LESS
		}
	} else if expense.Type == magic.EXPENSE_TYPE_2 {
		var employee model.Employee
		employee, code = model.SelectEmployee(expense.EmployeeUID)
		if employee.ID == 0 || employee.Office.Money < expense.Amount {
			canAdd = false
			code = msg.ERROR_EXPENSE_MONEY_LESS
		}
	}

	if canAdd {
		expense.Status = magic.EXPENSE_STATUS_NOT_APPROVAL_1
		code = model.InsertExpense(&expense)
	}

	msg.Message(c, code, expense)
}

//删除报销
func DelExpense(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteExpense(uid)
	msg.Message(c, code, nil)
}

//审核报销
func ApprovalExpense(c *gin.Context) {

	if model.SystemSettlement {
		msg.MessageForSystemSettlement(c)
		return
	}

	var expense, expenseDB model.Expense
	_ = c.ShouldBindJSON(&expense)
	expenseDB, code = model.SelectExpense(expense.UID)

	if code == msg.SUCCESS && expense.Status == expenseDB.Status &&
		expense.Status > 0 && expense.Status < 4 {
		var maps = make(map[string]interface{})
		oldStatus := expense.Status

		switch expense.Status {
		case magic.EXPENSE_STATUS_NOT_APPROVAL_1:
			maps["approver_uid1"] = c.MustGet("employeeUID").(string)
		case magic.EXPENSE_STATUS_NOT_APPROVAL_2:
			maps["approver_uid2"] = c.MustGet("employeeUID").(string)
		case magic.EXPENSE_STATUS_NOT_PAYMENT:
			maps["approver_uid3"] = c.MustGet("employeeUID").(string)
		}

		if expense.IsPass {
			expense.Status = expense.Status + 1
		} else {
			expense.Status = -1
		}
		maps["status"] = expense.Status
		code = model.ApprovalExpense(oldStatus, &expense, maps, c.MustGet("employeeUID").(string), expenseDB)
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

	pageSize, pageSizeErr := strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	pageNo, pageNoErr := strconv.Atoi(c.DefaultQuery("pageNo", "0"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	expenses, code, total = model.SelectExpenses(pageSize, pageNo, &expenseQuery)
	msg.MessageForList(c, msg.SUCCESS, expenses, pageSize, pageNo, total)
}
