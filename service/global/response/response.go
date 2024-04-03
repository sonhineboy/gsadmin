package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, msg string, data interface{}) {
	r := Response{
		Code:    200,
		Msg:     msg,
		Message: msg,
		Data:    data,
	}
	c.JSON(http.StatusOK, r)
}

func Failed(c *gin.Context, err string) {
	r := Response{
		Code:    422,
		Msg:     err,
		Message: err,
		Data:    []string{},
	}
	c.JSON(http.StatusOK, r)
}
