package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	Memory string = "memory"
	SQLite        = "sqlite"
	MySQL         = "mysql"
)

func OpenDB(conf Configure) {
	if db != nil {
		db.Close()
	}
	var err error
	switch conf.Mode {
	case Memory:
		db, err = gorm.Open("sqlite3", "file::memory:")

	case SQLite:
		db, err = gorm.Open("sqlite3", configure.Database.Path)

	default:
		log.Fatalln("Database mode not defined:", conf.Mode)
	}
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Company{}, &Media{}).Error; err != nil {
		panic(err)
	}
}
