package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/initialize"
	"github.com/sonhineboy/gsadmin/service/pkg/ctx"
	"github.com/sonhineboy/gsadmin/service/router"
	"golang.org/x/time/rate"
	"os"
)

func Start() {
	global.SuperAdmin = "administrator"
	global.GsR = gin.Default()
	global.Config = initialize.ConfigInit(global.GsAppPath)
	loadObject()
	router.RouteInit(global.GsR)
	initialize.PluginInit(ctx.NewDefaultAppCtx())

}

func TestLoad() {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("test Load getwd err %w", err))
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
}

func DiyDefer() {
	initialize.DbClose(global.Db)
	initialize.ZapSync(global.Logger)
}
