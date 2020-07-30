package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func OpenDB(conf Configure) {
	var err error
	if conf.Debug {
		// for testing
		db, err = gorm.Open("sqlite3", "file::memory:?cache=shared")
		if err != nil {
			panic(err)
		}
	}
	if err := db.AutoMigrate(&Company{}, &Media{}).Error; err != nil {
		panic(err)
	}
}
