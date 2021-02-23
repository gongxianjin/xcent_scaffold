package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/gin-gonic/gin" 
	"github.com/gongxianjin/xcent_scaffold/model"
	"github.com/gongxianjin/xcent_scaffold/model/request"
	"github.com/gongxianjin/xcent_scaffold/model/response"
	"github.com/gongxianjin/xcent_scaffold/service"
	"github.com/gongxianjin/xcent_scaffold/utils"
)


type SysMenuController struct {
}


// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenu [post]
func (SysMenu *SysMenuController)GetMenu(c *gin.Context) {
	if err, menus := service.GetMenuTree(getUserAuthorityId(c)); err != nil { 
		log.Printf("获取失败!:%v", err)
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuTree [post]
func (SysMenu *SysMenuController)GetBaseMenuTree(c *gin.Context) {
	if err, menus := service.GetBaseMenuTree(); err != nil { 
		log.Printf("获取失败!:%v", err)
		response.FailWithMessage("获取失败", c)
	} else {
		//response.OkWithDetailed(response.SysBaseMenusResponse{Menus: menus}, "获取成功", c)
		roleObj := `{"menu": [
						{"name":"dashboard","parentId": 0,"id": 1,"meta": { "keepAlive": false, "defaultMenu": false, "icon": "dashboard", "title": "测试" }, "component": "RouteView", "redirect": "/dashboard/workplace"},
						{"name":"workplace","parentId": 1,"id": 2,"meta": { "keepAlive": false, "defaultMenu": false, "icon": "dashboard", "title": "二级页面" }, "component": "Workplace", "redirect": ""}
					]}`
		var jsonData map[string]interface{}
		if e := json.Unmarshal([]byte(roleObj), &jsonData); e != nil {
			log.Fatalf("%s",e.Error())
		}
		fmt.Println(menus)
		fmt.Println(jsonData)
		response.OkWithDetailed(gin.H{
			"menu" : jsonData,
		}, "设置成功", c)
	}
}

// @Tags AuthorityMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMenuAuthorityInfo true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addMenuAuthority [post]
func (SysMenu *SysMenuController)AddMenuAuthority(c *gin.Context) {
	var authorityMenu request.AddMenuAuthorityInfo
	_ = c.ShouldBindJSON(&authorityMenu)
	if err := utils.Verify(authorityMenu, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil { 
		log.Printf("获取失败!:%v", err)
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// @Tags AuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuAuthority [post]
func (SysMenu *SysMenuController)GetMenuAuthority(c *gin.Context) {
	var param request.GetAuthorityId
	_ = c.ShouldBindJSON(&param)
	if err := utils.Verify(param, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menus := service.GetMenuAuthority(&param); err != nil { 
		log.Printf("获取失败!:%v", err)
		response.FailWithDetailed(response.SysMenusResponse{Menus: menus}, "获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addBaseMenu [post]
func (SysMenu *SysMenuController)AddBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.AddBaseMenu(menu); err != nil { 
		log.Printf("获取失败!:%v", err)

		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menu/deleteBaseMenu [post]
func (SysMenu *SysMenuController)DeleteBaseMenu(c *gin.Context) {
	var menu request.GetById
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteBaseMenu(menu.Id); err != nil { 
		log.Printf("获取失败!:%v", err)
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /menu/updateBaseMenu [post]
func (SysMenu *SysMenuController)UpdateBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateBaseMenu(menu); err != nil { 
		log.Printf("更新失败!:%v", err)
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func (SysMenu *SysMenuController)GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menu := service.GetBaseMenuById(idInfo.Id); err != nil { 
		log.Printf("获取失败!:%v", err)
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func (SysMenu *SysMenuController)GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menuList, total := service.GetInfoList(); err != nil {
		log.Printf("获取失败!:%v", err)
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		},"获取成功", c)
	}
}