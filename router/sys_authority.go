package router

import ( 
	"github.com/gin-gonic/gin" 
	"github.com/gongxianjin/xcent_scaffold/controller/v1"
	"github.com/gongxianjin/xcent_scaffold/middleware"
)

func InitAuthorityRouter(Router *gin.RouterGroup) {
	curd := v1.SysAuthorityController{} 
	AuthorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	{
		AuthorityRouter.POST("createAuthority", curd.CreateAuthority)   // 创建角色
		AuthorityRouter.POST("deleteAuthority", curd.DeleteAuthority)   // 删除角色
		AuthorityRouter.PUT("updateAuthority", curd.UpdateAuthority)    // 更新角色
		AuthorityRouter.POST("copyAuthority", curd.CopyAuthority)       // 更新角色
		AuthorityRouter.POST("getAuthorityList", curd.GetAuthorityList) // 获取角色列表
		AuthorityRouter.POST("setDataAuthority", curd.SetDataAuthority) // 设置角色资源权限
	}
}
