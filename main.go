package main

import (
	"github.com/gongxianjin/xcent_scaffold/initialize"
	"os"
	"os/signal"
	"syscall"

	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/router"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	initialize.MysqlTables()
	//initialize.InitMysqlData()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
