package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProcurementPlan(c *gin.Context) {
	var procurementPlan model.ProcurementPlan
	_ = c.ShouldBindJSON(&procurementPlan)
	procurementPlan.EmployeeUID = c.MustGet("employeeUID").(string)
	code = model.InsertProcurementPlan(&procurementPlan)
	msg.Message(c, code, procurementPlan)
}

func EditProcurementPlan(c *gin.Context) {
	var procurementPlan model.ProcurementPlan
	_ = c.ShouldBindJSON(&procurementPlan)
	code = model.UpdateProcurementPlan(&procurementPlan)
	msg.Message(c, code, procurementPlan)
}

func QueryProcurementPlan(c *gin.Context) {
	var procurementPlan model.ProcurementPlan
	uid := c.Param("uid")
	procurementPlan, code = model.SelectProcurementPlan(uid)
	msg.Message(c, code, procurementPlan)
}

func QueryProcurementPlans(c *gin.Context) {
	var procurementPlans []model.ProcurementPlan
	var total int64
	var procurementPlan model.ProcurementPlan

	_ = c.ShouldBindJSON(&procurementPlan)

	pageSize, pageSizeErr := strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	pageNo, pageNoErr := strconv.Atoi(c.DefaultQuery("pageNo", "0"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	procurementPlans, code, total = model.SelectProcurementPlans(pageSize, pageNo, &procurementPlan)
	msg.MessageForList(c, code, procurementPlans, pageSize, pageNo, total)
}
