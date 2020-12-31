package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/controller"
)

func InitBaseRouter(router *gin.RouterGroup) {
	curd := controller.ApiController{}
	router.POST("/login", curd.Login)
	router.GET("/loginout", curd.LoginOut)
}