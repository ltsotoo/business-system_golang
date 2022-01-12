package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func EditMonthPlan(c *gin.Context) {
	var monthPlan model.MonthPlan
	_ = c.ShouldBindJSON(&monthPlan)

	code = model.UpdateMonthPlan(&monthPlan)

	msg.Message(c, code, monthPlan)
}

func QueryMonthPlans(c *gin.Context) {
	var monthPlans []model.MonthPlan
	monthPlans, code = model.SelectMonthPlans()
	msg.Message(c, code, monthPlans)
}
