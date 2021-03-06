package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/controller"
	"github.com/gongxianjin/xcent_scaffold/docs"
	"github.com/gongxianjin/xcent_scaffold/middleware"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// programatically set swagger info
	docs.SwaggerInfo.Title = lib.GetStringConf("base.swagger.title")
	docs.SwaggerInfo.Description = lib.GetStringConf("base.swagger.desc")
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = lib.GetStringConf("base.swagger.host")
	docs.SwaggerInfo.BasePath = lib.GetStringConf("base.swagger.base_path")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	////demo
	//v1 := router.Group("/demo")
	//v1.Use(middleware.RecoveryMiddleware(), middleware.RequestLog(), middleware.IPAuthMiddleware(), middleware.TranslationMiddleware(),middleware.CasbinHandler())
	//{
	//	controller.DemoRegister(v1)
	//}

	//非登陆接口
	store := sessions.NewCookieStore([]byte("secret"))
	apiNormalGroup := router.Group("")
	apiNormalGroup.Use(sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		InitBaseRouter(apiNormalGroup)
	}

	//登陆接口
	apiAuthGroup := router.Group("")
	apiAuthGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.TranslationMiddleware(),
		middleware.JWTAuth(),
		middleware.CasbinHandler(),
	)
	{
		InitUserRouter(apiAuthGroup)										// 注册用户路由
		InitCasbinRouter(apiAuthGroup)                // 权限相关路由
		InitAuthorityRouter(apiAuthGroup)					  // 注册角色路由
		InitMenuRouter(apiAuthGroup)                  // 注册menu路由
    InitApiRouter(apiAuthGroup)                       // 注册功能api路由
	}

	//demo
	v1 := router.Group("/demo")
	v1.Use(
		middleware.RecoveryMiddleware(), 
		middleware.RequestLog(),
		middleware.IPAuthMiddleware(), 
		middleware.TranslationMiddleware(),
		middleware.JWTAuth(), 
		middleware.CasbinHandler())
	{
		controller.DemoRegister(v1)
	}

	return router
}
