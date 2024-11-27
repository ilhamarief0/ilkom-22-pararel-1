package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:111999@tcp(127.0.0.1:3306)/ecommerce"), &gorm.Config{})
	if err != nil {
		panic("Failed To Connect Database")
	}
	database.AutoMigrate(&Product{})

	DB = database
}
