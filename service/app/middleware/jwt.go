package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/global"
	"log"
	"net/http"
	"strings"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无权限访问，请求未携带token",
			})
			ctx.Abort() //结束后续操作
			return
		}
		log.Print("token:", authHeader)

		//按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			global.Response{}.Failed(ctx, "请求头中auth格式有误")
			ctx.Abort()
			return
		}

		claims, err := models.ParseToken(parts[1], global.Config.MyJwt.Secret)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "无权限访问，token无效" + err.Error(),
			})
			ctx.Abort()
			return
		}

		// 将当前请求的claims信息保存到请求的上下文c上
		ctx.Set("claims", claims)
		ctx.Set("permission", repositorys.NewPermissionRepository(claims))
		ctx.Next() // 后续的处理函数可以用过ctx.Get("claims")来获取当前请求的用户信息

	}
}
