package main

import (
	"context"
	"fmt"
	"github.com/sonhineboy/gsadmin/service/app"
	"github.com/sonhineboy/gsadmin/service/global"
	_ "github.com/sonhineboy/gsadmin/service/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	run()
}

// run 开始监听并启动web服务
func run() {
	svr := &http.Server{
		Addr:    global.Config.App.Port,
		Handler: global.GsE,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil {
			global.Logger.Errorf("listen: %s\n", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	<-sigs
	global.Logger.Infof("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		global.Logger.Errorf("stutdown err %v", err)
	}
	global.Logger.Infof("shutdown-->ok")
}
