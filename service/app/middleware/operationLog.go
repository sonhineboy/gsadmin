package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/global"
)

func OperationLog() gin.HandlerFunc {

	return func(c *gin.Context) {

		//ip
		println(c.ClientIP())

		//路径
		println(c.Request.URL.Path)

		var where = make(map[string]interface{})

		var d models.MenuApiList

		db := global.Db.Model(&models.MenuApiList{})
		where["url"] = c.Request.URL.Path
		db.Preload("Menu").Where(where).First(&d)

		println(d.Menu.Meta["title"].(string))

		//get/post
		method := c.Request.Method
		println(method)

		//用户信息
		claims, ok := repositorys.GetCustomClaims(c)
		if ok == true {
			println(claims.Name)
		}

		//参数
		if method == "GET" {
			b, _ := json.Marshal(c.Request.URL.Query())
			println(string(b))
		}

		if method == "POST" {
			s, _ := c.GetRawData()
			println(string(s))
		}

		c.Next()
	}

}
