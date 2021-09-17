package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/rbac"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

//录入员工
func EntryEmployee(c *gin.Context) {
	code = rbac.Check(c, "employee", "insert")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var employee model.Employee
		_ = c.ShouldBindJSON(&employee)
		code = model.CheckEmployee(employee.Phone)
		if code == msg.ERROR_EMPLOYEE_NOT_EXIST {
			//对录入员工的办事处和部门信息补充
			if employee.OfficeUID == "" || employee.DepartmentUID == "" {
				visitor, _ := model.SelectEmployee(c.MustGet("employeeUID").(string))
				if employee.OfficeUID == "" {
					employee.OfficeUID = visitor.OfficeUID
				}
				if employee.DepartmentUID == "" {
					employee.DepartmentUID = visitor.DepartmentUID
				}
			}
			//默认密码等于手机号
			employee.Password = employee.Phone
			code = model.InsertEmployee(&employee)
		}
		msg.Message(c, code, employee)
	}
}

//删除员工
func DelEmployee(c *gin.Context) {
	code = rbac.Check(c, "employee", "delete")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		uid := c.Param("uid")
		code = model.DeleteEmployee(uid)
		msg.Message(c, code, nil)
	}
}

//编辑员工信息
func EditEmployee(c *gin.Context) {
	code = rbac.Check(c, "employee", "update")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var employee model.Employee
		_ = c.ShouldBindJSON(&employee)
		code = model.UpdateEmployee(&employee)
		msg.Message(c, code, employee)
	}
}

//查询员工
func QueryEmployee(c *gin.Context) {
	var employee model.Employee
	uid := c.Param("uid")
	if uid == "my" {
		uid = c.MustGet("employeeUID").(string)
	}
	employee, code = model.SelectEmployee(uid)
	msg.Message(c, code, employee)
}

//查询员工列表
func QueryEmployees(c *gin.Context) {
	var employees []model.Employee
	var total int64
	var employeeQuery model.EmployeeQuery

	_ = c.ShouldBindJSON(&employeeQuery)

	if employeeQuery.OfficeUID == "" && employeeQuery.DepartmentUID == "" {
		visitor, _ := model.SelectEmployee(c.MustGet("employeeUID").(string))
		employeeQuery.OfficeUID = visitor.OfficeUID
		employeeQuery.DepartmentUID = visitor.DepartmentUID
	}

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	employees, code, total = model.SelectEmployees(pageSize, pageNo, &employeeQuery)
	msg.MessageForList(c, msg.SUCCESS, employees, pageSize, pageNo, total)
}
