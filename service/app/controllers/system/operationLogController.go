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
		params         requests.RoleList
		roleRepository repositorys.RoleRepository
	)
	_ = c.ShouldBind(&params)

	roleRepository.Where = params.Where

	response.Success(c, "ok", roleRepository.List(params.Page, params.PageSize, "sort"))
}
