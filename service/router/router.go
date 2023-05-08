package router

import (
	"github.com/sonhineboy/gsadmin/service/app/controllers/system"
	"github.com/sonhineboy/gsadmin/service/app/middleware"
	"github.com/sonhineboy/gsadmin/service/global"
)

func RouteInit() {

	global.GAD_R.Use(middleware.Limiter(), middleware.Event())
	global.GAD_R.GET("sss", system.Demo)
	global.GAD_R.GET("/api/common/captcha/img/:id/:w/:h", ApiControllers.CommonController.CaptchaImage)
	global.GAD_R.GET("/api/common/captcha/info", ApiControllers.CommonController.CaptchaInfo)
	global.GAD_R.POST("/api/user/login", ApiControllers.UserController.Login)
	global.GAD_R.Static("/api/system/common/file", ApiControllers.CommonController.GetFileBasePath())
	global.GAD_R.GET("/api/system/dept/list", system.DeptList)
	global.GAD_R.GET("/api/demo/page", system.DemoUser)

	r := global.GAD_R.Group("api")
	{
		SystemApiInit(r)
	}

}
