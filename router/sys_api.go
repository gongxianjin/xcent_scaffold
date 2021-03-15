package router

import ( 
	"github.com/gongxianjin/xcent_scaffold/controller/v1"
	"github.com/gongxianjin/xcent_scaffold/middleware"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api").Use(middleware.OperationRecord())
	{
		curd := v1.SysApiController{} 
		ApiRouter.POST("createApi", curd.CreateApi)   // 创建Api
		ApiRouter.POST("deleteApi", curd.DeleteApi)   // 删除Api
		ApiRouter.POST("getApiList", curd.GetApiList) // 获取Api列表
		ApiRouter.POST("getApiById", curd.GetApiById) // 获取单条Api消息
		ApiRouter.POST("updateApi", curd.UpdateApi)   // 更新api
		ApiRouter.POST("getAllApis", curd.GetAllApis) // 获取所有api
	}
}
