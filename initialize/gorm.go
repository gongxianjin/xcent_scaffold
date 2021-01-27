package initialize

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
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
		model.JwtBlacklist{},
		model.SysBaseMenu{},
		model.SysUser{},
		model.SysOperationRecord{},
		gormadapter.CasbinRule{},
	)
	log.Println("register table success")
}

func InitMysqlData(db *gorm.DB) {
	InitSysApi(db)
	InitSysUser(db)
	InitCasbinModel(db)
	InitSysAuthority(db)
	InitSysBaseMenus(db)
	InitSysAuthorityMenus(db)
	InitAuthorityMenu(db)
	InitSysDataAuthorityId(db)
}
