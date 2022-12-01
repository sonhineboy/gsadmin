package router

import "github.com/sonhineboy/gsadmin/service/app/controllers/system"

type Controllers struct {
	system.UserController
	system.CommonController
	system.MenuController
	system.RoleController
}

var ApiControllers = new(Controllers)
