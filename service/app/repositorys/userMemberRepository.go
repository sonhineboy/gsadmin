package repositorys

import (
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"gorm.io/gorm"
)

type UserMemberRepository struct {
	db *gorm.DB
}

// NewUserMemberRepository 实例化
func NewUserMemberRepository() *UserMemberRepository {
	return &UserMemberRepository{
		db: global.Db,
	}
}

//FindById 根据id 查询信息
func (re *UserMemberRepository) FindById(id int) (models.UserMember, error) {
	var (
		model models.UserMember
	)
	tx := re.db.First(&model, id)

	return model, global.GormTans(tx.Error)
}

//UpdateById 根据id 更新信息
func (re *UserMemberRepository) UpdateById(id int, data requests.UserMemberRequest) (int64, error) {
	var (
		model models.UserMember
	)
	tx := re.db.Model(&model).Where("id = ?", id).Updates(models.UserMember{
		
		NickName:	data.NickName,
		
		RealName:	data.RealName,
		
		Age:	data.Age,
		
		Status:	data.Status,
		
		Online:	data.Online,
		
	})
	return tx.RowsAffected, tx.Error
}

//DelByIds 根据id 删除数据
func (re *UserMemberRepository) DelByIds(ids []int) (int64, error) {
	var (
		model models.UserMember
	)
	tx := re.db.Delete(&model, ids)
	return tx.RowsAffected, tx.Error
}

//Page 返回分页数据
func (re *UserMemberRepository) Page(where map[string]interface{}, page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.UserMember
		offSet int
	)
	db := global.Db.Model(&models.UserMember{})

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
	db.Preload("Menus").Limit(pageSize).Order(sortField + " desc" + ",id desc").Offset(offSet)
	db.Find(&data)
	return global.Pages(page, pageSize, int(total), data)
}

//Insert 写入数据
func (re *UserMemberRepository) Insert(data requests.UserMemberRequest) (model models.UserMember, err error) {

	model = models.UserMember{
		NickName:	data.NickName,
		RealName:	data.RealName,
		Age:	data.Age,
		Status:	data.Status,
		Online:	data.Online,
		
	}

	result := re.db.Create(&model)
	err = result.Error
	return model, err
}
