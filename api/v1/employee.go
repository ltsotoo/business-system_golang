package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

//录入员工
func EntryEmployee(c *gin.Context) {
	var employee model.Employee
	_ = c.ShouldBindJSON(&employee)
	code = model.CheckEmployeePhone(employee.Phone)
	if code == msg.ERROR_EMPLOYEE_NOT_EXIST {
		model.CreateEmployee(&employee)
		code = msg.SUCCESS
	}
	msg.Message(c, code, employee)
}

//删除员工
func DelEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteEmployee(id)
	msg.Message(c, code, nil)
}

//编辑员工信息
func EditEmployee(c *gin.Context) {
	var employee model.Employee
	c.ShouldBindJSON(&employee)
	code = model.UpdateEmployee(&employee)
	msg.Message(c, code, employee)
}

//查询员工
func QueryEmployee(c *gin.Context) {
	var employee model.Employee
	id, _ := strconv.Atoi(c.Param("id"))
	employee, code = model.SelectEmployee(id)
	msg.Message(c, code, employee)
}

//查询员工列表
func QueryEmployees(c *gin.Context) {

	var employees []model.Employee
	var total int64

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize <= 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo <= 0 {
		pageNo = 1
	}
	employees, code, total = model.SelectEmployees(pageSize, pageNo)
	msg.MessageForList(c, msg.SUCCESS, employees, pageSize, pageNo, total)
}
