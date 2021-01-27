package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/public"
)

type SysAuthorityInput struct {
	AuthorityId   string `json:"authorityId" form:"authorityId" comment:"角色ID" example:"888" validate:"required"`
	AuthorityName string  `json:"authorityName" form:"authorityName" comment:"角色ID" example:"test" validate:"required"`
	ParentId      string `json:"parentId" form:"parentId" comment:"父角色ID" example:"1" validate:"required"`
}

func (params *SysAuthorityInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
