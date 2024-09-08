package global

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/global/response"
)

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success disabled
func (r Response) Success(c *gin.Context, msg string, data interface{}) {

	response.Success(c, msg, data)
}

// Failed Deprecated
func (r Response) Failed(c *gin.Context, err string) {
	response.Failed(c, err)
}
