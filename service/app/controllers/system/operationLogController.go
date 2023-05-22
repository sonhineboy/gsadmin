package system

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global/response"
)

type OperationLogController struct{}

func (o *OperationLogController) List(c *gin.Context) {

	var (
		params       requests.RoleList
		operationLog repositorys.OperationLogRepository
	)
	_ = c.ShouldBind(&params)

	operationLog.Where = params.Where

	response.Success(c, "ok", operationLog.List(params.Page, params.PageSize, "created_at"))
}
