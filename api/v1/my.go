package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func EditMyPwd(c *gin.Context) {
	var employeeQuery model.EmployeeQuery
	_ = c.ShouldBindJSON(&employeeQuery)

	code = model.UpdatePwd(&employeeQuery)
	msg.Message(c, code, nil)
}
