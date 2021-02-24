package model

import (
	"time"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	BaseModel
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`
	Username    string       `json:"user_name" gorm:"type:varchar(20);not null;comment:用户登录名"`
	Password    string       `json:"-"  gorm:"size:255;not null;comment:用户登录密码"`
	Phone       string       `json:"phone" gorm:"type:varchar(11);not null;unique;comment:用户手机号"`
	Email       string       `json:"email" gorm:"comment:邮箱号"`
	NickName    string       `json:"nickName" gorm:"default:'系统用户';comment:用户昵称" `
	HeaderImg   string       `json:"headerImg" gorm:"comment:用户头像"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}

type BaseModel struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
