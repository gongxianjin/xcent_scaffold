package initialize

import (
	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/model"
	"log"
)

// MysqlTables 注册数据库表专用
func MysqlTables() {
	lib.GORMDefaultPool.AutoMigrate(
		model.SysApi{},
		model.SysBaseMenuParameter{},
		model.SysAuthority{},
		model.SysBaseMenu{},
		model.SysUser{},
		)
	log.Println("register table success")
}
