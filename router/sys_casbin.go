package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/controller/v1"
	"github.com/gongxianjin/xcent_scaffold/middleware"
)

func InitCasbinRouter(router *gin.RouterGroup) {
	curd := v1.SysCasbinController{} 
	CasbinRouter := router.Group("casbin").Use(
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{ 
		CasbinRouter.POST("/updateCasbin", curd.UpdateCasbin)
		CasbinRouter.POST("/getPolicyPathByAuthorityId", curd.GetPolicyPathByAuthorityId)
	}

}
