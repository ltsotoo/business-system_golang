package v2

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

//查询客户列表
func QueryCustomers(c *gin.Context) {
	var customers []model.Customer
	var total int64
	var customerQuery model.CustomerQuery

	_ = c.ShouldBindJSON(&customerQuery)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	customers, code, total = model.SelectCustomers(pageSize, pageNo, customerQuery)
	msg.MessageForList(c, msg.SUCCESS, customers, pageSize, pageNo, total)
}
