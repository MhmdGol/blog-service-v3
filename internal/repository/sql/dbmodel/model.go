package dbmodel

import "gorm.io/gorm"

func Init(db *gorm.DB) error {
	return db.AutoMigrate(&Category{}, &Post{})
}
