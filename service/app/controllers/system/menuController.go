package system

import (
	"ginedu2/service/app/models"
	"ginedu2/service/app/repositorys"
	"ginedu2/service/app/requests"
	"ginedu2/service/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MenuController struct {
	MenuRepository repositorys.SystemMenuRepository
	res            global.Response
}

func (m MenuController) Add(c *gin.Context) {
	var postData requests.MenuPost
	_ = c.ShouldBind(&postData)
	result, model := m.MenuRepository.Add(postData)

	if result.Error == nil {
		m.res.Success(c, "ok", model)
	} else {
		m.res.Failed(c, result.Error.Error())
	}
}

func (m MenuController) Update(c *gin.Context) {
	var postData requests.MenuPost
	_ = c.ShouldBind(&postData)
	result, model := m.MenuRepository.Update(postData)
	if result.Error == nil {
		m.res.Success(c, "ok", model)
	} else {
		m.res.Failed(c, result.Error.Error())
	}
}

func (m MenuController) All(c *gin.Context) {
	m.res.Success(c, "ok", m.MenuRepository.MenuTree())
}

func (m MenuController) Del(c *gin.Context) {

	var delIds requests.MenuDel
	_ = c.ShouldBind(&delIds)
	result := global.Db.Delete(&models.AdminMenu{}, delIds.Ids)

	if result.Error == nil {
		m.res.Success(c, "ok", "")
	} else {
		m.res.Failed(c, result.Error.Error())
	}
}

func (m MenuController) MenuPermissions(c *gin.Context) {
	var (
		myMenus   []map[string]interface{}
		menus     []models.AdminMenu
		adminUser models.AdminUser
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
			myMenus = m.MenuRepository.ArrayToTree(menus, 0)
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
	m.res.Success(c, "ok", data)
}
