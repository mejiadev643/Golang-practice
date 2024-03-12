package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost port=5439 user=mejiadev password=mejiadev dbname=Golang sslmode=disable TimeZone=Asia/Shanghai"
var DB *gorm.DB

func Connect() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal("Error: ", error)
	} else {
		log.Println("Database connected")
	}
}
