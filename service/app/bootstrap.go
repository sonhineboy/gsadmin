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
	global.GAD_R = gin.Default()
	global.Config = initialize.ConfigInit(global.GAD_APP_PATH)
	loadObject()
	router.RouteInit(global.GAD_R)
}

func TestLoad() {
	dir, err := os.Getwd()
	if err != nil {
	}
	global.GAD_APP_PATH = dir + "/../"
	global.Config = initialize.ConfigInit(global.GAD_APP_PATH)
	loadObject()
}

func loadObject() {
	global.Db = initialize.DbInit(global.Config)
	global.EventDispatcher = initialize.EventInit()
	global.Limiter = rate.NewLimiter(global.Config.Rate.Limit, global.Config.Rate.Burst)

}
