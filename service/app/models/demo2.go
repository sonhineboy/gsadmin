package models

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/global"
)

type Demo2 struct {
	global.GAD_MODEL
	Name           string    `gorm:"column:name;type:varchar(255);not null;comment:名字;" json:"name"`
	Age            int8      `gorm:"column:age;index:age;not null;default:10;" json:"age"`
	
}

func (m *Demo2) TableName() string {
	return fmt.Sprint(global.Config.Db.TablePrefix, "demo2")
}
