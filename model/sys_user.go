package model

import "github.com/gongxianjin/xcent-common/gorm"

type SysUser struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"type:varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}
