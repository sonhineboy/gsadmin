package router

import (
	"ginedu2/service/app/controllers/system"
	"ginedu2/service/app/middelware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SystemApiInit(r *gin.RouterGroup) {
	r.Use(middelware.JWTAuth(), middelware.Permission())
	rUser := r.Group("/user")
	{
		rUser.POST("/add", system.UserController{}.Add)
		rUser.POST("/update", system.UserController{}.Up)
		rUser.POST("/del", system.UserController{}.Del)
	}

	rSystem := r.Group("/system")
	{
		rSystem.POST("/menu/add", system.MenuController{}.Add)
		rSystem.POST("/menu/up", system.MenuController{}.Update)
		rSystem.GET("/menu/list", system.MenuController{}.All)
		rSystem.POST("/menu/dels", system.MenuController{}.Del)
		rSystem.GET("/role/list", system.RoleController{}.List)
		rSystem.POST("/role/add", system.RoleController{}.Add)
		rSystem.POST("/role/up", system.RoleController{}.Up)
		rSystem.POST("/role/del", system.RoleController{}.Del)
		rSystem.POST("/role/upMenu", system.RoleController{}.RoleUpMenu)
		rSystem.GET("/user/list", system.UserController{}.List)
		rSystem.POST("/common/upload", system.NewCommonController().UpLoad)
	}

	r.GET("/system/menu/my/:version", system.MenuController{}.MenuPermissions)

	r.GET("/demo/ver", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})
}
