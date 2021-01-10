package initialize

import (
	"github.com/gongxianjin/xcent-common/gorm"
	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/model"
	"log"
	"time"
	uuid "github.com/satori/go.uuid"
)

var Users = []model.SysUser{
	{MODEL:gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Phone: "1311111111",Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", HeaderImg: "http://www.xcentiot.com/Public/images/logo_1.png", AuthorityId: "888"},
	{MODEL:gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Phone: "1511111111",Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "xcentUser", HeaderImg: "http://www.xcentiot.com/Public/images/logo_1.png", AuthorityId: "9528"},
}

func InitSysUser(db *gorm.DB) {
	if db.Where("id IN (?)", []int{1, 2}).Find(&[]model.SysUser{}).RowsAffected == 2 {
		log.Fatal("sys_apis表的初始数据已存在!")
	}
	db = db.Begin()
	traceCtx := lib.NewTrace()
	//设置trace信息
	db = db.SetCtx(traceCtx)
	for _,api := range Users {
		if err := db.Debug().Save(&api).Error; err != nil { // 遇到错误时回滚事务
			db.Rollback()
			log.Fatal(err)
		}
	}
	db.Commit()
	log.Println("init User success")
}
