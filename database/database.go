package database

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB * gorm.DB

func ConnectDB(){
	dsn :=  os.Getenv("databseUrl")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db

	fmt.Println("Database Connected Successfully!")
}