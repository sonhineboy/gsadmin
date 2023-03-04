package repositorys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/event"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	AdminUserModel models.AdminUser
	Where          map[string]interface{}
}

//添加一个用户
func (u *UserRepository) Add(password string, name string, data requests.UserAdd) (*gorm.DB, models.AdminUser) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	u.AdminUserModel.Password = string(pwd)
	u.AdminUserModel.Name = name
	u.AdminUserModel.RealName = data.RealName
	u.AdminUserModel.Avatar = data.Avatar
	for _, v := range data.Roles {
		var role models.Role
		role.ID = v
		u.AdminUserModel.Roles = append(u.AdminUserModel.Roles, role)

	}
	return global.Db.Create(&u.AdminUserModel), u.AdminUserModel
}

//更新用户
func (u *UserRepository) Update(data requests.UserUpdate) error {

	return global.Db.Transaction(func(sessionDb *gorm.DB) error {
		var model models.AdminUser
		model.ID = data.Id
		if len(data.PassWord) > 0 {
			pwd, err := bcrypt.GenerateFromPassword([]byte(data.PassWord), bcrypt.MinCost)
			if err != nil {
				fmt.Println(err)
			}
			u.AdminUserModel.Password = string(pwd)
		}
		u.AdminUserModel.Name = data.Name
		u.AdminUserModel.RealName = data.RealName
		u.AdminUserModel.Avatar = data.Avatar

		db := sessionDb.Where("id = ?", data.Id).Updates(&u.AdminUserModel)
		if db.Error == nil {
			var replace []models.Role
			for _, v := range data.Roles {
				var role models.Role
				role.ID = v
				replace = append(replace, role)
			}
			return sessionDb.Model(&model).Omit("Roles.*").Association("Roles").Replace(replace)

		} else {
			return db.Error
		}
	})

}

//登陆用户
func (u *UserRepository) Login(password string, name string, c *gin.Context) (bool, models.AdminUser) {
	re := global.Db.Where("name = ?", name).Preload("Roles").Preload("Roles.Menus").Preload("Roles.Menus.ApiList").First(&u.AdminUserModel)

	_ = global.GetEventDispatcher(c).Dispatch(event.NewLoginEvent("login", u.AdminUserModel))

	if re.Error == nil && bcrypt.CompareHashAndPassword([]byte(u.AdminUserModel.Password), []byte(password)) == nil {
		return true, u.AdminUserModel
	} else {
		return false, u.AdminUserModel
	}
}

func (u *UserRepository) List(page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.AdminUser
		offSet int
	)
	db := global.Db.Model(&u.AdminUserModel)
	offSet = (page - 1) * pageSize
	global.Db.Debug().Preload("Roles").Limit(pageSize).Order(sortField + " desc" + ",id desc").Offset(offSet)

	if u.Where != nil && len(u.Where) > 0 {
		db.Where("name = ?", u.Where["name"]).Or("real_name = ?", u.Where["name"]).Find(&data)
	} else {
		db.Find(&data)

	}
	db.Count(&total)

	return global.Pages(page, pageSize, int(total), data)
}
