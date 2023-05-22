package repositorys

import (
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/global"
)

type OperationLogRepository struct {
	Model models.OperationLog
	Where map[string]interface{}
}

func (o *OperationLogRepository) List(page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.OperationLog
		offSet int
	)
	db := global.Db.Model(&o.Model)
	db.Count(&total)
	offSet = (page - 1) * pageSize
	db.Preload("Menus").Limit(pageSize).Order(sortField + " desc" + ",id desc").Offset(offSet)

	if o.Where != nil && len(o.Where) > 0 {
		db.Where(o.Where).Find(&data)
	} else {
		db.Find(&data)
	}
	return global.Pages(page, pageSize, int(total), data)
}
