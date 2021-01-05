package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/controller"
)

func InitCasbinRouter(router *gin.RouterGroup) {
	curd := controller.SysCasbinController{}
	router.POST("/updateCasbin", curd.UpdateCasbin)
	router.POST("/getPolicyPathByAuthorityId", curd.GetPolicyPathByAuthorityId)
}
