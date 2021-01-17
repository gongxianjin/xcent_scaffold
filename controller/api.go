package controller

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gongxianjin/xcent_scaffold/utils"

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
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

type ApiController struct {
}

// @Tags Base
// @Summary 用户登录
// @Description 用户登录
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
	if !store.Verify(api.CaptchaId, api.Captcha, true) {
		middleware.ResponseError(c, 2002, errors.New("验证码错误"))
		return
	}
	U := &model.SysUser{Username: api.Username, Password: api.Password}
	if err, user := service.Login(U); err != nil {
		log.Printf("登陆失败! 用户名不存在或者密码错误:%v", err)
		middleware.ResponseError(c, 2002, errors.New("用户名不存在或者密码错误"))
		return
	} else {
		session := sessions.Default(c)
		session.Set("user", user.Username)
		session.Set("user_id", user.AuthorityId)
		session.Save()
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

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.Register true "用户名, 昵称, 手机号，密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]
func (demo *ApiController) Register(c *gin.Context) {
	var R request.Register
	_ = c.ShouldBindJSON(&R)
	if err := utils.Verify(R, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &model.SysUser{Username: R.Username, NickName: R.NickName,Phone: R.Phone,Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
	err, userReturn := service.Register(*user)
	if err != nil {
		log.Printf("注册失败：%v", err)
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

func (demo *ApiController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Delete("user_id")
	session.Save()
	return
}

func (demo *ApiController) ListPageBAK(c *gin.Context) {
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

// @Tags Base
// @Summary 生成图片验证码
// @accept application/json
// @Produce application/json 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/picCaptcha [post]
func (demo *ApiController) PicCaptcha(c *gin.Context) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	log.Printf("img-height:%v,width:%v,key-long:%v", lib.GetIntConf("base.Captcha.img-height"), lib.GetIntConf("base.Captcha.img-width"), lib.GetIntConf("base.Captcha.key-long"))
	driver := base64Captcha.NewDriverDigit(lib.GetIntConf("base.Captcha.img-height"), lib.GetIntConf("base.Captcha.img-width"), lib.GetIntConf("base.Captcha.key-long"), 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		log.Printf("验证码获取失败!:%v", err)
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}

// 识别手机号码
func isMobile(mobile string) bool{
	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	if result {
		log.Println(`正确的手机号`)
		return true
	} else {
		log.Println(`错误的手机号`)
		return false
	}
}

// @Tags Base
// @Summary 生成短信验证码
// @accept multipart/form-data
// @Produce  application/json
// @Param  phone formData  string true "手机号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func (demo *ApiController) MessageCaptcha(c *gin.Context) {
	//获取phoneNumb
	phone := c.Request.FormValue("phone")
	//验证手机号规则
	if !isMobile(phone) {
		c.JSON(http.StatusOK, response.Response{
			Code: 400,
			Data: "",
			Msg:  "手机号不符合规则",
		})
		return
	}
	//随机生成4位数
	randCode := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)
	log.Printf("phone:%v,randCode:%v",phone,randCode)
	//存入缓存中 60s 过期
	timer := 60
	redisObj, err := lib.RedisConnFactory("default")
	if err != nil {
		log.Fatalf("init redis:%v",err)
	}
	defer redisObj.Close()
	trace := lib.NewTrace()
	//code 存入 key 为手机号的主键
	lib.RedisLogDo(trace, redisObj, "SET", phone,randCode)
	lib.RedisLogDo(trace, redisObj, "expire", phone, timer)
	//todo 调用短信方接口发送短信
	middleware.ResponseSuccess(c,dto.SmsResponse{
		Code: randCode,
		Msg: fmt.Sprintf("短信验证码为:%d,请勿告知他人",randCode),
	})
}

// @Tags Base
// @Summary 生成微信验证码
// @accept multipart/form-data
// @Produce  application/json
// @Param  openId formData  string true "微信openID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/wechatCaptcha [post]
func (demo *ApiController) WechatCaptcha(c *gin.Context) {
	//获取openID
	openId := c.Request.FormValue("openId")
	//data := c.PostForm("openId")
	//buf := make([]byte, 1024)
	//n, _ := c.Request.Body.Read(buf)
	//log.Printf("data:%v,buf:%v",data,string(buf[0:n]))
	log.Printf("openId:%v",openId)
	//随机生成4位数
	randCode := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)
	log.Printf("phone:%v,randCode:%v",openId,randCode)
	//存入缓存中 60s 过期
	timer := 60
	redisObj, err := lib.RedisConnFactory("default")
	if err != nil {
		log.Fatalf("init redis:%v",err)
	}
	defer redisObj.Close()
	trace := lib.NewTrace()
	//code 存入 key 为手机号的主键
	lib.RedisLogDo(trace, redisObj, "SET", openId,randCode)
	lib.RedisLogDo(trace, redisObj, "expire", openId, timer)
	//todo 调用微信模板接口发送消息
	middleware.ResponseSuccess(c,dto.SmsResponse{
		Code: randCode,
		Msg: fmt.Sprintf("验证码为:%d,请勿告知他人",randCode),
	})
}
