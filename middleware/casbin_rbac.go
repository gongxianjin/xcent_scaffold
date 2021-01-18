package middleware

import (
	"fmt"
 
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent_scaffold/service"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		fmt.Println(claims)
		waitUse := claims.(*request.CustomClaims)
		fmt.Println(waitUse)
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
		if success {
			c.Next()
		} else {
			ResponseError(c, 501, err)
			c.Abort()
			return
		}
	}
}
