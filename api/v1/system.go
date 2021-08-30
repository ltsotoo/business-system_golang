package v1

import (
	"business-system_golang/middleware"
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var employee model.Employee
	var token string
	_ = c.ShouldBindJSON(&employee)

	employee, code = model.CheckLogin(employee.Phone, employee.Password)
	if code == msg.SUCCESS {
		permission_uid_set := model.SelectAllPermission(employee.UID, employee.DepartmentUID)
		token, _ = middleware.SetToken(employee.UID, permission_uid_set)
		token = "Bearer " + token
	}
	msg.Message(c, code, gin.H{"employee": employee, "token": token})
}

func Regist(c *gin.Context) {
	var employee model.Employee
	_ = c.ShouldBindJSON(&employee)
	code = model.CheckEmployee(employee.Phone)
	if code == msg.ERROR_EMPLOYEE_NOT_EXIST {
		code = model.InsertEmployee(&employee)
	}
	msg.Message(c, code, employee)
}
