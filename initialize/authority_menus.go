package initialize

import (
	"log"

	"github.com/gongxianjin/xcent-common/gorm"
)

type SysAuthorityMenus struct {
	SysAuthorityAuthorityId string
	SysBaseMenuId          uint
}

var AuthorityMenus = []SysAuthorityMenus{
	{"888", 1},
	{"888", 2},
	{"888", 3},
	{"888", 4},
	{"888", 5},
	{"888", 6},
	{"888", 7},
	{"888", 8},
	{"888", 9},
	{"888", 10},
	{"888", 11},
	{"888", 12},
	{"888", 13},
	{"888", 14},
	{"888", 15},
	{"888", 16},
	{"888", 17},
	{"888", 18},
	{"888", 19},
	{"888", 20},
	{"888", 21},
	{"888", 22},
	{"888", 23},
	{"888", 24},
	{"888", 25},
	{"888", 26},
	{"888", 27},
	{"888", 28},
	{"888", 29},
	{"888", 30},
	{"888", 31},
	{"888", 32},
	{"888", 33},
	{"8881", 1},
	{"8881", 2},
	{"8881", 8},
	{"8881", 17},
	{"8881", 18},
	{"8881", 19},
	{"8881", 20},
	{"9528", 1},
	{"9528", 2},
	{"9528", 3},
	{"9528", 4},
	{"9528", 5},
	{"9528", 6},
	{"9528", 7},
	{"9528", 8},
	{"9528", 9},
	{"9528", 10},
	{"9528", 11},
	{"9528", 12},
	{"9528", 13},
	{"9528", 14},
	{"9528", 15},
	{"9528", 17},
	{"9528", 18},
	{"9528", 19},
	{"9528", 20},
}

func InitSysAuthorityMenus(db *gorm.DB) {
	if db.Where("sys_authority_authority_id IN (?) ", []string{"888", "8881", "9528"}).Find(&[]SysAuthorityMenus{}).RowsAffected >= 59 {
		log.Println("sys_authority_menus表的初始数据已存在!")
		return
	}
	db = db.Begin()
	//去掉sys_data_authority_id中sys_authority_authority_id索引
	if err := db.Exec("ALTER TABLE`sys_authority_menus` DROP INDEX `sys_authority_authority_id`;").Error; err != nil {
		log.Println("删除索引sys_authority_authority_id失败!")
		return
	}
	for _, api := range AuthorityMenus {
		if err := db.Create(&api).Error; err != nil { // 遇到错误时回滚事务
			db.Rollback()
			log.Fatal(err)
		}
	}
	db.Commit()
	log.Println("init authorityMenus success")
}
