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
		model.User{}, 
		model.User_Friends{},
		model.SysBaseMenuParameter{},
		model.SysAuthority{},
		model.SysBaseMenu{},
		model.SysUser{},
		)
	log.Println("register table success")
}
