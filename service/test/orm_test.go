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

	db.Migrator()

	tables, err := db.Migrator().GetTables()
	if err != nil {
		return
	}

	for _, v := range tables {

		fmt.Println(v)
	}
	types, err := db.Migrator().ColumnTypes("member")
	if err != nil {
		return
	}

	for _, v := range types {

		null, _ := v.Nullable()
		fmt.Println(v.Name(), v.DatabaseTypeName(), null)
	}
}
