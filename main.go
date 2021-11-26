package main

import (
	"business-system_golang/model"
	"business-system_golang/routes"
)

func main() {

	//数据库连接初始化
	model.InitDb()

	//定时任务启动
	// middleware.UpdateAllEmployeesMoney()
	model.InitCronTabs()

	//路由初始化
	routes.InitRouter()
}
