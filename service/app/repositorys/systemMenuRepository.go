package repositorys

import (
	"errors"
	"ginedu2/service/app/models"
	"ginedu2/service/app/requests"
	"ginedu2/service/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type SystemMenuRepository struct {
	MenuModel models.AdminMenu
}

//添加菜单
func (menu SystemMenuRepository) Add(post requests.MenuPost) (*gorm.DB, models.AdminMenu) {
	menu.MenuModel.Name = post.Name
	menu.MenuModel.Component = post.Component
	menu.MenuModel.Meta = post.Meta
	menu.MenuModel.ParentId = post.ParentId
	menu.MenuModel.Path = post.Path
	menu.MenuModel.Redirect = post.Redirect
	for _, v := range post.ApiList {
		menu.MenuModel.ApiList = append(menu.MenuModel.ApiList, models.MenuApiList{
			Code: v["code"],
			Url:  v["url"],
		})
	}

	return global.Db.Create(&menu.MenuModel), menu.MenuModel
}

//更新菜单
func (menu SystemMenuRepository) Update(post requests.MenuPost) (error, models.AdminMenu) {
	var updateData models.AdminMenu
	updateData.Name = post.Name
	updateData.Component = post.Component
	updateData.Meta = post.Meta
	updateData.ParentId = post.ParentId
	updateData.Path = post.Path
	updateData.Redirect = post.Redirect
	updateData.Sort = post.Sort
	id, _ := strconv.Atoi(post.Id)
	menu.MenuModel.ID = uint(id)

	return global.Db.Transaction(func(sessionDb *gorm.DB) error {
		var notDelIds []uint
		for _, v := range post.ApiList {
			var apiList models.MenuApiList
			if len(v["id"]) > 0 {
				//更新业务逻辑
				id, _ := strconv.Atoi(v["id"])
				apiList.ID = uint(id)
				delete(v, "id")
				upDb := sessionDb.Model(&apiList).Updates(models.MenuApiList{
					Code: v["code"],
					Url:  v["url"],
				})

				if upDb.Error != nil {
					return upDb.Error
				}
				notDelIds = append(notDelIds, apiList.ID)
			} else {
				//新增业务逻辑
				var addModel = models.MenuApiList{
					MenuId: menu.MenuModel.ID,
					Code:   v["code"],
					Url:    v["url"],
				}
				if adDb := sessionDb.Create(&addModel); adDb.Error != nil {
					return adDb.Error
				}
				notDelIds = append(notDelIds, addModel.ID)

			}
		}

		//同步，自动删除不存在的id
		if len(notDelIds) > 0 {
			if syncDb := sessionDb.Not(notDelIds).Where("menu_id = ?", id).Delete(&models.MenuApiList{}); syncDb.Error != nil {
				return syncDb.Error
			}
		}
		return sessionDb.Model(&menu.MenuModel).Updates(updateData).Error
	}), menu.MenuModel
}

func (menu SystemMenuRepository) ArrayToTree(arr []models.AdminMenu, pid uint) []map[string]interface{} {
	var newArr []map[string]interface{}
	for _, v := range arr {
		mapData := make(map[string]interface{})
		if v.ParentId == pid {
			children := menu.ArrayToTree(arr, v.ID)

			if children == nil {
				mapData["children"] = make([]interface{}, 0)
			} else {
				mapData["children"] = children
			}
			mapData["id"] = v.ID
			mapData["meta"] = v.Meta
			mapData["component"] = v.Component
			mapData["name"] = v.Name
			mapData["parent_id"] = v.ParentId
			mapData["path"] = v.Path
			mapData["sort"] = v.Sort
			mapData["redirect"] = v.Redirect
			mapData["apiList"] = v.ApiList
			newArr = append(newArr, mapData)
		}
	}
	return newArr
}

func (menu SystemMenuRepository) MenuTree() []map[string]interface{} {
	var all []models.AdminMenu
	global.Db.Preload("ApiList").Order("sort desc").Find(&all)
	return menu.ArrayToTree(all, 0)
}

func (menu SystemMenuRepository) GetCustomClaims(c *gin.Context) (*models.CustomClaims, bool) {
	v, ok := c.Get("claims")
	claims, err := v.(*models.CustomClaims)
	if ok && err {
		return claims, true
	} else {
		return &models.CustomClaims{}, false
	}
}

//根据当前登陆得用户获得api 权限
func (menu SystemMenuRepository) GetApiList(c *gin.Context, apiList *[]models.MenuApiList) error {
	claims, ok := menu.GetCustomClaims(c)
	if ok {
		var adminUser models.AdminUser
		adminUser.ID = claims.Id
		global.Db.Model(&adminUser).Preload("Roles.Menus.ApiList").First(&adminUser)
		for _, role := range adminUser.Roles {
			for _, menu := range role.Menus {
				*apiList = append(*apiList, menu.ApiList...)
			}
		}
		return nil
	} else {
		return errors.New("无法处理")
	}
}

//根据传递用户对象
func (menu SystemMenuRepository) GetApiListByUser(adminUser models.AdminUser, apiList *[]models.MenuApiList) error {
	for _, role := range adminUser.Roles {
		for _, menu := range role.Menus {
			*apiList = append(*apiList, menu.ApiList...)
		}
	}
	return nil
}

//根据传递用户对象
func (menu SystemMenuRepository) GetPermissionByUser(adminUser models.AdminUser, permission *[]string) error {
	for _, role := range adminUser.Roles {
		for _, menu := range role.Menus {
			*permission = append(*permission, menu.Name)
		}
	}
	return nil
}

//获取map apiList
func (menu SystemMenuRepository) GetApiListToMap(c *gin.Context, apiListMap *map[string]string) error {
	var apiList []models.MenuApiList
	err := menu.GetApiList(c, &apiList)
	if err != nil {
		return err
	} else {
		for _, v := range apiList {
			(*apiListMap)[v.Url] = v.Code
		}
		return nil
	}
}

//获取map apiList 根据用户
func (menu SystemMenuRepository) GetApiListToMapByUser(adminUser models.AdminUser, apiListMap *map[string]string) error {
	var apiList []models.MenuApiList
	_ = menu.GetApiListByUser(adminUser, &apiList)
	for _, v := range apiList {
		(*apiListMap)[v.Url] = v.Code
	}
	return nil
}
