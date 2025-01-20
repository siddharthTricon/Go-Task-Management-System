package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	dsn := "root:Siddharth@03@/taskmanager?charset=utf8&parseTime=True&loc=Local" 
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Failed to connect to the database: ", err)
	}

	DB = db

	log.Println("Database connected successfully")

}


