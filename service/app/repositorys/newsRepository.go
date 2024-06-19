package repositorys

import (
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"gorm.io/gorm"
)

type NewsRepository struct {
	db *gorm.DB
}

// NewNewsRepository 实例化
func NewNewsRepository() *NewsRepository {
	return &NewsRepository{
		db: global.Db,
	}
}

//FindById 根据id 查询信息
func (re *NewsRepository) FindById(id int) (models.News, error) {
	var (
		model models.News
	)
	tx := re.db.First(&model, id)

	return model, global.GormTans(tx.Error)
}

//UpdateById 根据id 更新信息
func (re *NewsRepository) UpdateById(id int, data requests.NewsRequest) (int64, error) {
	var (
		model models.News
	)
	tx := re.db.Model(&model).Where("id = ?", id).Updates(models.News{
		
		Title:	data.Title,
		
		Author:	data.Author,
		
		Content:	data.Content,
		
		Image:	data.Image,
		
	})
	return tx.RowsAffected, tx.Error
}

//DelByIds 根据id 删除数据
func (re *NewsRepository) DelByIds(ids []int) (int64, error) {
	var (
		model models.News
	)
	tx := re.db.Delete(&model, ids)
	return tx.RowsAffected, tx.Error
}

//Page 返回分页数据
func (re *NewsRepository) Page(where map[string]interface{}, page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.News
		offSet int
	)
	db := global.Db.Model(&models.News{})

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
func (re *NewsRepository) Insert(data requests.NewsRequest) (model models.News, err error) {

	model = models.News{
		Title:	data.Title,
		Author:	data.Author,
		Content:	data.Content,
		Image:	data.Image,
		
	}

	result := re.db.Create(&model)
	err = result.Error
	return model, err
}
