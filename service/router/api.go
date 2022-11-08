package router

import (
	"ginedu2/service/app/controllers"
	"ginedu2/service/app/middelware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiInit(r *gin.RouterGroup) {
	r.Use(middelware.JWTAuth(), middelware.Permission())
	rUser := r.Group("/user")
	{
		rUser.POST("/add", controllers.UserController{}.Add)
		rUser.POST("/update", controllers.UserController{}.Up)
		rUser.POST("/del", controllers.UserController{}.Del)
	}

	rSystem := r.Group("/system")
	{
		rSystem.POST("/menu/add", controllers.MenuController{}.Add)
		rSystem.POST("/menu/up", controllers.MenuController{}.Update)
		rSystem.GET("/menu/list", controllers.MenuController{}.All)
		rSystem.POST("/menu/dels", controllers.MenuController{}.Del)
		rSystem.GET("/role/list", controllers.RoleController{}.List)
		rSystem.POST("/role/add", controllers.RoleController{}.Add)
		rSystem.POST("/role/up", controllers.RoleController{}.Up)
		rSystem.POST("/role/del", controllers.RoleController{}.Del)
		rSystem.POST("/role/upMenu", controllers.RoleController{}.RoleUpMenu)
		rSystem.GET("/user/list", controllers.UserController{}.List)
		rSystem.POST("/common/upload", controllers.NewCommonController().UpLoad)
	}

	r.GET("/system/menu/my/:version", controllers.MenuController{}.MenuPermissions)

	r.GET("/demo/ver", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})
}
