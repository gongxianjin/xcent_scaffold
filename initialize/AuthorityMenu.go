package initialize

import (
	"log"

	"github.com/gongxianjin/xcent-common/gorm"
)

func InitAuthorityMenu(db *gorm.DB) {
	if err := db.Exec("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menu`.`id` AS `id`,`sys_base_menu`.`created_at` AS `created_at`, `sys_base_menu`.`updated_at` AS `updated_at`, `sys_base_menu`.`deleted_at` AS `deleted_at`, `sys_base_menu`.`menu_level` AS `menu_level`,`sys_base_menu`.`parent_id` AS `parent_id`,`sys_base_menu`.`path` AS `path`,`sys_base_menu`.`name` AS `name`,`sys_base_menu`.`hidden` AS `hidden`,`sys_base_menu`.`component` AS `component`, `sys_base_menu`.`title`  AS `title`,`sys_base_menu`.`icon` AS `icon`,`sys_base_menu`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menu`.`keep_alive` AS `keep_alive`,`sys_base_menu`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menu` on ((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menu`.`id`)))").Error; err != nil {
		log.Println("authority_menu视图已存在!")
		return
	}
	log.Println("authority_menu视图创建成功!")
}
