package initialize

import (
	"log"

	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/model"
)

// MysqlTables 注册数据库表专用
func MysqlTables() {
	lib.GORMDefaultPool.AutoMigrate(
		model.SysApi{},
		model.Test1{},
		model.SysBaseMenuParameter{},
		model.SysAuthority{},
		model.Sys_Data_Authority_Id{},
		model.SysBaseMenu{},
		model.SysUser{},
	)
	log.Println("register table success")
}

func InitMysqlData() {
	InitSysApi()
}
