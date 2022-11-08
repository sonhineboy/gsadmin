package models

import "ginedu2/service/global"

type Role struct {
	global.GAD_MODEL
	Alias  string      `gorm:"type:varchar(50);column:alias;" json:"alias"`
	Label  string      `gorm:"type:varchar(100);column:label;" json:"label"`
	Remark string      `gorm:"type:varchar(255);column:remark" json:"remark"`
	Sort   int         `gorm:"type:int(3);column:sort" json:"sort"`
	Status *int        `json:"status" gorm:"type:int(3);column:status"`
	Menus  []AdminMenu `json:"menus" gorm:"many2many:role_menu"`
}
