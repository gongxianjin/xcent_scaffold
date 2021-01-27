package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/dto"
	"github.com/gongxianjin/xcent_scaffold/middleware"
	"github.com/gongxianjin/xcent_scaffold/model"
	"github.com/gongxianjin/xcent_scaffold/model/request"
	"github.com/gongxianjin/xcent_scaffold/model/response"
	"github.com/gongxianjin/xcent_scaffold/service"
	"github.com/gongxianjin/xcent_scaffold/utils"
)

type SysAuthorityController struct {
}

// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.SysAuthorityInput true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]
func (SysAuthority *SysAuthorityController) CreateAuthority(c *gin.Context) {
	// if err := utils.Verify(authority, utils.AuthorityVerify); err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	params := &dto.SysAuthorityInput{}
	if err := params.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	var authority model.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	authority.AuthorityId = params.AuthorityId
	authority.AuthorityName = params.AuthorityName
	authority.ParentId = params.ParentId
	if err, authBack := service.CreateAuthority(authority); err != nil {
		log.Printf("创建失败!:%v", err)
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}

// @Tags Authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.SysAuthorityCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router /authority/copyAuthority [post]
func (SysAuthority *SysAuthorityController) CopyAuthority(c *gin.Context) {
	var copyInfo response.SysAuthorityCopyResponse
	_ = c.ShouldBindJSON(&copyInfo)
	if err := utils.Verify(copyInfo, utils.OldAuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(copyInfo.Authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := service.CopyAuthority(copyInfo); err != nil {
		log.Printf("拷贝失败!:%v", err)
		response.FailWithMessage("拷贝失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "拷贝成功", c)
	}
}

// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authority/deleteAuthority [post]
func (SysAuthority *SysAuthorityController) DeleteAuthority(c *gin.Context) {
	var authority model.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		log.Printf("删除失败!:%v", err)
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authority/updateAuthority [post]
func (SysAuthority *SysAuthorityController) UpdateAuthority(c *gin.Context) {
	var auth model.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authority := service.UpdateAuthority(auth); err != nil {
		log.Printf("更新失败!:%v", err)
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authority}, "更新成功", c)
	}
}

// @Tags Authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]
func (SysAuthority *SysAuthorityController) GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetAuthorityInfoList(pageInfo); err != nil {
		log.Printf("获取失败!:%v", err)
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]
func (SysAuthority *SysAuthorityController) SetDataAuthority(c *gin.Context) {
	var auth model.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.SetDataAuthority(auth); err != nil {
		log.Printf("设置失败!:%v", err)
		response.FailWithMessage("设置失败"+err.Error(), c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}
