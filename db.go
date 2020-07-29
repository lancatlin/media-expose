package main

import "github.com/jinzhu/gorm"

func OpenDB(conf Configure) (db *gorm.DB) {
	// for testing
	if conf.Debug {
		db, _ = gorm.Open("sqlite3", "/tmp/gorm.db")
	}
	if err := db.AutoMigrate(&Company{}, &Media{}).Error; err != nil {
		panic(err)
	}
	return
}
