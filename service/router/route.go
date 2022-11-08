package router

import (
	"fmt"
	"ginedu2/service/app/controllers"
	"ginedu2/service/global"
)

func RouteInit() {

	commonController := controllers.NewCommonController()

	fmt.Println(commonController.GetFileBasePath())

	global.GAD_R.GET("sss", controllers.Demo)
	global.GAD_R.GET("/api/common/captcha/img/:id/:w/:h", commonController.CaptchaImage)
	global.GAD_R.GET("/api/common/captcha/info", commonController.CaptchaInfo)
	global.GAD_R.POST("/api/user/login", controllers.UserController{}.Login)
	global.GAD_R.Static("/api/system/common/file", commonController.GetFileBasePath())
	global.GAD_R.GET("/api/system/dept/list", controllers.DeptList)
	global.GAD_R.GET("/api/demo/page", controllers.DemoUser)

	r := global.GAD_R.Group("api")
	{
		ApiInit(r)
	}
}
