package response

import "github.com/gongxianjin/xcent_scaffold/model"

type SysUserResponse struct {
	User model.SysUser `json:"user"`
}

type LoginResponse struct {
	User      model.SysUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}

type SysUserInfoResponse struct {
	Id         uint          `json:"id"`
	Name       string        `json:"name"`
	Username   string        `json:"username"`
	Password   string        `json:"password"`
	Status     string        `json:"status"`
	RoleId     string        `json:"roleId"`
	Role       model.SysAuthority `json:"role"`
	CreateTime string        `json:"createTime"`
}
