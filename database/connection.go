package database

import (
	"go-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:ffGG#2004#@/yt_go_auth"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	// Migrate tables and print logs if migration fails
	err = connection.AutoMigrate(&models.User{}, &models.File{})
	if err != nil {
		panic("Migration failed: " + err.Error()) // Log the migration error
	}
}
