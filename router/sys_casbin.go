package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/controller"
	"github.com/gongxianjin/xcent_scaffold/middleware"
)

func InitCasbinRouter(router *gin.RouterGroup) {
	curd := controller.SysCasbinController{} 
	CasbinRouter := router.Group("casbin").Use(
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{ 
		CasbinRouter.POST("/updateCasbin", curd.UpdateCasbin)
		CasbinRouter.POST("/getPolicyPathByAuthorityId", curd.GetPolicyPathByAuthorityId)
	}

}
