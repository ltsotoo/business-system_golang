package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/rbac"
	"strconv"

	"github.com/gin-gonic/gin"
)

//录入客户
func EntryCustomer(c *gin.Context) {
	code = rbac.Check(c, "customer", "insert")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var customer model.Customer
		_ = c.ShouldBindJSON(&customer)
		code = model.InsertCustomer(&customer)
		msg.Message(c, code, customer)
	}
}

//删除客户
func DelCustomer(c *gin.Context) {
	code = rbac.Check(c, "customer", "delete")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		uid := c.Param("uid")
		code = model.DeleteCustomer(uid)
		msg.Message(c, code, nil)
	}
}

//编辑客户
func EditCustomer(c *gin.Context) {
	code = rbac.Check(c, "customer", "update")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var customer model.Customer
		_ = c.ShouldBindJSON(&customer)
		code = model.UpdateCustomer(&customer)
		msg.Message(c, code, customer)
	}
}

//查询客户
func QueryCustomer(c *gin.Context) {
	code = rbac.Check(c, "customer", "select")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var customer model.Customer
		uid := c.Param("uid")
		customer, code = model.SelectCustomer(uid)
		msg.Message(c, code, customer)
	}
}

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

	customers, code, total = model.SelectCustomers(pageSize, pageNo, &customerQuery)
	msg.MessageForList(c, msg.SUCCESS, customers, pageSize, pageNo, total)
}

//查询客户公司列表
func QueryCompanys(c *gin.Context) {
	var companys []model.CustomerCompany
	companys, code = model.SelectCompanys("")
	msg.Message(c, code, companys)
}
