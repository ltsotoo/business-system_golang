package routes

import (
	v1 "business-system_golang/api/v1"
	"business-system_golang/config"
	"business-system_golang/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.SystemConfig.Server.Mode)
	// r := gin.Default()
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Cors())

	router := r.Group("api/v1")
	{
		//SYSTEM接口
		router.POST("login", v1.Login)
		router.POST("regist", v1.Regist)
	}

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//字典表接口
		auth.POST("dictionary", v1.AddDictionary)
		auth.DELETE("dictionary/:uid", v1.DelDictionary)
		auth.PUT("dictionary", v1.EditDictionary)
		auth.GET("dictionaries", v1.QueryDictionaries)
		auth.GET("dictionarieTypes", v1.QueryDictionarieTypes)
		//员工模块接口
		auth.POST("employee", v1.EntryEmployee)
		auth.DELETE("employee/:uid", v1.DelEmployee)
		auth.PUT("employee", v1.EditEmployee)
		auth.GET("employee/:uid", v1.QueryEmployee)
		auth.POST("employees", v1.QueryEmployees)
		auth.GET("my", v1.QueryMy)
		auth.GET("resetPwd/:uid", v1.ResetEmployeePwd)
		auth.POST("editMyPwd", v1.EditMyPwd)
		//员工模块接口PLUS
		auth.POST("office", v1.EntryOffice)
		auth.DELETE("office/:uid", v1.DelOffice)
		auth.PUT("office", v1.EditOffice)
		auth.POST("offices", v1.QueryOffices)
		auth.POST("department", v1.EntryDepartment)
		auth.DELETE("department/:uid", v1.DelDepartment)
		auth.PUT("department", v1.EditDepartment)
		auth.POST("departments", v1.QueryDepartments)
		auth.POST("role", v1.AddRole)
		auth.PUT("role", v1.EditRole)
		auth.GET("role/:uid", v1.QueryRole)
		auth.GET("allRoles", v1.QueryAllRoles)
		auth.GET("roles", v1.QueryRoles)
		auth.GET("permissions", v1.QueryPermissions)
		//合同模块接口
		auth.POST("contract", v1.EntryContract)
		auth.DELETE("contract/:uid", v1.DelContract)
		auth.GET("contract/:uid", v1.QueryContract)
		auth.POST("contracts", v1.QueryContracts)
		auth.POST("rejectContract", v1.RejectContract)
		//合同流程模块接口
		auth.PUT("task/contract/approve", v1.ApproveContract)
		//任务模块接口
		auth.POST("tasks", v1.QueryTasks)
		auth.GET("taskRemarks", v1.QueryTaskRemarks)
		auth.POST("mytasks", v1.QueryMyTasks)
		//任务流程模块接口
		auth.PUT("task/flow/approve", v1.ApproveTask)
		auth.PUT("task/flow/next", v1.NextTask)
		//回款模块接口
		auth.POST("payment", v1.AddPayment)
		auth.PUT("payment", v1.EditPayment)
		auth.GET("payments/:contractUID", v1.QueryPaymentsForContract)
		auth.POST("finishPayments", v1.FinishPayments)
		//客户模块接口
		auth.POST("customer", v1.EntryCustomer)
		auth.DELETE("customer/:uid", v1.DelCustomer)
		auth.PUT("customer", v1.EditCustomer)
		auth.GET("customer/:uid", v1.QueryCustomer)
		auth.POST("customers", v1.QueryCustomers)
		//客户模块接口PLUS
		auth.POST("company", v1.AddCustomerCompany)
		auth.DELETE("company/:uid", v1.DelCustomerCompany)
		auth.PUT("company", v1.EditCustomerCompany)
		auth.POST("companys", v1.QueryCustomerCompanys)
		//产品模块接口
		auth.POST("product", v1.EntryProduct)
		auth.DELETE("product/:uid", v1.DelProduct)
		auth.PUT("productBase", v1.EditProductBase)
		auth.PUT("productPrice", v1.EditProductPrice)
		auth.PUT("productNumber", v1.EditProductNumber)
		auth.GET("product/:uid", v1.QueryProduct)
		auth.POST("products", v1.QueryProducts)
		//
		auth.POST("productType", v1.AddProductType)
		auth.DELETE("productType/:uid", v1.DelProductType)
		auth.PUT("productType", v1.EditProductType)
		auth.POST("productTypes", v1.QueryProductTypes)
		//供应商模块接口
		auth.POST("supplier", v1.EntrySupplier)
		auth.DELETE("supplier/:uid", v1.DelSupplier)
		auth.PUT("supplier", v1.EditSupplier)
		auth.GET("supplier/:uid", v1.QuerySupplier)
		auth.POST("suppliers", v1.QuerySuppliers)
		//财务模块接口
		auth.POST("expense", v1.AddExpense)
		auth.PUT("expense", v1.ApprovalExpense)
		auth.GET("expense/:uid", v1.QueryExpense)
		auth.POST("expenses", v1.QueryExpenses)
		auth.POST("myexpenses", v1.QueryMyExpenses)
		//预研究模块
		auth.POST("preResearch", v1.CreatePreResearch)
		auth.DELETE("preResearch/:uid", v1.DelPreResearch)
		auth.GET("preResearch/:uid", v1.QueryPreResearch)
		auth.POST("preResearchs", v1.QueryPreResearchs)
		auth.GET("preResearchTask/:uid", v1.QueryPreResearchTask)
		auth.POST("preResearchTasks", v1.QueryPreResearchTasks)
		// auth.PUT("preResearch", v1.EditPreResearch)
		auth.PUT("preResearch", v1.UpdatePreResearch)
		auth.PUT("preResearchTask", v1.UpdatePreResearchTask)
		//投标保证金模块
		auth.POST("bidBond", v1.AddBidBond)
		auth.DELETE("bidBond/:uid", v1.DelBidBond)
		auth.PUT("bidBond", v1.EditBidBond)
		auth.PUT("bidBond/approve/:uid", v1.ApproveBidBond)
		auth.POST("bidBonds", v1.QueryBidBonds)
		//合同开票
		auth.POST("invoice", v1.AddInvoice)
		auth.DELETE("invoice/:uid", v1.DelInvoice)
		auth.PUT("invoice", v1.EditInvoice)
		auth.PUT("invoice/approve/:uid", v1.ApproveInvoice)
		auth.POST("invoices", v1.QueryInvoices)
	}

	_ = r.Run(config.SystemConfig.Server.Port)
}
