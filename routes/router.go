package routes

import (
	v1 "business-system_golang/api/v1"
	v2 "business-system_golang/api/v2"
	"business-system_golang/config"
	"business-system_golang/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.SystemConfig.Server.Mode)
	r := gin.Default()

	r.Use(middleware.Cors())

	V2 := r.Group("api/v2")
	{
		//SYSTEM接口
		V2.POST("customers", v2.QueryCustomers)
		V2.POST("dictionaryType", v2.AddDictionaryType)
	}

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
		auth.GET("dictionaryType", v1.QueryDictionaryType)
		auth.GET("dictionaryTypes", v1.QueryDictionaryTypes)
		auth.GET("dictionaries", v1.QueryDictionaries)
		//员工模块接口
		auth.POST("employee", v1.EntryEmployee)
		auth.DELETE("employee/:uid", v1.DelEmployee)
		auth.PUT("employee", v1.EditEmployee)
		auth.GET("employee/:uid", v1.QueryEmployee)
		auth.POST("employees", v1.QueryEmployees)
		//员工模块接口PLUS
		auth.POST("office", v1.EntryOffice)
		auth.DELETE("office/:uid", v1.DelOffice)
		auth.GET("offices", v1.QueryOffices)
		auth.POST("area", v1.EntryArea)
		auth.DELETE("area/:uid", v1.DelArea)
		auth.PUT("area", v1.EditArea)
		auth.POST("areas", v1.QueryAreas)
		auth.POST("department", v1.EntryDepartment)
		auth.DELETE("department/:uid", v1.DelDepartment)
		auth.POST("departments", v1.QueryDepartments)
		auth.GET("roles", v1.QueryRoles)
		auth.GET("permissions", v1.QueryPermissions)
		//合同模块接口
		auth.POST("contract", v1.EntryContract)
		auth.DELETE("contract/:uid", v1.DelContract)
		auth.PUT("contract", v1.EditContract)
		auth.GET("contract/:uid", v1.QueryContract)
		auth.POST("contracts", v1.QueryContracts)
		//任务模块接口
		auth.DELETE("task/:uid", v1.DelTask)
		auth.POST("tasks", v1.QueryTasks)
		//客户模块接口
		auth.POST("customer", v1.EntryCustomer)
		auth.DELETE("customer/:uid", v1.DelCustomer)
		auth.PUT("customer", v1.EditCustomer)
		auth.GET("customer/:uid", v1.QueryCustomer)
		auth.POST("customers", v1.QueryCustomers)
		//客户模块接口PLUS
		auth.POST("companys", v1.QueryCompanys)
		//产品模块接口
		auth.POST("product", v1.EntryProduct)
		auth.DELETE("product/:uid", v1.DelProduct)
		auth.PUT("product", v1.EditProduct)
		auth.GET("product/:uid", v1.QueryProduct)
		auth.POST("products", v1.QueryProducts)
		//供应商模块接口
		auth.POST("supplier", v1.EntrySupplier)
		auth.DELETE("supplier/:uid", v1.DelSupplier)
		auth.PUT("supplier", v1.EditSupplier)
		auth.GET("supplier/:uid", v1.QuerySupplier)
		auth.POST("suppliers", v1.QuerySuppliers)
	}

	_ = r.Run(config.SystemConfig.Server.Port)
}
