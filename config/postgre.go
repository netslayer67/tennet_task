package config

import (
	"fmt"

	// "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	var DB_HOST = "localhost"
	var DB_USER = "postgres"
	var DB_PASSWORD = "my1fb000"
	var DB_NAME = "task"
	var DB_PORT = "8000"

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	// conn := "root:@tcp(127.0.0.1:3306)/dumbflix?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected")
}
