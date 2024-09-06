package repositorys

import (
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"gorm.io/gorm"
)

type Demo2Repository struct {
	db *gorm.DB
}

// NewDemo2Repository 实例化
func NewDemo2Repository() *Demo2Repository {
	return &Demo2Repository{
		db: global.Db,
	}
}

//FindById 根据id 查询信息
func (re *Demo2Repository) FindById(id int) (models.Demo2, error) {
	var (
		model models.Demo2
	)
	tx := re.db.First(&model, id)

	return model, global.GormTans(tx.Error)
}

//UpdateById 根据id 更新信息
func (re *Demo2Repository) UpdateById(id int, data requests.Demo2Request) (int64, error) {
	var (
		model models.Demo2
	)
	tx := re.db.Model(&model).Where("id = ?", id).Updates(models.Demo2{
		
		Name:	data.Name,
		
		Age:	data.Age,
		
	})
	return tx.RowsAffected, tx.Error
}

//DelByIds 根据id 删除数据
func (re *Demo2Repository) DelByIds(ids []int) (int64, error) {
	var (
		model models.Demo2
	)
	tx := re.db.Delete(&model, ids)
	return tx.RowsAffected, tx.Error
}

//Page 返回分页数据
func (re *Demo2Repository) Page(where map[string]interface{}, page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.Demo2
		offSet int
	)
	db := global.Db.Model(&models.Demo2{})

	for k, v := range where {
		if len(v.(string)) == 0 {
			delete(where, k)
		}
	}

	if where != nil && len(where) > 0 {
		db.Where(where)
	}
	db.Count(&total)

	if page <= 0 {
		page = 1
	}
	offSet = (page - 1) * pageSize
	db.Limit(pageSize).Order(sortField + " desc" + ",id desc").Offset(offSet)
	db.Find(&data)
	return global.Pages(page, pageSize, int(total), data)
}

//Insert 写入数据
func (re *Demo2Repository) Insert(data requests.Demo2Request) (model models.Demo2, err error) {

	model = models.Demo2{
		Name:	data.Name,
		Age:	data.Age,
		
	}

	result := re.db.Create(&model)
	err = result.Error
	return model, err
}
