package database

import (
	"fmt"

	"gorm.io/driver/postgres"  
	"gorm.io/gorm"
)

var DB * gorm.DB

func ConnectDB(){
	dsn := "host=localhost user=postgres password=0406 dbname=online_banking_system port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db

	fmt.Println("Database Connected Successfully!")
}