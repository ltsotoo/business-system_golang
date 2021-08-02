package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var employee model.Employee
	_ = c.ShouldBindJSON(&employee)

	code = model.CheckEmployeeByPhoneAndPwd(&employee)

	msg.Message(c, code, employee)
}
