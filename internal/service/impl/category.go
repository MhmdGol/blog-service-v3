package service

import (
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/repository"
	"blog-service-v3/internal/service"

	"go.uber.org/zap"
)

type CategoryService struct {
	catRepo repository.CategoryRepository
	logger  *zap.Logger
}

var _ service.CategoryService = (*CategoryService)(nil)

func NewCategoryService(r repository.CategoryRepository, logger *zap.Logger) *CategoryService {
	logger.Info("NewCategoryService")

	return &CategoryService{
		catRepo: r,
		logger:  logger,
	}
}

func (cs *CategoryService) Create(c model.Category) error {
	cs.logger.Info("cs.Create")

	return cs.catRepo.Create(c)
}

func (cs *CategoryService) All() ([]model.Category, error) {
	cs.logger.Info("cs.All")

	return cs.catRepo.All()
}

func (cs *CategoryService) UpdateByID(c model.Category) error {
	cs.logger.Info("cs.UpdateByID")

	return cs.catRepo.UpdateByID(c)
}

func (cs *CategoryService) DeleteByID(id model.ID) error {
	cs.logger.Info("cs.DeleteByID")

	return cs.catRepo.DeleteByID(id)
}
