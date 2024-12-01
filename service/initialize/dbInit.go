package initialize

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/config"
	"github.com/sonhineboy/gsadmin/service/global"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func DbInit(c *config.Config) *gorm.DB {

	switch c.Db.Type {
	case "mysql":
		return initMysql(c)
	case "sqlite":
		return initSqlite(c)
	default:
		panic(fmt.Sprintf("not support type: %s", c.Db.Type))
	}

}

func initMysql(c *config.Config) *gorm.DB {
	dsn := fmt.Sprint(c.Db.User, ":", c.Db.PassWord, "@tcp(", c.Db.Host, ":", c.Db.Port, ")/", c.Db.Database, "?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Db.TablePrefix,
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err.Error())
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(c.Db.MaxOpenConns)
	sqlDb.SetMaxIdleConns(c.Db.MaxIdleConns)
	sqlDb.SetConnMaxIdleTime(time.Hour)
	return db
}

func initSqlite(c *config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(c.Db.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Db.TablePrefix,
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err.Error())
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(c.Db.MaxOpenConns)
	sqlDb.SetMaxIdleConns(c.Db.MaxIdleConns)
	sqlDb.SetConnMaxIdleTime(time.Hour)
	return db
}

func AutoMigrate(db *gorm.DB) {

	if global.Config.Db.Type == "mysql" {
		db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	err := db.AutoMigrate(
		//slot start not delete
		&models.AdminUser{},
		&models.AdminMenu{},
		&models.MenuApiList{},
		&models.Role{},
		&models.OperationLog{},
		&models.News{},
		//slot end not delete
	)
	if err != nil {
		panic(err)
	}
}

func DbClose(db *gorm.DB) func() {
	return func() {
		db, _ := db.DB()
		_ = db.Close()
	}
}
