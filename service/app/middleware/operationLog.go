package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/global"
)

func OperationLog() gin.HandlerFunc {

	return func(c *gin.Context) {

		var (
			doData []byte
			err    error
		)

		contentType := c.GetHeader("Content-Type")

		switch c.Request.Method {
		case "POST":
			if strings.Contains(contentType, "multipart/form-data") {
				doData = []byte("图片上传")
			} else {
				doData, err = io.ReadAll(c.Request.Body)
				if err != nil {
					fmt.Println("中间件错误--POST数据》", err)
					return
				}
				c.Request.Body = io.NopCloser(bytes.NewBuffer(doData))
			}
		case "GET":
			doData, err = json.Marshal(c.Request.URL.Query())
			if err != nil {
				fmt.Println("中间件错误--GET数据》", err)
				return
			}
		}
		claims, ok := repositorys.GetCustomClaims(c)
		if !ok {
			return
		}
		go func(doData []byte, claims *models.CustomClaims, ip string, method string, path string) {
			var (
				log models.OperationLog
			)

			log.UserId = claims.Id
			log.UserName = claims.Name

			var where = make(map[string]interface{})
			var d models.MenuApiList
			db := global.Db.Model(&models.MenuApiList{})
			where["url"] = path
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
			log.Method = method
			log.DoData = string(doData)
			log.Ip = ip
			log.UrlPath = path
			global.Db.Create(&log)

		}(doData, claims, c.ClientIP(), c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}
