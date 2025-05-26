package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/initialize"
	"github.com/sonhineboy/gsadmin/service/router"
	"golang.org/x/time/rate"
	"os"
)

func Start() {
	global.SuperAdmin = "administrator"
	global.GsE = gin.Default()
	global.Config = initialize.ConfigInit(global.GsAppPath)
	loadObject()
	router.RouteInit(global.GsE)
}

func TestLoad() {
	dir, err := os.Getwd()
	if err != nil {
	}
	global.GsAppPath = dir + "/../"
	global.Config = initialize.ConfigInit(global.GsAppPath)
	loadObject()

}

func loadObject() {
	global.Db = initialize.DbInit(global.Config)
	initialize.AutoMigrate(global.Db)
	global.EventDispatcher = initialize.EventInit()
	global.Limiter = rate.NewLimiter(global.Config.Rate.Limit, global.Config.Rate.Burst)
	global.Logger = initialize.ZapInit(global.Config)
	global.ValidatorManager = initialize.InitValidator()
}

func DiyDefer() {
	initialize.DbClose(global.Db)
	initialize.ZapSync(global.Logger)
}
