package service

import (
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/repository"
	"blog-service-v3/internal/service"
)

type PostService struct {
	postRepo repository.PostRepository
}

var _ service.PostService = (*PostService)(nil)

func NewPostService(r repository.PostRepository) *PostService {
	return &PostService{
		postRepo: r,
	}
}

func (ps *PostService) Create(p model.Post) error {
	return ps.postRepo.Create(p)
}

func (ps *PostService) All() ([]model.Post, error) {
	return ps.postRepo.All()
}

func (ps *PostService) Paginated(pageNumber, pageSize int) ([]model.Post, error) {
	return ps.postRepo.Paginated(pageNumber, pageSize)
}

func (ps *PostService) UpdateByID(p model.Post) error {
	return ps.postRepo.UpdateByID(p)
}

func (ps *PostService) DeleteByID(id model.ID) error {
	return ps.postRepo.DeleteByID(id)
}
