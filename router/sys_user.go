package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/controller"
	"github.com/gongxianjin/xcent_scaffold/middleware"
)

func InitUserRouter(router *gin.RouterGroup) {
	curd := controller.ApiController{}
	BaseRouter := router.Group("/api").Use(
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		BaseRouter.POST("/login", curd.Login)
		BaseRouter.GET("/loginout", curd.LoginOut)
	}
}
