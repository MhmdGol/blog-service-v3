package repository

import "blog-service-v3/internal/model"

type PostRepository interface {
	Create(model.Post) error
	All() ([]model.Post, error)
	Paginated(pageNumber, pageSize int) ([]model.Post, error)
	UpdateByID(model.Post) error
	DeleteByID(id model.ID) error
}

type CategoryRepository interface {
	Create(model.Category) error
	All() ([]model.Category, error)
	UpdateByID(model.Category) error
	DeleteByID(id model.ID) error
}
