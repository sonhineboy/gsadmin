package main

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/app"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/global"
	_ "github.com/sonhineboy/gsadmin/service/router"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("GetWd err:%v", err))
	}
	global.GAD_APP_PATH = dir + string(os.PathSeparator)
	app.Start()
	//自动迁移开始
	db, _ := global.Db.DB()

	amErr := global.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.AdminUser{},
		&models.AdminMenu{},
		&models.MenuApiList{},
		&models.Role{},
		&models.OperationLog{},
		&models.News{},
	)
	if amErr != nil {
		fmt.Println(amErr)
	}
	//自动迁移结束

	err = global.GAD_R.Run(global.Config.App.Port)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = db.Close()
	}()
}
