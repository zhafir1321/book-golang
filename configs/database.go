package configs

import (
	"book-golang/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/?parseTime=true"), &gorm.Config{})

	db.Exec("CREATE DATABASE IF NOT EXISTS `go-book`")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/go-book?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.User{}, &models.Book{}, &models.BorrowReturn{})
	
	DB = db

	log.Println("Database connection successfully!")
}