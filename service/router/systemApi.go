package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/middleware"
	"net/http"
)

func SystemApiInit(r *gin.RouterGroup) {

	r.Use(middleware.JWTAuth(), middleware.Permission())
	rUser := r.Group("/user")
	{
		rUser.POST("/add", ApiControllers.UserController.Add)
		rUser.POST("/update", ApiControllers.UserController.Up)
		rUser.POST("/del", ApiControllers.UserController.Del)
	}

	rSystem := r.Group("/system")
	{
		rSystem.POST("/menu/add", ApiControllers.MenuController.Add)
		rSystem.POST("/menu/up", ApiControllers.MenuController.Update)
		rSystem.GET("/menu/list", ApiControllers.MenuController.All)
		rSystem.POST("/menu/dels", ApiControllers.MenuController.Del)
		rSystem.GET("/role/list", ApiControllers.RoleController.List)
		rSystem.POST("/role/add", ApiControllers.RoleController.Add)
		rSystem.POST("/role/up", ApiControllers.RoleController.Up)
		rSystem.POST("/role/del", ApiControllers.RoleController.Del)
		rSystem.POST("/role/upMenu", ApiControllers.RoleController.RoleUpMenu)
		rSystem.GET("/user/list", ApiControllers.UserController.List)
		rSystem.POST("/common/upload", ApiControllers.CommonController.UpLoad)
	}

	r.GET("/system/menu/my/:version", ApiControllers.MenuController.MenuPermissions)

	r.GET("/demo/ver", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})
}
