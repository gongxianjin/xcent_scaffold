package model

import (
	"github.com/gongxianjin/xcent-common/gorm"
)

type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
