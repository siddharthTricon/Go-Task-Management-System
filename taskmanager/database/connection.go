package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connect(){
	dsn := "root:SIddharth@03@/taskmanager?charset=utf8&parseTime=True&loc=Local" 
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Database connected successfully")

}


