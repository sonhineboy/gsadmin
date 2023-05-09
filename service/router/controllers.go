package router

import (
	"github.com/sonhineboy/gsadmin/service/app/controllers/system"
)

type Controllers struct {
	system.UserController
	system.CommonController
	system.MenuController
	system.RoleController
	system.OperationLogController
}

var ApiControllers = new(Controllers)
