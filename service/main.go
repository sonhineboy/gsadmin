package main

import (
	"fmt"
	"ginedu2/service/app"
	"ginedu2/service/app/models"
	"ginedu2/service/global"
	_ "ginedu2/service/router"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
	}
	global.GAD_APP_PATH = dir + "/"
	app.Start()

	//自动迁移开始
	db, _ := global.Db.DB()
	defer db.Close()
	amErr := global.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.AdminUser{},
		&models.AdminMenu{},
		&models.MenuApiList{},
		&models.Role{})
	if amErr != nil {
		fmt.Println(amErr)
	}
	//自动迁移结束

	err = global.GAD_R.Run(global.Config.App.Port)
	if err != nil {
		panic(err)
	}
}
