package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/middleware"
	"github.com/gongxianjin/xcent_scaffold/model/request"
	"github.com/gongxianjin/xcent_scaffold/service"
	"github.com/gongxianjin/xcent_scaffold/utils"
)

type SysCasbinController struct {
}

// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /casbin/UpdateCasbin [post]
func (sysCasbin *SysCasbinController)UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	if err := utils.Verify(cmr, utils.AuthorityIdVerify); err != nil {
		middleware.ResponseError(c,400,err)
		return
	}
	if err := service.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos); err != nil {
		middleware.ResponseError(c,400,err)
	} else {
		middleware.ResponseSuccess(c,"")
	}
}

// @Tags Casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func (sysCasbin *SysCasbinController)GetPolicyPathByAuthorityId(c *gin.Context) {
	var casbin request.CasbinInReceive
	_ = c.ShouldBindJSON(&casbin)
	if err := utils.Verify(casbin, utils.AuthorityIdVerify); err != nil {
		middleware.ResponseError(c,401,err)
		return
	}
	paths := service.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	//response.OkWithDetailed(response.PolicyPathResponse{Paths: paths}, "获取成功", c)
	middleware.ResponseSuccess(c, paths)
}

