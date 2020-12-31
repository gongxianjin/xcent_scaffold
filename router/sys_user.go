package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/controller"
)

func InitUserRouter(router *gin.RouterGroup) {
	curd := controller.ApiController{}
	router.GET("/user/listpage", curd.ListPage)
	router.POST("/user/add", curd.AddUser)
	router.POST("/user/edit", curd.EditUser)
	router.POST("/user/remove", curd.RemoveUser)
	router.POST("/user/batchremove", curd.RemoveUser)
}
