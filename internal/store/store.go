package store

import (
	"blog-service-v3/internal/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	// must be replaced by viper configs
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=blog-service sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&repository.Category{}, &repository.Post{})
	if err != nil {
		panic(err)
	}

	return db
}
