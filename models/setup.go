package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")
	fmt.Println("Db", database)
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})
	DB = database
}
