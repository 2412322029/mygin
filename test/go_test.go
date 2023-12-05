package test

import (
	"fmt"
	"gorm.io/gorm"
	"mygin/pkg/adapter"
	"mygin/pkg/cfg"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Printf("%v", cfg.GetConfig().Db)
	var db *gorm.DB
	if cfg.GetConfig().Db.Dbtype == "sqlite" {
		db = adapter.ConnectSqlite()
	} else if cfg.GetConfig().Db.Dbtype == "mysql" {
		db = adapter.ConnectMysql()
	}
	if db == nil {
		panic("unknown dbtype")
	}
	sqlDB, _ := db.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数

	println(db)
}
