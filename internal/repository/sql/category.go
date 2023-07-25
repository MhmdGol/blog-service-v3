package sql

import (
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/repository"
	"blog-service-v3/internal/repository/sql/dbmodel"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

// &CategoryRepository{}
var _ repository.CategoryRepository = (*CategoryRepository)(nil)

func NewCategoryRepo(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (cr *CategoryRepository) Create(c model.Category) error {
	err := cr.db.Create(&dbmodel.Category{
		Name: c.Name,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoryRepository) All() ([]model.Category, error) {
	var categories []dbmodel.Category

	err := cr.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	result := make([]model.Category, len(categories))
	for i, c := range categories {
		result[i] = model.Category{
			ID:   model.ID(c.ID),
			Name: c.Name,
		}
	}
	return result, nil
}

func (cr *CategoryRepository) UpdateByID(c model.Category) error {
	var categoryToUpdate dbmodel.Category
	err := cr.db.First(&categoryToUpdate, c.ID).Error
	if err != nil {
		return err
	}

	categoryToUpdate.Name = c.Name

	err = cr.db.Save(&categoryToUpdate).Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoryRepository) DeleteByID(id model.ID) error {
	var category dbmodel.Category
	err := cr.db.First(&category, id).Error
	if err != nil {
		return err
	}

	err = cr.db.Delete(&category).Error
	if err != nil {
		return err
	}

	return nil
}