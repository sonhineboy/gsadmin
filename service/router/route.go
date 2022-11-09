package router

import (
	"fmt"
	"ginedu2/service/app/controllers/system"
	"ginedu2/service/global"
)

func RouteInit() {

	commonController := system.NewCommonController()

	fmt.Println(commonController.GetFileBasePath())

	global.GAD_R.GET("sss", system.Demo)
	global.GAD_R.GET("/api/common/captcha/img/:id/:w/:h", commonController.CaptchaImage)
	global.GAD_R.GET("/api/common/captcha/info", commonController.CaptchaInfo)
	global.GAD_R.POST("/api/user/login", system.UserController{}.Login)
	global.GAD_R.Static("/api/system/common/file", commonController.GetFileBasePath())
	global.GAD_R.GET("/api/system/dept/list", system.DeptList)
	global.GAD_R.GET("/api/demo/page", system.DemoUser)

	r := global.GAD_R.Group("api")
	{
		ApiInit(r)
	}
}
