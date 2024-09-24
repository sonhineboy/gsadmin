package main

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/app"
	"github.com/sonhineboy/gsadmin/service/global"
	_ "github.com/sonhineboy/gsadmin/service/router"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("GetWd err:%v", err))
	}
	global.GsAppPath = dir + string(os.PathSeparator)
	app.Start()

	defer func() {
		app.DiyDefer()
	}()
	err = global.GsR.Run(global.Config.App.Port)
	if err != nil {
		panic(err)
	}

}
