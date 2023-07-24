package service

import (
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/repository"
	"blog-service-v3/internal/service"
)

type CategoryService struct {
	catRepo repository.CategoryRepository
}

var _ service.CategoryService = (*CategoryService)(nil)

func NewCategoryService(r repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		catRepo: r,
	}
}

func (cs *CategoryService) Create(c model.Category) error {
	return cs.catRepo.Create(c)
}

func (cs *CategoryService) All() ([]model.Category, error) {
	return cs.catRepo.All()
}

func (cs *CategoryService) UpdateByID(c model.Category) error {
	return cs.catRepo.UpdateByID(c)
}

func (cs *CategoryService) DeleteByID(id model.ID) error {
	return cs.catRepo.DeleteByID(id)
}
