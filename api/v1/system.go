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

	code = model.CheckLogin(&employee)

	if code == msg.SUCCESS {
		token, _ = middleware.SetToken(employee.ID, employee.Phone)
		token = "Bearer " + token
	}

	msg.Message(c, code, gin.H{"employee": employee, "token": token})
}
