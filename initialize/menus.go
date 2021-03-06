package initialize

import (
	"github.com/gongxianjin/xcent-common/gorm" 
	"github.com/gongxianjin/xcent_scaffold/model"
	"log"
	"time"
)

var BaseMenus = []model.SysBaseMenu{
	{BaseModel: model.BaseModel{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId:"0", Redirect: "/dashboard/workplace", Name: "dashboard", Hidden: false, Component: "RouteView", Sort: 1, Meta: model.Meta{Title: "仪表盘", Icon: "setting",Show: true}},
	{BaseModel: model.BaseModel{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId:"1", Path: "", Name: "workplace", Component: "Workplace", Sort: 0, Meta: model.Meta{Title: "工作台", Icon: "info",Show: true}},
	{BaseModel: model.BaseModel{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Redirect: "/sys/menu", Name: "sys", Component: "RouteView", Sort: 3, Meta: model.Meta{Title: "超级管理员", Icon: "user-solid",Show: true}},
	{BaseModel: model.BaseModel{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Name: "RoleList", Component: "RoleList", Sort: 1, Meta: model.Meta{Title: "角色管理", Icon: "s-custom",Show: true}},
	{BaseModel: model.BaseModel{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Name: "sysMenu", Component: "sysMenu", Sort: 2, Meta: model.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true,Show: true}},
	{BaseModel: model.BaseModel{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Name: "PermissionList", Component: "PermissionList", Sort: 3, Meta: model.Meta{Title: "api管理", Icon: "s-platform",Show: true}},
	{BaseModel: model.BaseModel{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Name: "UserList", Component: "UserList", Sort: 4, Meta: model.Meta{Title: "用户管理", Icon: "coordinate",Show: true}},
	{BaseModel: model.BaseModel{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: model.Meta{Title: "个人信息", Icon: "message-solid",Show: true}},
	{BaseModel: model.BaseModel{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: model.Meta{Title: "示例文件", Icon: "s-management",Show: true}},
	{BaseModel: model.BaseModel{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "table", Name: "table", Component: "view/example/table/table.vue", Sort: 1, Meta: model.Meta{Title: "表格示例", Icon: "s-order",Show: true}},
	{BaseModel: model.BaseModel{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "form", Name: "form", Component: "view/example/form/form.vue", Sort: 2, Meta: model.Meta{Title: "表单示例", Icon: "document",Show: true}},
	{BaseModel: model.BaseModel{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "rte", Name: "rte", Component: "view/example/rte/rte.vue", Sort: 3, Meta: model.Meta{Title: "富文本编辑器", Icon: "reading",Show: true}},
	{BaseModel: model.BaseModel{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: model.Meta{Title: "excel导入导出", Icon: "s-marketing",Show: true}},
	{BaseModel: model.BaseModel{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: model.Meta{Title: "上传下载", Icon: "upload",Show: true}},
	{BaseModel: model.BaseModel{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: model.Meta{Title: "断点续传", Icon: "upload",Show: true}},
	{BaseModel: model.BaseModel{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: model.Meta{Title: "客户列表（资源示例）", Icon: "s-custom",Show: true}},
	{BaseModel: model.BaseModel{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: model.Meta{Title: "系统工具", Icon: "s-cooperation",Show: true}},
	{BaseModel: model.BaseModel{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: model.Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true,Show: true}},
	{BaseModel: model.BaseModel{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: model.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
	{BaseModel: model.BaseModel{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: model.Meta{Title: "系统配置", Icon: "s-operation",Show: true}},
	{BaseModel: model.BaseModel{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "iconList", Name: "iconList", Component: "view/iconList/index.vue", Sort: 2, Meta: model.Meta{Title: "图标集合", Icon: "star-on",Show: true}},
	{BaseModel: model.BaseModel{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: model.Meta{Title: "字典管理", Icon: "notebook-2"}},
	{BaseModel: model.BaseModel{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: model.Meta{Title: "字典详情", Icon: "s-order"}},
	{BaseModel: model.BaseModel{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: model.Meta{Title: "操作历史", Icon: "time"}},
	{BaseModel: model.BaseModel{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6, Meta: model.Meta{Title: "断点续传（插件版）", Icon: "upload"}},
	{BaseModel: model.BaseModel{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "1",  Path: "https://www.baidu.com", Name:"test", Hidden: false, Component: "/", Sort: 4, Meta: model.Meta{Title: "官方网站", Icon: "s-home",Show:true}},
	{BaseModel: model.BaseModel{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: model.Meta{Title: "服务器状态", Icon: "cloudy"}},
	{BaseModel: model.BaseModel{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "workflow", Name: "workflow", Hidden: false, Component: "view/workflow/index.vue", Sort: 5, Meta: model.Meta{Title: "工作流功能", Icon: "phone"}},
	{BaseModel: model.BaseModel{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "28", Path: "workflowCreate", Name: "workflowCreate", Hidden: false, Component: "view/workflow/workflowCreate/workflowCreate.vue", Sort: 0, Meta: model.Meta{Title: "工作流绘制", Icon: "circle-plus"}},
	{BaseModel: model.BaseModel{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "28", Path: "workflowProcess", Name: "workflowProcess", Hidden: false, Component: "view/workflow/workflowProcess/workflowProcess.vue", Sort: 0, Meta: model.Meta{Title: "工作流列表", Icon: "s-cooperation"}},
	{BaseModel: model.BaseModel{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "28", Path: "workflowUse", Name: "workflowUse", Hidden: true, Component: "view/workflow/workflowUse/workflowUse.vue", Sort: 0, Meta: model.Meta{Title: "使用工作流", Icon: "video-play"}},
	{BaseModel: model.BaseModel{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "28", Path: "started", Name: "started", Hidden: false, Component: "view/workflow/userList/started.vue", Sort: 0, Meta: model.Meta{Title: "我发起的", Icon: "s-order"}},
	{BaseModel: model.BaseModel{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "28", Path: "need", Name: "need", Hidden: false, Component: "view/workflow/userList/need.vue", Sort: 0, Meta: model.Meta{Title: "我的待办", Icon: "s-platform"}},
}

func InitSysBaseMenus(db *gorm.DB) {
	if db.Where("id IN (?)", []int{1, 27}).Find(&[]model.SysBaseMenu{}).RowsAffected == 2 {
		log.Println("sys_base_menus表的初始数据已存在!")
		return
	}
	db = db.Begin()
	for _,api := range BaseMenus {
		if err := db.Debug().Save(&api).Error; err != nil { // 遇到错误时回滚事务
			db.Rollback()
			log.Fatal(err)
		}
	}
	db.Commit()
	log.Println("init BaseMenu success")
}
