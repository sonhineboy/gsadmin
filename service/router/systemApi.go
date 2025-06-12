package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/controllers/genExample"
	"github.com/sonhineboy/gsadmin/service/app/controllers/system"
	"github.com/sonhineboy/gsadmin/service/app/middleware"
)

func SystemApiInit(r *gin.RouterGroup) {

	// 只验证token
	r.Use(middleware.JWTAuth(), middleware.OperationLog())
	{
		r.POST("/user/logout", ApiControllers.UserController.Logout)
	}
	// 查权限
	r.Use(middleware.Permission())
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

		rSystem.GET("/operationLog/list", ApiControllers.OperationLogController.List)

	}

	r.GET("/system/menu/my/:version", ApiControllers.MenuController.MenuPermissions)

	r.GET("/demo/ver", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})

	//gen
	gen := r.Group("/gen")
	{
		var genController system.GenController
		gen.GET("/tables", genController.GetTables)
		gen.GET("/fields", genController.TableFields)
		gen.Use(middleware.EnvCheck()).POST("/genCode", genController.GenCode)
	}

	//router gen start not delete

	//gen_news
	News := r.Group("/news")
	{
		var NewsController genExample.NewsController
		News.GET("/index", NewsController.Index)
		News.POST("/save", NewsController.Save)
		News.POST("/delete", NewsController.Delete)
		News.GET("/:id", NewsController.Get)
		News.POST("/edit/:id", NewsController.Edit)
	}

	//router gen end not delete

}
