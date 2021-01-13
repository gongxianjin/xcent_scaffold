package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/controller"
	"github.com/gongxianjin/xcent_scaffold/middleware"
)

func InitUserRouter(router *gin.RouterGroup) {
	curd := controller.ApiController{}
	user := controller.UserController{}
	BaseRouter := router.Group("/api").Use(
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		BaseRouter.POST("/user/add", curd.AddUser)
		BaseRouter.POST("/user/edit", curd.EditUser)
		BaseRouter.POST("/user/remove", curd.RemoveUser)
		BaseRouter.POST("/user/batchremove", curd.RemoveUser)
		
		BaseRouter.GET("/user/listpage", user.ListPage)
	}
}
