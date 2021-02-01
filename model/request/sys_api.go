package request

import "github.com/gongxianjin/xcent_scaffold/model"

// api分页条件查询及排序结构体
type SearchApiParams struct {
	model.SysApi
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}