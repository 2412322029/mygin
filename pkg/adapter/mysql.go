package adapter

import (
	"fmt"
	"mygin/pkg/cfg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql() *gorm.DB {
	username := cfg.GetConfig().Db.Mysql.Username //账号
	password := cfg.GetConfig().Db.Mysql.Password //密码
	host := cfg.GetConfig().Db.Mysql.Host         //数据库地址
	port := cfg.GetConfig().Db.Mysql.Port         //数据库端口
	Dbname := cfg.GetConfig().Db.Mysql.Dbname     //数据库名
	timeout := "10s"                              //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
