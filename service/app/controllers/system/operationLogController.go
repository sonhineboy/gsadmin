package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/global/response"
)

type OperationLogController struct{}

func (o *OperationLogController) List(c *gin.Context) {

	var (
		params       global.List
		operationLog repositorys.OperationLogRepository
	)
	_ = c.ShouldBind(&params)
	value, ok := params.Where["created_at"].(map[string]interface{})

	fmt.Print(111111111, ok)

	if ok {
		fmt.Print(value["end"])
	}
	operationLog.Where = params.Where

	response.Success(c, "ok", operationLog.List(params.Page, params.PageSize, "created_at"))
}
