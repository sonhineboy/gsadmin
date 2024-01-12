package repositorys

import (
	"github.com/sonhineboy/gsadmin/service/global"
	"gorm.io/gorm"
)

type Repository interface {
	//获取数据库连接
	GetDb() *gorm.DB
	SetDb()
	//分页
	Page(page int, pageSize int, sortField string, data interface{}) map[string]interface{}
	//模型
	GetModel() interface{}
	SetModel()
	//添加
	Add(data interface{}) error
	//按id更新
	Update(id uint, data interface{}) error
	//按条件更新
	UpdateByWhere(where interface{}, data interface{}) error
	//删除数据
	Delete(condos ...interface{}) error
}

type BaseRepository struct {
	Db    *gorm.DB
	Model interface{}
	Where map[string]interface{}
}

func (r *BaseRepository) SetDb() {
	r.Db = global.Db
}

func (r *BaseRepository) GetDb() *gorm.DB {

	if r.Db == nil {
		r.SetDb()
	}
	return r.Db
}

func (r *BaseRepository) Add(data interface{}) error {
	return r.GetDb().Create(&data).Error
}

func (r *BaseRepository) Delete(condos ...interface{}) error {
	return r.GetDb().Delete(&r.Model, condos...).Error
}

func (r *BaseRepository) Update(id uint, data interface{}) error {
	return r.GetDb().Model(&r.Model).Where("id = ?", id).Updates(data).Error
}

func (r *BaseRepository) Page(page int, pageSize int, sortField string, data interface{}) map[string]interface{} {
	var (
		total  int64
		offSet int
	)
	db := r.GetDb().Model(&r.Model)
	if r.Where != nil && len(r.Where) > 0 {
		db.Where(r.Where)
	}
	db.Count(&total)
	offSet = (page - 1) * pageSize
	db.Preload("Menus").Limit(pageSize).Order(sortField + " desc" + ",id desc").Offset(offSet)
	db.Find(&data)
	return global.Pages(page, pageSize, int(total), data)
}

//按条件更新
func (r *BaseRepository) UpdateByWhere(where interface{}, data interface{}) error {
	return nil
}

func (r *TestRepository) GetModel() interface{} {
	if r.Model == nil {
		r.SetModel()
	}
	return r.Model
}
