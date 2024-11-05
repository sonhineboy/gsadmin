package models

import "github.com/sonhineboy/gsadmin/service/global"

type OperationLog struct {
	global.GsModel
	UserId   uint   `gorm:"column:user_id;type:int(11);comment:用户ID" json:"user_id"`
	UrlPath  string `gorm:"column:user_path;type:varchar(100);comment:访问路径" json:"url_path"`
	Ip       string `gorm:"column:ip;type:varchar(50);comment:IP" json:"ip"`
	Method   string `gorm:"column:method;type:varchar(50);comment:请求方式" json:"method"`
	PathName string `gorm:"column:path_name;type:varchar(100);comment:请求名称" json:"path_name"`
	DoData   string `gorm:"column:do_data;type:text;comment:处理数据;default:null" json:"do_data"`
	UserName string `gorm:"column:user_name;type:varchar(40);comment:用户名;default:未知;NOT NULL" json:"user_name"`
}
