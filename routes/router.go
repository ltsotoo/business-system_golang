package routes

import (
	"business-system_golang/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.SystemConfig.Server.Mode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(config.SystemConfig.Server.Port)
}
