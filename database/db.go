package database

import (
	"log"

	"github.com/prranavv/Backend_Project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBinstance struct {
	Db *gorm.DB
}

var Database DBinstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connected into the database successfully")
	//Add Migration
	db.AutoMigrate(&models.Task{})
	Database = DBinstance{Db: db}
}
