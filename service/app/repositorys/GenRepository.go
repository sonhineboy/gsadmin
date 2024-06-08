package repositorys

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadminGen"
	"github.com/sonhineboy/gsadminGen/pkg"
	"gorm.io/gorm"
)

type GenRepository struct {
	Db *gorm.DB
}

func NewGenRepository() *GenRepository {
	return &GenRepository{
		Db: global.Db,
	}
}

func (r *GenRepository) GetTables() (tableSlice []map[string]string, err error) {
	tableSlice = make([]map[string]string, 0, 10)
	tables, err := r.Db.Migrator().GetTables()
	if err != nil {
		return nil, err
	}
	for _, v := range tables {
		tableSlice = append(tableSlice, map[string]string{"label": v, "value": v})
	}
	return
}

func (r *GenRepository) TableField(name string, function func(fieldsSlices []map[string]interface{}, columnType gorm.ColumnType, r *GenRepository) []map[string]interface{}) ([]map[string]interface{}, error) {

	column, err := r.Db.Migrator().ColumnTypes(name)

	if err != nil {
		return nil, err
	}

	ignoreField := r.getIgnoreField()

	fieldsSlices := make([]map[string]interface{}, 0, 20)

	for _, v := range column {
		_, ok := ignoreField[v.Name()]
		if ok {
			continue
		}
		fieldsSlices = function(fieldsSlices, v, r)
	}

	return fieldsSlices, nil
}

func (r *GenRepository) GetIndexType(column string, Indexes map[string]map[string]interface{}) string {

	v, ok := Indexes[column]
	if ok {
		if v["Non_unique"] == 0 {
			return "UNIQUE"
		}
		if v["Index_type"] == "FULLTEXT" {
			return "FULLTEXT"
		}
		return "NORMAL"
	}
	return "Null"
}

func (r *GenRepository) GetTablesIndexes(tables string) (indexMap map[string]map[string]interface{}) {
	var indexes []map[string]interface{}
	r.Db.Raw(fmt.Sprint("show Index from ", tables)).Scan(&indexes)
	indexMap = make(map[string]map[string]interface{}, 20)
	for _, v := range indexes {
		columnName, _ := v["Column_name"]
		if columnNameString, ok := columnName.(string); ok {
			indexMap[columnNameString] = map[string]interface{}{
				"Non_unique": v["Non_unique"],
				"Index_type": v["Index_type"],
			}
		} else {
			continue
		}
	}
	return
}

func (r *GenRepository) getIgnoreField() map[string]string {
	return map[string]string{
		"id":         "true",
		"created_at": "true",
		"updated_at": "true",
		"deleted_at": "true",
	}
}

func (r *GenRepository) GenCode(data requests.GenCode) error {

	v := pkg.TableModal{
		Name:   data.TableDiyName,
		Fields: data.Fields,
	}
	var err error
	if err = gsadminGen.GenController("./app/controllers/"+data.ControllerPackage+"/"+gsadminGen.UnderToConvertSoreLow(v.Name)+"Controller.go", v, data.ControllerPackage); err != nil {
		return err
	}
	if err = gsadminGen.GenModel("./app/models/"+gsadminGen.UnderToConvertSoreLow(v.Name)+".go", v); err != nil {
		return err
	}
	if err = gsadminGen.GenRequest("./app/requests/"+gsadminGen.UnderToConvertSoreLow(v.Name)+"Request.go", v); err != nil {
		return err
	}
	err = gsadminGen.GenRepository("./app/repositorys/"+gsadminGen.UnderToConvertSoreLow(v.Name)+"Repository.go", v)
	if err != nil {
		return err
	}
	return nil
}
