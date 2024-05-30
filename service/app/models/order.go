package models

import "github.com/sonhineboy/gsadmin/service/global"

type Order struct {
	*global.GAD_MODEL
	UserName string `gorm:"column:user_name;index:user_name,class:fulltext;comment:用户名;" json:"user_name" binding:"required"`
	Age      int    `gorm:"column:age;default:0;comment:年龄;" json:"age" binging:"required"`
}

func (m *Order) TableName() string {
	return "order"
}
