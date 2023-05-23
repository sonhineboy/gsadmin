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

	if o.Where != nil && len(o.Where) > 0 {
		createdAt, ok := o.Where["created_at"]
		if ok {
			delete(o.Where, "created_at")
			createdAtMap, ok := createdAt.(map[string]interface{})
			if ok {
				start, startOk := createdAtMap["begin"]
				end, endOk := createdAtMap["end"]
				if startOk && endOk {
					db.Where("created_at BETWEEN ? and ?", start, end)
				}
			}

		}

		db.Where(o.Where)
	}
	db.Count(&total)

	offSet = (page - 1) * pageSize
	db.Preload("Menus").Limit(pageSize).Order(sortField + " desc" + ",id desc").Offset(offSet)
	db.Find(&data)
	return global.Pages(page, pageSize, int(total), data)
}
