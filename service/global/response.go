package global

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

//disabled
func (r Response) Success(c *gin.Context, msg string, data interface{}) {
	r.Code = 200
	r.Msg = msg
	r.Message = msg
	r.Data = data
	c.JSON(http.StatusOK, r)
}

//disabled
func (r Response) Failed(c *gin.Context, err string) {
	r.Code = 422
	r.Msg = err
	r.Message = err
	r.Data = []string{}
	c.JSON(http.StatusOK, r)
}
