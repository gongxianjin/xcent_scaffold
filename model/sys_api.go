package model

import "github.com/gongxianjin/xcent-common/gorm"

type SysApi struct {
	gorm.Model
	Path        string `json:"path" gorm:"comment:api路径"`
	Description string `json:"description" gorm:"comment:api中文描述"`
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`
	Method      string `json:"method" gorm:"default:'POST'" gorm:"comment:方法"`
}

type User struct {
	gorm.Model
	Friends []*User `gorm:"many2many:user_friends;foreignKey:FriendID;"` 
}

type User_Friends struct{
	FriendID int
}

// type Language struct {
//   gorm.Model
// 	Name string
// 	UserRefer uint
// }

// type User struct {
// 	gorm.Model
// 	Languages []Language `gorm:"many2many:user_languages;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;JoinReferences:UserRefer"`
// 	Refer    uint
// }
