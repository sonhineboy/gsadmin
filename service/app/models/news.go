package models

import "github.com/sonhineboy/gsadmin/service/global"

type News struct {
	global.GAD_MODEL
	Title          string    `gorm:"column:title;type:varchar(255);not null;comment:标题;" json:"title"`
	Author         string    `gorm:"column:author;type:varchar(255);not null;comment:作者;" json:"author"`
	Content        string    `gorm:"column:content;type:text;not null;comment:内容;" json:"content"`
	
}

func (m *News) TableName() string {
	return "news"
}
