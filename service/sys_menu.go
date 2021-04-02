package service

import (
	"errors"
	"strconv"

	"github.com/gongxianjin/xcent-common/gorm"
	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/model"
	"github.com/gongxianjin/xcent_scaffold/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: authorityId string
//@return: err error, treeMap map[string][]model.SysMenu

func getMenuTreeMap(authorityId string) (err error, treeMap map[string][]model.SysMenu) {
	var allMenus []model.SysMenu
	treeMap = make(map[string][]model.SysMenu)
	err = lib.GORMDefaultPool.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	//for _, v := range allMenus {
	//	treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	//}
	return err, treeMap
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuTree
//@description: 获取动态菜单树
//@param: authorityId string
//@return: err error, menus []model.SysMenu

func GetMenuTree(authorityId string) (err error, menus []model.SysMenu) {
	err, menuTree := getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getChildrenList
//@description: 获取子菜单
//@param: menu *model.SysMenu, treeMap map[string][]model.SysMenu
//@return: err error

func getChildrenList(menu *model.SysMenu, treeMap map[string][]model.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetInfoList
//@description: 获取路由分页
//@return: err error, list interface{}, total int64

func GetInfoList(info request.PageInfo, id int,parentId string, name string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNo - 1)
	db :=  lib.GORMDefaultPool.Model(&model.SysBaseMenu{})
	var menuList []model.SysBaseMenu

	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if id != 0 {
		db = db.Where("id = ?", id)
	}
	if parentId != "" {
		db = db.Where("parent_id = ?", parentId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, menuList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Order("sort").Preload("Parameters").Find(&menuList).Error
	}

	return err, menuList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu
//@return: err error

func getBaseChildrenList(menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddBaseMenu
//@description: 添加基础路由
//@param: menu model.SysBaseMenu
//@return: err error

func AddBaseMenu(menu model.SysBaseMenu) (err error) {
	if !errors.Is(lib.GORMDefaultPool.Where("name = ?", menu.Name).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		err = errors.New("存在重复name，请修改name")
	}
	err = lib.GORMDefaultPool.Create(&menu).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: err error, treeMap map[string][]model.SysBaseMenu

func getBaseMenuTreeMap() (err error, treeMap map[string][]model.SysBaseMenu) {
	var allMenus []model.SysBaseMenu
	treeMap = make(map[string][]model.SysBaseMenu)
	err = lib.GORMDefaultPool.Order("sort").Preload("Parameters").Find(&allMenus).Error
	//for _, v := range allMenus {
	//	treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	//}
	return err, treeMap
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: err error, menus []model.SysBaseMenu

func GetBaseMenuTree() (err error, menus []model.SysBaseMenu) {
	err, treeMap := getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: err error, menus []model.SysBaseMenu

func GetBaseMenu() (err error, menus []model.SysBaseMenu) {
	var allMenus []model.SysBaseMenu
	err = lib.GORMDefaultPool.Order("sort").Preload("Parameters").Find(&allMenus).Error
	return err, allMenus
}



//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []model.SysBaseMenu, authorityId string
//@return: err error

func AddMenuAuthority(menus []model.SysBaseMenu, authorityId string) (err error) {
	var auth model.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = SetMenuAuthority(&auth)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: err error, menus []model.SysMenu

func GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []model.SysMenu) {
	err = lib.GORMDefaultPool.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	//err = lib.GORMDefaultPool.Raw(sql, authorityId).Scan(&menus).Error
	return err, menus
}
 

//@author: [granty1](https://github.com/granty1)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: BatchSetMenuStatusByIds
//@description: 批量设置菜单状态
//@param: ids request.IdsReq
//@return: err error

func BatchSetMenuStatusByIds(ids request.IdsReq,status uint) (err error) { 
	err = lib.GORMDefaultPool.Where("id in (?)", ids.Ids).Updates(map[string]interface{}{ 
		"show":status,
	}).Error 
	return err
}
