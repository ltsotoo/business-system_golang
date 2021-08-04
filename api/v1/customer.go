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
	var customerQuery model.CustomerQuery

	_ = c.ShouldBindJSON(&customerQuery)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize <= 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo <= 0 {
		pageNo = 1
	}

	customers, code, total = model.SelectCustomers(pageSize, pageNo, customerQuery)
	msg.MessageForList(c, msg.SUCCESS, customers, pageSize, pageNo, total)
}

//查询客户公司列表
func QueryCompanys(c *gin.Context) {
	var companys []model.Company
	companys, code = model.SelectCompanys()
	msg.Message(c, code, companys)
}

//查询客户公司课题组列表
func QueryResearchGroupsByCompanyID(c *gin.Context) {
	var researchGroups []model.ResearchGroup
	companyID, companyIDErr := strconv.Atoi(c.Query("companyID"))
	if companyIDErr == nil {
		researchGroups, code = model.SelectResearchGroupsByCompanyID(companyID)
	}
	msg.Message(c, code, researchGroups)
}

//通过公司ID和课题组ID查询客户
func QueryCustomersByCompanyIDAndResearchGroupID(c *gin.Context) {
	var customers []model.Customer
	companyID, companyIDErr := strconv.Atoi(c.DefaultQuery("companyID", "0"))
	researchGroupID, researchGroupIDErr := strconv.Atoi(c.DefaultQuery("researchGroupID", "0"))

	if companyIDErr == nil || researchGroupIDErr == nil {
		customers, code = model.SelectCustomersByCompanyIDAndResearchGroupID(companyID, researchGroupID)
	}

	msg.Message(c, code, customers)
}
