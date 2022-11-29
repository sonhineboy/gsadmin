package repositorys

import (
	"ginedu2/service/app/models"
	"ginedu2/service/app/requests"
	"ginedu2/service/global"
	"gorm.io/gorm"
)

type RoleRepository struct {
	Model models.Role
	Where map[string]interface{}
}

func (r RoleRepository) List(page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.Role
		offSet int
	)
	global.Db.Model(&r.Model).Count(&total)
	offSet = (page - 1) * pageSize
	db := global.Db.Preload("Menus").Limit(pageSize).Order(sortField + " desc" + ",id desc").Offset(offSet)

	if r.Where != nil {
		db.Where(r.Where).Find(&data)
	} else {
		db.Find(&data)

	}

	return global.Pages(page, pageSize, int(total), data)
}

/*
添加角色
*/
func (r RoleRepository) Add(post requests.Role) error {
	db := global.Db.Create(&models.Role{
		Alias:  post.Alias,
		Label:  post.Label,
		Sort:   post.Sort,
		Remark: post.Remark,
		Status: &post.Status,
	})
	return db.Error
}

/*
更新角色
*/
func (r RoleRepository) Update(post requests.Role) error {

	return global.Db.Transaction(func(sessionDb *gorm.DB) error {
		return sessionDb.Debug().Where("id = ?", post.Id).Updates(&models.Role{
			Alias:  post.Alias,
			Label:  post.Label,
			Sort:   post.Sort,
			Remark: post.Remark,
			Status: &post.Status,
		}).Error
	})

}

func (r RoleRepository) UpMenus(post requests.RoleUpMenus) error {
	var role models.Role
	role.ID = post.Id

	if len(post.Menus) > 0 {

		var replace []models.AdminMenu

		for _, v := range post.Menus {
			var li models.AdminMenu
			li.ID = v
			replace = append(replace, li)

		}
		return global.Db.Model(&role).Omit("Menus.*").Association("Menus").Replace(replace)
	} else {
		return global.Db.Model(&role).Association("Menus").Clear()
	}

}
