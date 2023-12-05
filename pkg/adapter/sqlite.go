package adapter

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mygin/pkg/cfg"
	"mygin/pkg/model"
)

func ConnectSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cfg.GetConfig().Db.Sqlite.Filename), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Message{})
	return db
}