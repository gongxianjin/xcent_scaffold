package controller

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/dao"
	"github.com/gongxianjin/xcent_scaffold/dto"
	"github.com/gongxianjin/xcent_scaffold/middleware"
	"github.com/gongxianjin/xcent_scaffold/model"
	"github.com/gongxianjin/xcent_scaffold/model/request"
	"github.com/gongxianjin/xcent_scaffold/model/response"
	"github.com/gongxianjin/xcent_scaffold/service"
)

type ApiController struct {
}

// @Tags Base
// @Summary 用户登录
// @Description 用户登录
// @Tags 用户
// @ID  /base/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.LoginInput true "body"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (demo *ApiController) Login(c *gin.Context) {
	api := &dto.LoginInput{}
	if err := api.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	U := &model.SysUser{Username: api.Username, Password: api.Password}
	if err, user := service.Login(U); err != nil {
		log.Printf("登陆失败! 用户名不存在或者密码错误:%v", err)
		middleware.ResponseError(c, 2002, errors.New("用户名不存在或者密码错误"))
	} else {
		tokenNext(c, *user)
	}
	// if api.Username == "admin" && api.Password == "123456" {
	// 	session := sessions.Default(c)
	// 	session.Set("user", api.Username)
	// 	session.Set("user_id", "888")
	// 	session.Save()
	// 	middleware.ResponseSuccess(c, "")
	// 	return
	// }
	// middleware.ResponseError(c, 2002, errors.New("账号或密码错误"))
	return
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	log.Println("gen Token begin")
	j := &middleware.JWT{SigningKey: []byte(lib.GetStringConf("base.jwt.signing-key"))} // 唯一签名
	claims := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		BufferTime:  60 * 60 * 24, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 7天
			Issuer:    "XCENT",                        // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		log.Fatalf("获取token失败:%v", err)
		middleware.ResponseError(c, 2002, errors.New("获取token失败"))
		return
	}
	if !lib.GetBoolConf("base.system.use-multipoint") {
		log.Printf("gen Token to no-multipoint:%v", token)
		middleware.ResponseSuccess(c, response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		})
		return
	}
	if err, jwtStr := service.GetRedisJWT(user.Username); err == redis.Nil {
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			log.Fatalf("设置登录状态失败:%v", err)
			middleware.ResponseError(c, 2002, errors.New("设置登录状态失败"))
			return
		}
		log.Printf("gen Token to redis:%v", token)
		middleware.ResponseSuccess(c, "")
	} else if err != nil {
		log.Fatalf("设置登录状态失败:%v", err)
		middleware.ResponseError(c, 2002, errors.New("设置登录状态失败"))
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			middleware.ResponseError(c, 2002, errors.New("jwt作废失败"))
			return
		}
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			middleware.ResponseError(c, 2002, errors.New("设置登录状态失败"))
			return
		}
		log.Printf("gen Token to black:%v", token)
		middleware.ResponseSuccess(c, "")
	}
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
