package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/controller"
	"github.com/gongxianjin/xcent_scaffold/middleware"
)

func InitUserRouter(router *gin.RouterGroup) {
	curd := controller.ApiController{}
	user := controller.UserController{}
	BaseRouter := router.Group("").Use(
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		BaseRouter.POST("/user/add", curd.AddUser)
		BaseRouter.POST("/user/edit", curd.EditUser)
		BaseRouter.POST("/user/remove", curd.RemoveUser)
		BaseRouter.POST("/user/batchremove", curd.RemoveUser)

		BaseRouter.PUT("/user/changePassword", user.ChangePassword)      // 修改密码
		BaseRouter.GET("/user/ListPage", user.ListPage)                  // 分页获取用户列表
		BaseRouter.POST("/user/setUserAuthority", user.SetUserAuthority) // 设置用户权限
		BaseRouter.DELETE("/user/deleteUser", user.DeleteUser)           // 删除用户
		BaseRouter.PUT("/user/setUserInfo", user.SetUserInfo)            // 设置用户信息
		BaseRouter.GET("/user/info", user.GetUserInfo)            //获取用户信息
		BaseRouter.POST("/user/logout", user.LoginOut)            //退出登录

	}
}
