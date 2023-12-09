package initializers

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	// sqlite
	DB_URL := os.Getenv("DB_URL")
	DB, err = gorm.Open(sqlite.Open(DB_URL), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
