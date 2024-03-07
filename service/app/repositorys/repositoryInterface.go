package repositorys

import (
	"github.com/sonhineboy/gsadmin/service/global"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	// GetDb 获取数据库连接
	GetDb() *gorm.DB
	SetDb()
	// Page 分页
	Page(page int, pageSize int, sortField string, data interface{}) map[string]interface{}
	// GetModel 模型
	GetModel() interface{}
	SetModel()
	// Add 添加
	Add(data interface{}) error
	// UpdateById 按id更新
	UpdateById(id uint, data interface{}) *gorm.DB
	// UpdateByWhere 按条件更新
	UpdateByWhere(data interface{}, query interface{}, args ...interface{}) *gorm.DB
	// Delete 删除数据
	Delete(condos ...interface{}) error
}

type BaseRepository struct {
	Db      *gorm.DB
	Model   interface{}
	Where   map[string]interface{}
	Preload []string
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
	return r.GetDb().Create(data).Error
}

func (r *BaseRepository) Delete(condos ...interface{}) error {
	return r.GetDb().Delete(&r.Model, condos...).Error
}

func (r *BaseRepository) UpdateById(id uint, data interface{}) *gorm.DB {
	return r.GetDb().Model(&r.Model).Where("id = ?", id).Updates(data)
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
	if r.Preload != nil && len(r.Preload) > 0 {
		for _, v := range r.Preload {
			db.Preload(v)
		}
	}
	db.Limit(pageSize).Order(sortField + " desc" + ",id desc").Offset(offSet)
	db.Find(data)
	return global.Pages(page, pageSize, int(total), data)
}

// UpdateByWhere 按条件更新
func (r *BaseRepository) UpdateByWhere(data interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return r.GetDb().Model(&r.Model).Where(query, args).Updates(data)
}

func (r *BaseRepository) GetModel() interface{} {
	return r.Model
}
