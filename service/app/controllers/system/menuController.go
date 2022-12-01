package system

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/global/response"
	"gorm.io/gorm"
)

type MenuController struct{}

func (m *MenuController) Add(c *gin.Context) {

	var (
		postData       requests.MenuPost
		menuRepository repositorys.SystemMenuRepository
	)
	_ = c.ShouldBind(&postData)
	result, model := menuRepository.Add(postData)

	if result.Error == nil {
		response.Success(c, "ok", model)
	} else {
		response.Failed(c, result.Error.Error())
	}
}

func (m *MenuController) Update(c *gin.Context) {
	var (
		postData       requests.MenuPost
		menuRepository repositorys.SystemMenuRepository
	)
	_ = c.ShouldBind(&postData)
	err, model := menuRepository.Update(postData)
	if err == nil {
		response.Success(c, "ok", model)
	} else {
		response.Failed(c, err.Error())
	}
}

func (m *MenuController) All(c *gin.Context) {
	var (
		menuRepository repositorys.SystemMenuRepository
	)
	response.Success(c, "ok", menuRepository.MenuTree())
}

func (m *MenuController) Del(c *gin.Context) {

	var (
		delIds requests.MenuDel
	)
	_ = c.ShouldBind(&delIds)
	result := global.Db.Delete(&models.AdminMenu{}, delIds.Ids)

	if result.Error == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, result.Error.Error())
	}
}

func (m *MenuController) MenuPermissions(c *gin.Context) {
	var (
		myMenus        []map[string]interface{}
		menus          []models.AdminMenu
		adminUser      models.AdminUser
		menuRepository repositorys.SystemMenuRepository
	)

	v, ok := c.Get("claims")

	if ok {
		claims, err := v.(*models.CustomClaims)
		if err {

			if global.IsSuperAdmin(claims.Roles, global.SuperAdmin) {
				global.Db.Preload("ApiList").Order("sort desc").Find(&menus)
			} else {
				adminUser.ID = claims.Id
				global.Db.Model(&adminUser).Preload("Roles").Preload("Roles.Menus", func(db *gorm.DB) *gorm.DB {
					return db.Order("sort desc")
				}).Preload("Roles.Menus.ApiList").First(&adminUser)
				for _, v := range adminUser.Roles {
					menus = append(menus, v.Menus...)
				}
			}
			myMenus = menuRepository.ArrayToTree(menus, 0)
		}
	}

	var data = make(map[string]interface{})
	data["menu"] = myMenus
	var permissions []string
	for _, v := range menus {
		if len(v.Name) > 0 {
			permissions = append(permissions, v.Name)
		}
	}
	data["permissions"] = permissions
	response.Success(c, "ok", data)
}
