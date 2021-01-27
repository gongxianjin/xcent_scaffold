package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/public"
)

type CreateAuthorityInput struct {
	AuthorityId   string `json:"authorityId" form:"authorityId" comment:"角色ID" example:"888" validate:"required"`
	AuthorityName string `json:"authorityName" form:"authorityName" comment:"角色ID" example:"test" validate:"required"`
	ParentId      string `json:"parentId" form:"parentId" comment:"父角色ID" example:"1" validate:"required"`
}


func (params *CreateAuthorityInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type CopyAuthorityInput struct {
	AuthorityId   string `json:"authorityId" form:"authorityId" comment:"新权限id" example:"888" validate:"required"`
	AuthorityName string `json:"authorityName" form:"authorityName" comment:"新权限名" example:"test" validate:"required"`
	ParentId      string `json:"parentId" form:"parentId" comment:"新父角色id" example:"1" validate:"required"`
	OldAuthorityId   string `json:"oldAuthorityId" form:"oldAuthorityId" comment:"旧角色id" example:"2" validate:"required"`
}

func (params *CopyAuthorityInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}