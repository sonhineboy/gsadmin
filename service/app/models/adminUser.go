package models

import "github.com/sonhineboy/gsadmin/service/global"

type AdminUser struct {
	global.GAD_MODEL
	Nickname string `gorm:"column:nickname;type:varchar(255);comment:昵称" json:"nickname"`
	RealName string `gorm:"column:real_name;type:varchar(255);comment:真实名称" json:"real_name"`
	Password string `gorm:"column:password;type:varchar(255)" json:"password"`
	Email    string `gorm:"column:email;type:varchar(255)" json:"email"`
	Name     string `gorm:"uniqueIndex;type:varchar(100);default:" json:"name"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);default:''" json:"avatar"`
	Roles    []Role `json:"group" gorm:"many2many:user_role"`
}
