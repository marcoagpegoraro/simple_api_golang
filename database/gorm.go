package database

import (
	"github.com/marcoagpegoraro/simple_api_golang/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func InitializeDatabase() {
	var err error

	Instance, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	Instance.AutoMigrate(&models.Cliente{})
}
