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
	code = model.InsertCustomer(&customer)
	msg.Message(c, code, customer)
}

//删除客户
func DelCustomer(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteCustomer(uid)
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
	uid := c.Param("uid")
	customer, code = model.SelectCustomer(uid)
	msg.Message(c, code, customer)
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
	msg.MessageForList(c, code, customers, pageSize, pageNo, total)
}

func AddCustomerCompany(c *gin.Context) {
	var customerCompany model.CustomerCompany
	_ = c.ShouldBindJSON(&customerCompany)
	code = model.InsertCustomerCompany(&customerCompany)
	msg.Message(c, code, customerCompany)
}

func DelCustomerCompany(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteCustomerCompany(uid)
	msg.Message(c, code, nil)
}

func EditCustomerCompany(c *gin.Context) {
	var customerCompany model.CustomerCompany
	_ = c.ShouldBindJSON(&customerCompany)
	code = model.UpdateCustomerCompany(&customerCompany)
	msg.Message(c, code, customerCompany)
}

//查询客户公司列表
func QueryCustomerCompanys(c *gin.Context) {
	var companys []model.CustomerCompany
	var total int64
	var company model.CustomerCompany

	_ = c.ShouldBindJSON(&company)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	companys, code, total = model.SelectCustomerCompanys(pageSize, pageNo, &company)
	msg.MessageForList(c, code, companys, pageSize, pageNo, total)
}
