package initialize

import (
	"log"
	"time"

	"github.com/gongxianjin/xcent-common/gorm"
	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/model"
)

func InitSysApi() {
	//获取链接池
	dbpool, err := lib.GetGormPool("default")
	if err != nil {
		log.Fatal(err)
	}
	if dbpool.Where("id IN (?)", []int64{1, 67}).Find(&[]model.SysApi{}).RowsAffected == 2 {
		log.Fatal("sys_apis表的初始数据已存在!")
	}
	db := dbpool.Begin()
	traceCtx := lib.NewTrace()
	//设置trace信息
	db = db.SetCtx(traceCtx)
	t1 := model.SysApi{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/base/login", "用户登录", "base", "POST"}

	Apis := model.SysApi{
	{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/base/login", "用户登录", "base", "POST"},
	}
	if err := dbpool.Debug().Save(&t1).Error; err != nil { // 遇到错误时回滚事务
		db.Rollback()
		log.Fatal(err)
	}
}
