package controllers

import (
	"ginedu2/service/app/models"
	"ginedu2/service/app/repositorys"
	"ginedu2/service/app/requests"
	"ginedu2/service/global"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	repository repositorys.RoleRepository
	res        global.Response
}

func (r RoleController) List(c *gin.Context) {

	var params requests.RoleList
	_ = c.ShouldBind(&params)

	if params.Where != nil {
		r.repository.Where = params.Where
	}

	r.res.Success(c, "ok", r.repository.List(params.Page, params.PageSize, "sort"))
}

func (r RoleController) Up(c *gin.Context) {
	var post requests.Role

	err := c.ShouldBind(&post)

	if err != nil {
		r.res.Failed(c, err.Error())
		return
	}
	if err = r.repository.Update(post); err == nil {
		r.res.Success(c, "ok", "")
	} else {
		r.res.Failed(c, err.Error())
	}
}

func (r RoleController) Add(c *gin.Context) {
	var post requests.Role

	err := c.ShouldBind(&post)

	if err != nil {
		r.res.Failed(c, err.Error())
		return
	}
	if err = r.repository.Add(post); err == nil {
		r.res.Success(c, "ok", "")
	} else {
		r.res.Failed(c, err.Error())
	}
}

func (r RoleController) Del(c *gin.Context) {
	var delIds requests.RoleDel
	_ = c.ShouldBind(&delIds)
	result := global.Db.Delete(&models.Role{}, delIds.Ids)
	if result.Error == nil {
		r.res.Success(c, "ok", "")
	} else {
		r.res.Failed(c, result.Error.Error())
	}
}

func (r RoleController) RoleUpMenu(c *gin.Context) {

	var data requests.RoleUpMenus
	_ = c.ShouldBind(&data)
	result := r.repository.UpMenus(data)
	if result == nil {
		r.res.Success(c, "ok", "")
	} else {
		r.res.Failed(c, result.Error())
	}

}
