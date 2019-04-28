package models

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "data/storage.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}