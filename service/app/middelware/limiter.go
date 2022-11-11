package middelware

import (
	"ginedu2/service/global"
	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.Limiter.Allow() == false {
			global.Response{}.Failed(c, "当前请求过快，请稍后再试！")
			c.Abort()
			return
		}
		c.Next()
	}
}
