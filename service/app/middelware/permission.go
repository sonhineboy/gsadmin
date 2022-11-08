package middelware

import (
	"ginedu2/service/app/models"
	"ginedu2/service/app/repositorys"
	"ginedu2/service/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Permission() gin.HandlerFunc {

	return func(c *gin.Context) {

		var apiList []models.MenuApiList
		err := repositorys.SystemMenuRepository{}.GetApiList(c, &apiList)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  "无权限访问!",
			})
			c.Abort()
			return
		}

		isAllow := false
		for _, api := range apiList {
			if api.Url == c.Request.URL.Path {
				isAllow = true
			}
		}
		Claims, _ := repositorys.SystemMenuRepository{}.GetCustomClaims(c)

		if !isAllow && !global.IsSuperAdmin(Claims.Roles, global.SuperAdmin) {
			c.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  "无权限访问!",
			})
			c.Abort()
			return
		}
		c.Next()

	}
}
