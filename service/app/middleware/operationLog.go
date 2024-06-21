package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/global"
	"strings"
)

func OperationLog() gin.HandlerFunc {

	return func(c *gin.Context) {

		cCp := c.Copy()
		go func() {
			contentType := cCp.GetHeader("Content-Type")
			var (
				doData []byte
				log    models.OperationLog
			)
			method := c.Request.Method
			//参数
			if method == "GET" {
				doData, _ = json.Marshal(cCp.Request.URL.Query())
			}
			if method == "POST" {

				if strings.Contains(contentType, "multipart/form-data") {
					doData = []byte("图片上传")
				} else {
					doData, _ = cCp.GetRawData()
				}
			}

			claims, ok := repositorys.GetCustomClaims(c)
			if ok == true {
				log.UserId = claims.Id
				log.UserName = claims.Name
			} else {
				log.UserId = 0
			}

			var where = make(map[string]interface{})
			var d models.MenuApiList
			db := global.Db.Model(&models.MenuApiList{})
			where["url"] = cCp.Request.URL.Path
			tx := db.Preload("Menu").Where(where).First(&d)
			if tx.Error != nil {
				fmt.Println("日志记录错误：", tx.Error.Error())
				log.PathName = "未知请求"
			} else {
				title, ok := d.Menu.Meta["title"]

				if ok {
					log.PathName = title.(string)
				}
			}
			log.Method = cCp.Request.Method
			log.DoData = string(doData)
			log.Ip = cCp.ClientIP()

			log.UrlPath = cCp.Request.URL.Path

			global.Db.Create(&log)

		}()

		c.Next()
	}

}
