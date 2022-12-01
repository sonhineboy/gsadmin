package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/event"
	"github.com/sonhineboy/gsadmin/service/app/listener"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/router"
	"github.com/sonhineboy/gsadmin/service/src"
	"golang.org/x/time/rate"
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
	global.Limiter = rate.NewLimiter(global.Config.Rate.Limit, global.Config.Rate.Burst)
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
	dsn := global.Config.Db.User + ":" + global.Config.Db.PassWord + "@tcp(" + global.Config.Db.Host + ":" + global.Config.Db.Port + ")/" + global.Config.Db.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
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
