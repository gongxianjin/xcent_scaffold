package initialize

import ( 
	"github.com/gongxianjin/xcent_scaffold/model"
	"log" 
	"gorm.io/gorm"
)

// MysqlTables 注册数据库表专用
func MysqlTables(db gorm.DB) {
 db.AutoMigrate(
		model.SysApi{},
		model.User{},
		model.SysBaseMenuParameter{},
		model.SysAuthority{},
		model.SysBaseMenu{},
		model.SysUser{},
		)
	log.Println("register table success")
}
