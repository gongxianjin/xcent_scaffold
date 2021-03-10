package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/dao"
	"github.com/gongxianjin/xcent_scaffold/dto"
	"github.com/gongxianjin/xcent_scaffold/middleware"
	"github.com/gongxianjin/xcent_scaffold/model"
	"github.com/gongxianjin/xcent_scaffold/model/request"
	"github.com/gongxianjin/xcent_scaffold/model/response"
	"github.com/gongxianjin/xcent_scaffold/service"
	"github.com/gongxianjin/xcent_scaffold/utils"
)

type UserController struct {
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query int false "页码"
// @Param pageSize query int false "页条数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/ListPage [GET]
func (demo *UserController) ListPage(c *gin.Context) {
	// listInput :=  &dto.ListPageInput{}
	// if err := listInput.BindingValidParams(c); err != nil {
	// 	middleware.ResponseError(c, 2001, err)
	// 	return
	// }
	// if listInput.pageSize == 0 {
	// 	listInput.pageSize = 10
	// }
	// tx, err := lib.GetGormPool("default")
	// if err != nil {
	// 	middleware.ResponseError(c, 2002, err)
	// 	return
	// }
	// userList, total, err := (&dao.User{}).PageList(c, tx, listInput)
	// if err != nil {
	// 	middleware.ResponseError(c, 2003, err)
	// 	return
	// }
	// m := &dao.ListPageOutput{
	// 	List:  userList,
	// 	Total: total,
	// }
	//middleware.ResponseSuccess(c, m)
	log.Println("begin List")
	var pageSize = 10
	var pageIndex = 1
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, _ = strconv.Atoi(size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, _ = strconv.Atoi(index)
	}

	log.Printf("pageSize:%v,pageIndex:%v", pageSize, pageIndex)
	var pageInfo request.PageInfo
	pageInfo.PageNo = pageIndex
	pageInfo.PageSize = pageSize
	if err, list, total := service.GetUserInfoList(pageInfo); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		// response.OkWithDetailed(response.PageResult{
		// 	List:     list,
		// 	Total:    total,
		// 	Page:     pageInfo.Page,
		// 	PageSize: pageInfo.PageSize,
		// }, "获取成功", c)
		m := response.PageResult{
			List:     list,
			TotalCount:    total,
			PageNo:     pageInfo.PageNo,
			PageSize: pageInfo.PageSize,
		}
		middleware.ResponseSuccess(c, m)
	}
	return
}

// @Tags SysUser
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/info [GET]
func (demo *UserController) GetUserInfo(c *gin.Context) {
	if err, ReqUser := service.GetUserInfo(getUserUuid(c)); err != nil {
		log.Printf("获取失败!:%v", err)
		response.FailWithMessage("获取失败", c)
	} else {
		fmt.Println(ReqUser)
		//todo 格式化输出字段   
		// response.OkWithDetailed(response.SysUserInfoResponse{
		// 	Id: ReqUser.ID,
		// 	Role: ReqUser.Authority,
		// 	RoleId: ReqUser.AuthorityId,
		// }, "设置成功", c)

		// roleObj := `{"id": "admin","name": "管理员","describe": "拥有所有权限","status": 1, "creatorId": "system","createTime": 1497160610259,"deleted": 0,"permissions": [{"roleId": "admin","permissionId": "dashboard","permissionName": "仪表盘","actions": "[{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}]","actionEntitySet": [{"action": "add","describe": "新增","defaultCheck": false}, {"action": "query","describe": "查询","defaultCheck": false}, {"action": "get","describe": "详情","defaultCheck": false}, {"action": "update","describe": "修改","defaultCheck": false}, {"action": "delete","describe": "删除","defaultCheck": false}],"actionList": null,"dataAccess": null}}`
		roleObj := `{"id": "admin","name": "管理员","describe": "拥有所有权限","status": 1, "creatorId": "system","createTime": 1497160610259,"deleted": 0,
		"permissions": [
		{"roleId": "admin","permissionId": "dashboard","permissionName": "仪表盘",
		"actions": [{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}],
		"actionEntitySet":[{"action": "add","describe": "新增","defaultCheck": false}, {"action": "query","describe": "查询","defaultCheck": false}, {"action": "get","describe": "详情","defaultCheck": false}, {"action": "update","describe": "修改","defaultCheck": false}, {"action": "delete","describe": "删除","defaultCheck": false}],"actionList": null,"dataAccess": null}
		,{"roleId":"admin","permissionId":"exception","permissionName": "异常页面权限",
		"actions": [{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}],
		"actionEntitySet":[{"action": "add", "describe": "新增", "defaultCheck": false}, {"action": "query","describe": "查询","defaultCheck": false}, {"action": "get","describe": "详情","defaultCheck": false}, {"action": "update","describe": "修改","defaultCheck": false}, {"action": "delete","describe": "删除","defaultCheck": false}],"actionList": null,"dataAccess": null}
		]}`
		var jsonData map[string]interface{} 
		if e := json.Unmarshal([]byte(roleObj), &jsonData); e != nil { 
				log.Fatalf("%s",e.Error());
		}  
		fmt.Println(jsonData)
		response.OkWithDetailed(gin.H{ 
		"id": ReqUser.ID,
    "name": ReqUser.NickName,
    "username": ReqUser.Username,
    "password": ReqUser.Password,
    "avatar": ReqUser.HeaderImg, 
    "telephone": ReqUser.Phone,
    "lastLoginIp": "",
    "lastLoginTime": "",
    "creatorId": ReqUser.AuthorityId,
    "createTime": ReqUser.CreatedAt,
    "merchantCode": "",
    "deleted": ReqUser.DeletedAt,
    "roleId": ReqUser.AuthorityId,
    "role": jsonData,
		}, "设置成功", c)
	}
}


