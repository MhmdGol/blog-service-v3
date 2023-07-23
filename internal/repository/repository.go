package repository

import (
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(title, text string, categories []string) error
	AllPosts() ([]Post, error)
	PagePosts(pageNumber, pageSize int) ([]Post, error)
	UpdatePost(postID int, title, text string, categories []string) error
	DeletePost(postID int) error
}

type CategoryRepository interface {
	CreateCategory(name string) error
	AllCategories() ([]Category, error)
	UpdateCategory(categoryID int, name string) error
	DeleteCategory(categoryID int) error
}

type categoryRepository struct {
	db *gorm.DB
}

type postRepository struct {
	db *gorm.DB
}

func CategoryNew(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func PostNew(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}
