package middleware

import "github.com/gin-gonic/gin"

func OperationLog() gin.HandlerFunc {

	return func(c *gin.Context) {
		println(c.Request.URL.Path)
		c.Next()
	}

}
