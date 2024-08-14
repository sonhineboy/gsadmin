package initialize

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func DbInit(c *config.Config) *gorm.DB {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("数据库连接失败", err)
		}
	}()

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
