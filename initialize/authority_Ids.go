package initialize

import (
	"github.com/gongxianjin/xcent-common/gorm"
	"github.com/gongxianjin/xcent-common/lib"
	"log"
)

type SysDataAuthorityId struct {
	SysAuthorityAuthorityId    string
	DataAuthorityIdAuthorityId string
}

var DataAuthorityId = []SysDataAuthorityId{
	{"888", "888"},
	{"888", "8881"},
	{"888", "9528"},
	{"9528", "8881"},
	{"9528", "9528"},
}

func InitSysDataAuthorityId(db *gorm.DB) {
	if db.Where("sys_authority_authority_id IN (?) ", []string{"888", "9528"}).Find(&[]SysDataAuthorityId{}).RowsAffected == 5 {
		log.Println("sys_data_authority_id表的初始数据已存在!")
		return
	}
	db = db.Begin()
	traceCtx := lib.NewTrace()
	//设置trace信息
	db = db.SetCtx(traceCtx)
	for _,api := range DataAuthorityId {
		if err := db.Debug().Save(&api).Error; err != nil { // 遇到错误时回滚事务
			db.Rollback()
			log.Fatal(err)
		}
	}
	db.Commit()
	log.Println("init dataAuthorityId success")
}
