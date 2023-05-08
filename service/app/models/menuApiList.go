package models

import (
	"github.com/sonhineboy/gsadmin/service/global"
)

type MenuApiList struct {
	global.GAD_MODEL
	Code   string    `gorm:"column:code;type:varchar(100);comment:关键字" json:"code"`
	Url    string    `gorm:"column:url;type:varchar(100);comment:地址" json:"url"`
	MenuId uint      `gorm:"column:menu_id;type:int;" json:"menu_id"`
	Menu   AdminMenu `gorm:"foreignKey:id;references:menu_id"`
}
