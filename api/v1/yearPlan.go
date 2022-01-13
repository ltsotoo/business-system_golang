package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func StartYearPlan(c *gin.Context) {
	code = model.StartYearPlan()
	msg.Message(c, code, nil)
}

func EndYearPlan(c *gin.Context) {
	code = model.EndYearPlan()
	msg.Message(c, code, nil)
}

func EditYearOffice(c *gin.Context) {
	var office model.Office
	_ = c.ShouldBindJSON(&office)

	code = model.UpdateYearOffice(&office)
	msg.Message(c, code, nil)
}
