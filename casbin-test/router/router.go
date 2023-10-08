package router

import (
	"casbin-test/common"
	"casbin-test/router/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// CasBin权限认证
	router := gin.Default()
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/addPolicy", handler.AddPolicy)
		authGroup.GET("/testPolicy", common.CasBinMiddleware(), handler.TestListPolicics)
	}
	return router
}
