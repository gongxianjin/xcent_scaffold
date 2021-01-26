package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/model/request"
	"github.com/gongxianjin/xcent_scaffold/service"
	"github.com/gongxianjin/xcent_scaffold/model/response"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims") 
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		fmt.Println(obj)
		// 获取请求方法
		act := c.Request.Method
		fmt.Println(act)
		// 获取用户的角色
		// Session := sessions.Default(c)
		// sub := Session.Get("user_id")
		sub := waitUse.AuthorityId
		e := service.Casbin()
		// 判断策略中是否存在
		success, err := e.Enforce(sub, obj, act)
		fmt.Println(sub,obj,act,success)
		if success {
			c.Next()
		} else {
			log.Printf("rabc role:%v",err)
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return 
		}
	}
}
