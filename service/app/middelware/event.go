package middelware

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/global"
)

func Event() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("e", global.EventDispatcher)
		c.Next()
	}
}
