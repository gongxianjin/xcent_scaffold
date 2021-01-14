package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gongxianjin/xcent_scaffold/initialize"

	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/router"
)


// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	//获取链接池
	db, err := lib.GetGormPool("default")
	if err != nil {
		log.Fatal(err)
	}
	initialize.MysqlTables(db)
	initialize.InitMysqlData(db)
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
