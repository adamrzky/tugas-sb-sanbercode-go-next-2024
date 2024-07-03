package config

import (
	"final-project-golang-individu/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/culinary?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.User{}, &models.Profile{}, &models.Restaurant{}, &models.Food{}, &models.Review{}, &models.Comment{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
}
