package models

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/global"
)

type News struct {
	global.GsModel
	Title   string `gorm:"column:title;type:varchar(255);not null;comment:标题;" json:"title"`
	Author  string `gorm:"column:author;type:varchar(255);not null;comment:作者;" json:"author"`
	Content string `gorm:"column:content;type:text;not null;comment:内容;" json:"content"`
	Image   string `gorm:"column:image;comment:缩略图;" json:"image"`
}

func (m *News) TableName() string {
	return fmt.Sprint(global.Config.Db.TablePrefix, "news")
}
