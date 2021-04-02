package request

import "github.com/gongxianjin/xcent_scaffold/model"

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu
	AuthorityId string
}


// menu分页条件查询及排序结构体
type SearchMenuParams struct {
	PageInfo
	Id  int `json:"id"`
	ParentId string `json:"parentId"`
	Name     string   `json:"name"`
}

type BatchSetMenuParams struct { 
	Ids []int `json:"ids" form:"ids"` 
	Show     int   `json:"show"`
}
