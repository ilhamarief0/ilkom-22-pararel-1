package models

import (
	"gorm.io/drive/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/product_service"), &gorm.Config{})
	if err != nil{
		panic ("Failed To Connect Database")
	}
	database.AutoMigrate(&Post{})

	DB = database
}