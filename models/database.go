package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {

	database, err := gorm.Open(sqlite.Open("budgeter.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&User{}, &Expense{}, &Budget{})

	DB = database
}
