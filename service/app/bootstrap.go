package app

import (
	"fmt"
	"ginedu2/service/app/event"
	"ginedu2/service/app/listener"
	"ginedu2/service/global"
	"ginedu2/service/router"
	"ginedu2/service/src"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"time"
)

func Start() {
	global.SuperAdmin = "administrator"
	global.GAD_R = gin.Default()

	loadConfig()
	loadObject()
	global.EventDispatcher = InitEvent()
	router.RouteInit()
}

func loadConfig() {
	configFile, err := ioutil.ReadFile(global.GAD_APP_PATH + "config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err2 := yaml.Unmarshal(configFile, &global.Config)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
}

func loadObject() {
	//初始化数据
	initMysql()
}

func initMysql() {

	dsn := "root:@tcp(" + global.Config.Db.Host + ":" + global.Config.Db.Port + ")/" + global.Config.Db.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.Config.Db.TablePrefix,
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err.Error())
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(global.Config.Db.MaxOpenConns)
	sqlDb.SetMaxIdleConns(global.Config.Db.MaxIdleConns)
	sqlDb.SetConnMaxIdleTime(time.Hour)
	global.Db = db

}

//初始化事件
func InitEvent() *src.EventDispatcher {
	EventDispatcher := src.NewDispatcher()
	EventDispatcher.Register(event.TestEvent{}.GetEventName(), listener.NewTestListener())
	EventDispatcher.Register(event.LoginEvent{}.GetEventName(), listener.NewTestListener())
	return EventDispatcher
}
