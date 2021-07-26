package routes

import (
	v1 "business-system_golang/api/v1"
	"business-system_golang/config"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.SystemConfig.Server.Mode)
	r := gin.Default()

	routeV1 := r.Group("api/v1")
	{
		//员工模块接口
		routeV1.POST("employee", v1.EntryEmployee)
		routeV1.DELETE("employee/:id", v1.DelEmployee)
		routeV1.PUT("employee", v1.EditEmployee)
		routeV1.GET("employee/:id", v1.QueryEmployee)
		routeV1.GET("employees", v1.QueryEmployees)
		//合同模块接口
		//任务模块接口
		//客户模块接口
		routeV1.POST("customer", v1.EntryCustomer)
		routeV1.DELETE("customer/:id", v1.DelCustomer)
		routeV1.PUT("customer", v1.EditCustomer)
		routeV1.GET("customer/:id", v1.QueryCustomer)
		routeV1.GET("customer", v1.QueryCustomers)
		//产品模块接口
		//供应商模块接口
	}

	r.Run(config.SystemConfig.Server.Port)
}
