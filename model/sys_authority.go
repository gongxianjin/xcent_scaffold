package model

import (
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID"`
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
}

//type Sys_Data_Authority_Id struct {
//	Data_Authority_Id_Authority_Id string
//}


type SysDataAuthorityId struct {
	SysAuthorityAuthorityId string    `json:"sys_authority_authority_id" gorm:"primary_key"`
	DataAuthorityIdAuthorityId string `json:"data_authority_id_authority_id" gorm:"primary_key"`
}
