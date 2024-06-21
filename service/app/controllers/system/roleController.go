package system

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/global/response"
)

type RoleController struct {
}

func (r *RoleController) List(c *gin.Context) {
	var (
		params         requests.RoleList
		roleRepository repositorys.RoleRepository
	)
	_ = c.ShouldBind(&params)

	roleRepository.Where = params.Where

	response.Success(c, "ok", roleRepository.List(params.Page, params.PageSize, "sort"))
}

func (r *RoleController) Up(c *gin.Context) {
	var (
		post           requests.Role
		roleRepository repositorys.RoleRepository
	)

	err := c.ShouldBind(&post)

	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	if err = roleRepository.Update(post); err == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, err.Error())
	}
}

func (r *RoleController) Add(c *gin.Context) {
	var (
		post           requests.Role
		roleRepository repositorys.RoleRepository
	)

	err := c.ShouldBind(&post)

	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	if err = roleRepository.Add(post); err == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, err.Error())
	}
}

func (r *RoleController) Del(c *gin.Context) {
	var delIds requests.RoleDel
	_ = c.ShouldBind(&delIds)
	result := global.Db.Delete(&models.Role{}, delIds.Ids)
	if result.Error == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, result.Error.Error())
	}
}

func (r *RoleController) RoleUpMenu(c *gin.Context) {

	var (
		data           requests.RoleUpMenus
		roleRepository repositorys.RoleRepository
	)
	_ = c.ShouldBind(&data)
	result := roleRepository.UpMenus(data)
	if result == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, result.Error())
	}

}
