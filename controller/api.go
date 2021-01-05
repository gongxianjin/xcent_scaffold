package controller

import (
	"errors"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/dao"
	"github.com/gongxianjin/xcent_scaffold/dto"
	"github.com/gongxianjin/xcent_scaffold/middleware"
)

type ApiController struct {
}

func (demo *ApiController) Login(c *gin.Context) {
	api := &dto.LoginInput{}
	if err := api.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	if api.Username == "admin" && api.Password == "123456" {
		session := sessions.Default(c)
		session.Set("user", api.Username)
		session.Set("user_id", "888")
		session.Save()
		middleware.ResponseSuccess(c, "")
		return
	}
	middleware.ResponseError(c, 2002, errors.New("账号或密码错误"))
	return
}

func (demo *ApiController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Delete("user_id")
	session.Save()
	return
}

func (demo *ApiController) ListPage(c *gin.Context) {
	listInput := &dto.ListPageInput{}
	if err := listInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	if listInput.PageSize == 0 {
		listInput.PageSize = 10
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	userList, total, err := (&dao.User{}).PageList(c, tx, listInput)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	m := &dao.ListPageOutput{
		List:  userList,
		Total: total,
	}
	middleware.ResponseSuccess(c, m)
	return
}

func (demo *ApiController) AddUser(c *gin.Context) {
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

func (demo *ApiController) EditUser(c *gin.Context) {
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

func (demo *ApiController) RemoveUser(c *gin.Context) {
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
