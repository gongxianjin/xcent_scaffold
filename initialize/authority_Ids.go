package initialize

import (
	"github.com/gongxianjin/xcent_scaffold/model"
	"log"

	"github.com/gongxianjin/xcent-common/gorm"
)

//type SysDataAuthorityId struct {
//	SysAuthorityAuthorityId    string
//	DataAuthorityIdAuthorityId string
//}

var DataAuthorityId = []model.SysDataAuthorityId{
	{"888", "888"},
	{"888", "8881"},
	{"888", "9528"},
	{"9528", "8881"},
	{"9528", "9528"},
}

func InitSysDataAuthorityId(db *gorm.DB) {
	if db.Where("sys_authority_authority_id IN (?) ", []string{"888", "9528"}).Find(&[]model.SysDataAuthorityId{}).RowsAffected == 5 {
		log.Println("sys_data_authority_id表的初始数据已存在!")
		return
	}
	db = db.Begin() 
		//去掉sys_data_authority_id中sys_authority_authority_id索引
	//if err := db.Exec("ALTER TABLE `sys_data_authority_id` DROP INDEX `sys_authority_authority_id`;").Error; err != nil {
	//	log.Println("删除索引sys_authority_authority_id失败!")
	//	return
	//}
	//加上主键sys_data_authority_id中data_authority_id_authority_id
	//if err := db.Exec("ALTER TABLE `sys_data_authority_id` MODIFY COLUMN `data_authority_id_authority_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL AFTER `sys_authority_authority_id`,DROP PRIMARY KEY,ADD PRIMARY KEY (`sys_authority_authority_id`, `data_authority_id_authority_id`) USING BTREE;").Error; err != nil {
	//	log.Println("创建主键sys_authority_authority_id失败!")
	//	return
	//}
	for _, api := range DataAuthorityId {
		if err := db.Debug().Save(&api).Error; err != nil { // 遇到错误时回滚事务
			db.Rollback()
			log.Fatal(err)
		}
	}
	db.Commit()
	log.Println("init dataAuthorityId success")
}
