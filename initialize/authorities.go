package initialize

import (
	"github.com/gongxianjin/xcent-common/gorm"
	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/model"
	"log"
	"time"
)

var Authorities = []model.SysAuthority{
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "普通用户", ParentId: "0"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: "普通用户子角色", ParentId: "888"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: "测试角色", ParentId: "0"},
}

func InitSysAuthority(db *gorm.DB) {
	if db.Where("authority_id IN (?) ", []string{"888", "9528"}).Find(&[]model.SysAuthority{}).RowsAffected == 2 {
		log.Println("sys_authorities表的初始数据已存在!")
		return
	}
	db = db.Begin()
	traceCtx := lib.NewTrace()
	//设置trace信息
	db = db.SetCtx(traceCtx)
	for _,api := range Authorities {
		if err := db.Debug().Save(&api).Error; err != nil { // 遇到错误时回滚事务
			db.Rollback()
			log.Fatal(err)
		}
	}
	db.Commit()
	log.Println("init authority success")
}
