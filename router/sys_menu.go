package router

import ( 
	"github.com/gongxianjin/xcent_scaffold/controller/v1"
	"github.com/gongxianjin/xcent_scaffold/middleware"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	{
		curd := v1.SysMenuController{} 
		MenuRouter.POST("getMenu", curd.GetMenu)                   // 获取菜单树
		MenuRouter.POST("getMenuList", curd.GetMenuList)           // 分页获取基础menu列表
		MenuRouter.POST("addBaseMenu", curd.AddBaseMenu)           // 新增菜单
		MenuRouter.POST("getBaseMenuTree", curd.GetBaseMenuTree)   // 获取基础路由树
		MenuRouter.POST("addMenuAuthority", curd.AddMenuAuthority) //	增加menu和角色关联关系
		MenuRouter.POST("getMenuAuthority", curd.GetMenuAuthority) // 获取指定角色menu
		MenuRouter.POST("deleteBaseMenu", curd.DeleteBaseMenu)     // 删除菜单
		MenuRouter.POST("updateBaseMenu", curd.UpdateBaseMenu)     // 更新菜单
		MenuRouter.POST("getBaseMenuById", curd.GetBaseMenuById)   // 根据id获取菜单
		MenuRouter.POST("batchSetMenuStatus", curd.BatchSetMenuStatus)     // 批量设置菜单
	}
	return MenuRouter
}
