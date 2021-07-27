package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//录入客户
func EntryCustomer(c *gin.Context) {
	var customer model.Customer
	_ = c.ShouldBindJSON(&customer)
	code = model.CreateCustomer(&customer)
	msg.Message(c, code, customer)
}

//删除客户
func DelCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCustomer(id)
	msg.Message(c, code, nil)
}

//编辑客户
func EditCustomer(c *gin.Context) {
	var customer model.Customer
	_ = c.ShouldBindJSON(&customer)
	code = model.UpdateCustomer(&customer)
	msg.Message(c, code, customer)
}

//查询客户
func QueryCustomer(c *gin.Context) {
	var customer model.Customer
	id, _ := strconv.Atoi(c.Param("id"))
	customer, code = model.SelectCustomer(id)
	msg.Message(c, code, customer)
}

//查询客户列表
func QueryCustomers(c *gin.Context) {
	var customers []model.Customer
	var total int64

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize <= 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo <= 0 {
		pageNo = 1
	}

	customers, code, total = model.SelectCustomers(pageSize, pageNo)
	msg.MessageForList(c, msg.SUCCESS, customers, pageSize, pageNo, total)
}
