package routes

import (
	v1 "business-system_golang/api/v1"
	"business-system_golang/config"
	"business-system_golang/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.SystemConfig.Server.Mode)
	r := gin.Default()

	r.Use(middleware.Cors())
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"*"}
	// r.Use(cors.New(corsConfig))

	routeV1 := r.Group("api/v1")
	{
		//员工模块接口
		routeV1.POST("employee", v1.EntryEmployee)
		routeV1.DELETE("employee/:id", v1.DelEmployee)
		routeV1.PUT("employee", v1.EditEmployee)
		routeV1.GET("employee/:id", v1.QueryEmployee)
		routeV1.GET("employees", v1.QueryEmployees)
		//合同模块接口
		routeV1.POST("contract", v1.EntryContract)
		routeV1.DELETE("contract/:id", v1.DelContract)
		routeV1.PUT("contract", v1.EditContract)
		routeV1.GET("contract/:id", v1.QueryContract)
		routeV1.GET("contracts", v1.QueryContracts)
		//任务模块接口
		//客户模块接口
		routeV1.POST("customer", v1.EntryCustomer)
		routeV1.DELETE("customer/:id", v1.DelCustomer)
		routeV1.PUT("customer", v1.EditCustomer)
		routeV1.GET("customer/:id", v1.QueryCustomer)
		routeV1.GET("customers", v1.QueryCustomers)
		//产品模块接口
		routeV1.POST("product", v1.EntryProduct)
		routeV1.DELETE("product/:id", v1.DelProduct)
		routeV1.PUT("product", v1.EditProduct)
		routeV1.GET("product/:id", v1.QueryProduct)
		routeV1.POST("products", v1.QueryProducts)
		//供应商模块接口
		routeV1.POST("supplier", v1.EntrySupplier)
		routeV1.DELETE("supplier/:id", v1.DelSupplier)
		routeV1.PUT("supplier", v1.EditSupplier)
		routeV1.GET("supplier/:id", v1.QuerySupplier)
		routeV1.GET("suppliers", v1.QuerySuppliers)
		//字典表模块
		routeV1.GET("systemDictionaryValuesByKeyId", v1.QuerySystemDictionaryValuesByKeyId)
		routeV1.GET("systemDictionaryValuesByParentId", v1.QuerySystemDictionaryValuesByParentId)
	}

	_ = r.Run(config.SystemConfig.Server.Port)
}
