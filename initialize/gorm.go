package initialize

import (
	"log"
	"github.com/gongxianjin/xcent_scaffold/model"
	"github.com/gongxianjin/xcent-common/gorm"
)

// MysqlTables 注册数据库表专用
func MysqlTables(db *gorm.DB) {
	db.AutoMigrate(
		model.SysApi{},
		model.SysBaseMenuParameter{},
		model.SysAuthority{},
		model.Sys_Data_Authority_Id{},
		model.SysBaseMenu{},
		model.SysUser{},
	)
	log.Println("register table success")
}

func InitMysqlData(db *gorm.DB) {
	InitSysApi(db)
	InitSysUser(db)
}
