package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/controller"
	"github.com/gongxianjin/xcent_scaffold/middleware"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("casbin").Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog())
	{
		curd := controller.SysCasbinController{}
		CasbinRouter.POST("updateCasbin", curd.UpdateCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", curd.GetPolicyPathByAuthorityId)
	}
}
