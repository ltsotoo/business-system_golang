package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryHistoryOffices(c *gin.Context) {
	var historyOffices []model.HistoryOffice
	var historyOffice model.HistoryOffice
	var total int64

	_ = c.ShouldBindJSON(&historyOffice)

	pageSize, pageSizeErr := strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	pageNo, pageNoErr := strconv.Atoi(c.DefaultQuery("pageNo", "0"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	historyOffices, code, total = model.SelectHistoryOffices(pageSize, pageNo, &historyOffice)
	msg.MessageForList(c, code, historyOffices, pageSize, pageNo, total)
}

func QueryHistoryEmployees(c *gin.Context) {
	var historyEmployees []model.HistoryEmployee
	var historyEmployee model.HistoryEmployee
	var total int64

	_ = c.ShouldBindJSON(&historyEmployee)

	pageSize, pageSizeErr := strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	pageNo, pageNoErr := strconv.Atoi(c.DefaultQuery("pageNo", "0"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	historyEmployees, code, total = model.SelectHistoryEmployees(pageSize, pageNo, &historyEmployee)
	msg.MessageForList(c, code, historyEmployees, pageSize, pageNo, total)
}
