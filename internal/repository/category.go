package repository

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string  `json:"name"`
	Posts []*Post `gorm:"many2many:post_categories;"`
}

func (s *categoryRepository) CreateCategory(name string) error {
	err := s.db.Create(&Category{
		Name: name,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *categoryRepository) AllCategories() ([]Category, error) {
	var categories []Category

	err := s.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *categoryRepository) UpdateCategory(categoryID int, name string) error {
	var categoryToUpdate Category
	err := s.db.First(&categoryToUpdate, categoryID).Error
	if err != nil {
		return err
	}

	categoryToUpdate.Name = name

	err = s.db.Save(&categoryToUpdate).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *categoryRepository) DeleteCategory(categoryID int) error {
	var category Category
	err := s.db.First(&category, categoryID).Error
	if err != nil {
		return err
	}

	err = s.db.Delete(&category).Error
	if err != nil {
		return err
	}

	return nil
}
