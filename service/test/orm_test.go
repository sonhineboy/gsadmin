package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestTableInfo(t *testing.T) {
	dsn := "root:@tcp(127.0.0.1:3306)/gin_scuiadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	//types, err := db.Migrator().ColumnTypes("member")
	//if err != nil {
	//	return
	//}
	//
	//for _, columnType := range types {
	//
	//	v, _ := columnType.DefaultValue()
	//
	//	fmt.Println("default", v, "len", len(v))
	//}

	var indexes []map[string]interface{}
	db.Raw("show Index from member").Scan(&indexes)

	indexMap := make(map[string]map[string]interface{}, 20)

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

	fmt.Println(indexMap)
}

func GetIndexType(field string) {

}
