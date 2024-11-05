package models

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/global"
)

type Article struct {
	Title        string `gorm:"column:title;type:varchar(100);NOT NULL" json:"title"`
	Cid          uint64 `gorm:"column:cid;type:bigint(20) unsigned;NOT NULL" json:"cid"`
	Desc         string `gorm:"column:desc;type:varchar(200)" json:"desc"`
	Content      string `gorm:"column:content;type:longtext" json:"content"`
	Img          string `gorm:"column:img;type:varchar(100)" json:"img"`
	CommentCount int64  `gorm:"column:comment_count;type:bigint(20);default:0;NOT NULL" json:"comment_count"`
	ReadCount    int64  `gorm:"column:read_count;type:bigint(20);default:0;NOT NULL" json:"read_count"`
	global.GsModel
}

func (m *Article) TableName() string {
	return fmt.Sprint(global.Config.Db.TablePrefix, "article")
}
