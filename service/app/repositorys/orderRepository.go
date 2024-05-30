package repositorys

import (
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository 实例化
func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		db: global.Db,
	}
}

//FindById 根据id 查询信息
func (re *OrderRepository) FindById(id int) (models.Order, error) {
	var order models.Order
	tx := re.db.First(&order, id)
	return order, tx.Error
}

//UpdateById 根据id 更新信息
func (re *OrderRepository) UpdateById(id int, data requests.OrderRequest) (int64, error) {
	var (
		model models.Order
	)
	tx := re.db.Model(&model).Where("id = ?", id).Updates(models.Order{
		Age:      data.Age,
		UserName: data.UserName,
	})
	return tx.RowsAffected, tx.Error
}

//DelByIds 根据id 删除数据
func (re *OrderRepository) DelByIds(ids []int) (int64, error) {
	var (
		model models.Order
	)
	tx := re.db.Delete(&model, ids)
	return tx.RowsAffected, tx.Error
}

//Page 返回分页数据
func (re *OrderRepository) Page(where map[string]interface{}, page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.Order
		offSet int
	)
	db := global.Db.Model(&models.Order{})

	if where != nil && len(where) > 0 {
		db.Where(where)
	}
	db.Count(&total)

	if page <= 0 {
		page = 1
	}
	offSet = (page - 1) * pageSize
	db.Preload("Menus").Limit(pageSize).Order(sortField + " desc" + ",id desc").Offset(offSet)
	db.Find(&data)
	return global.Pages(page, pageSize, int(total), data)
}

//Insert 写入数据
func (re *OrderRepository) Insert(data requests.OrderRequest) (model models.Order, err error) {

	model = models.Order{
		Age:      data.Age,
		UserName: data.UserName,
	}

	result := re.db.Create(&model)
	err = result.Error
	return model, err
}
