package models

import "github.com/sonhineboy/gsadmin/service/global"

type UserMember struct {
	global.GAD_MODEL
	NickName       string    `gorm:"column:nick_name;type:varchar(255);not null;comment:昵称;" json:"nick_name"`
	RealName       string    `gorm:"column:real_name;type:varchar(255);comment:真实姓名;" json:"real_name"`
	Age            int32     `gorm:"column:age;default:0;comment:0 未知;" json:"age"`
	Status         int8      `gorm:"column:status;index:status;default:1;comment:1 正常，2 禁用 ;" json:"status"`
	Online         string    `gorm:"column:online;type:varchar(255);index:online;default:0;comment:不在线;" json:"online"`
	
}

func (m *UserMember) TableName() string {
	return "user_member"
}