// @Tags SysUser
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登出成功"}"
// @Router /user/logout [POST]
func (demo *UserController) LoginOut(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := model.JwtBlacklist{Jwt: token}
	session := sessions.Default(c)
	session.Set("user", "")
	session.Set("user_id", "")
	session.Save()
	if err := service.JsonInBlacklist(jwt); err != nil{
		log.Printf("获取失败!:%v", err)
		response.FailWithMessage("登出失败", c)
	} else { 
		response.OkWithMessage("登出成功", c)
	}
}



func (demo *UserController) AddUser(c *gin.Context) {
	addInput := &dto.AddUserInput{}
	if err := addInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	user := &dao.User{
		Name:  addInput.Name,
		Sex:   addInput.Sex,
		Age:   addInput.Age,
		Birth: addInput.Birth,
		Addr:  addInput.Addr,
	}
	if err := user.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

func (demo *UserController) EditUser(c *gin.Context) {
	editInput := &dto.EditUserInput{}
	if err := editInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	user, err := (&dao.User{}).Find(c, tx, int64(editInput.Id))
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	user.Name = editInput.Name
	user.Sex = editInput.Sex
	user.Age = editInput.Age
	user.Birth = editInput.Birth
	user.Addr = editInput.Addr
	if err := user.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

func (demo *UserController) RemoveUser(c *gin.Context) {
	removeInput := &dto.RemoveUserInput{}
	if err := removeInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	if err := (&dao.User{}).Del(c, tx, strings.Split(removeInput.IDS, ",")); err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func (demo *UserController) ChangePassword(c *gin.Context) {
	var user request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	U := &model.SysUser{Username: user.Username, Password: user.Password}
	if err, _ := service.ChangePassword(U, user.NewPassword); err != nil {
		log.Printf("修改失败:%v", err)
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func (demo *UserController) SetUserAuthority(c *gin.Context) {
	var sua request.SetUserAuth
	_ = c.ShouldBindJSON(&sua)
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	if err := service.SetUserAuthority(sua.UUID, sua.AuthorityId); err != nil {
		log.Printf("修改失败:%v", err)
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
func (demo *UserController) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := getUserID(c)
	if jwtId == uint(reqId.Id) {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	if err := service.DeleteUser(reqId.Id); err != nil {
		log.Printf("修改失败:%v", err)
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func getUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		log.Print("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件")
		return 0
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.ID
	}
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /user/setUserInfo [put]
func (demo *UserController) SetUserInfo(c *gin.Context) {
	var user model.SysUser
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.SetUserVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, ReqUser := service.SetUserInfo(user); err != nil {
		log.Printf("设置失败:%v", err)
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
	}
}

// 从Gin的Context中获取从jwt解析出来的用户UUID
func getUserUuid(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		log.Printf("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UUID.String()
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func getUserAuthorityId(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		log.Printf("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.AuthorityId
	}
}
