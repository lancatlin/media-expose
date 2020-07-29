package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func OpenDB(conf Configure) (db *gorm.DB) {
	// for testing
	if conf.Debug {
		var err error
		db, err = gorm.Open("sqlite3", "/tmp/gorm.db")
		if err != nil {
			panic(err)
		}
	}
	if err := db.AutoMigrate(&Company{}, &Media{}).Error; err != nil {
		panic(err)
	}
	return
}
