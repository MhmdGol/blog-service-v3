package service

import (
	"blog-service-v3/internal/repository"
)

type PostApp interface {
	CreatePost(title, text string, categories []string) error
	AllPosts() ([]Post, error)
	PagePosts(pageNumber, pageSize int) ([]Post, error)
	UpdatePost(postID int, title, text string, categories []string) error
	DeletePost(postID int) error
}

type CategoryApp interface {
	CreateCategory(name string) error
	AllCategories() ([]Category, error)
	UpdateCategory(categoryID int, name string) error
	DeleteCategory(categoryID int) error
}

type postApp struct {
	postRepo repository.PostRepository
}

type categoryApp struct {
	catRepo repository.CategoryRepository
}

func PostNew(r repository.PostRepository) *postApp {
	return &postApp{
		postRepo: r,
	}
}

func CategoryNew(r repository.CategoryRepository) *categoryApp {
	return &categoryApp{
		catRepo: r,
	}
}
