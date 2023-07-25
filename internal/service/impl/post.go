package service

import (
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/repository"
	"blog-service-v3/internal/service"

	"go.uber.org/zap"
)

type PostService struct {
	postRepo repository.PostRepository
	logger   *zap.Logger
}

var _ service.PostService = (*PostService)(nil)

func NewPostService(r repository.PostRepository, logger *zap.Logger) *PostService {
	logger.Info("NewPostService")

	return &PostService{
		postRepo: r,
		logger:   logger,
	}
}

func (ps *PostService) Create(p model.Post) error {
	ps.logger.Info("ps.Create")

	return ps.postRepo.Create(p)
}

func (ps *PostService) All() ([]model.Post, error) {
	ps.logger.Info("ps.All")

	return ps.postRepo.All()
}

func (ps *PostService) Paginated(pageNumber, pageSize int) ([]model.Post, error) {
	ps.logger.Info("ps.Paginated")

	return ps.postRepo.Paginated(pageNumber, pageSize)
}

func (ps *PostService) UpdateByID(p model.Post) error {
	ps.logger.Info("ps.UpdateByID")

	return ps.postRepo.UpdateByID(p)
}

func (ps *PostService) DeleteByID(id model.ID) error {
	ps.logger.Info("ps.DeleteByID")

	return ps.postRepo.DeleteByID(id)
}
