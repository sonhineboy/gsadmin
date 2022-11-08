package models

import (
	"ginedu2/service/global"
)

type MenuApiList struct {
	global.GAD_MODEL
	Code   string `gorm:"column:code;type:varchar(100);comment:关键字" json:"code"`
	Url    string `gorm:"column:url;type:varchar(100);comment:地址" json:"url"`
	MenuId uint   `gorm:"column:menu_id;type:int;" json:"menu_id"`
}
