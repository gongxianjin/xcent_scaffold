package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/controller"
	"github.com/gongxianjin/xcent_scaffold/middleware"
)

func InitBaseRouter(router *gin.RouterGroup) {
	curd := controller.ApiController{}
	BaseRouter := router.Group("base").Use(
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		BaseRouter.POST("login", curd.Login)
		BaseRouter.POST("captcha", curd.WechatCaptcha)
		BaseRouter.POST("register", curd.Register)
		BaseRouter.GET("loginOut", curd.LoginOut)
	}
}