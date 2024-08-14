package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/controllers/system"
	"github.com/sonhineboy/gsadmin/service/app/middleware"
)

func RouteInit(e *gin.Engine) {
	e.Use(middleware.Limiter(), middleware.Event())
	e.GET("/api/common/captcha/img/:id/:w/:h", ApiControllers.CommonController.CaptchaImage)
	e.GET("/api/common/captcha/info", ApiControllers.CommonController.CaptchaInfo)
	e.GET("/api/common/version", ApiControllers.CommonController.GetVersion)
	e.POST("/api/user/login", ApiControllers.UserController.Login)
	e.Static("/api/system/common/file", ApiControllers.CommonController.GetFileBasePath())
	e.GET("/api/system/dept/list", system.DeptList)
	e.GET("/api/demo/page", system.DemoUser)
	e.POST("/api/demo/order", system.OrderDemo)

	r := e.Group("api")
	{
		SystemApiInit(r)
	}

}
