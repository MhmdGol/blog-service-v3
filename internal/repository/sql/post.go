package sql

import (
	"blog-service-v3/internal/repository/sql/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRopo(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (s *PostRepository) CreatePost(title, text string, categories []string) error {
	var cats []*model.Category
	for _, item := range categories {
		var findCat model.Category
		s.db.Where("name = ?", item).First(&findCat)

		if findCat.ID == 0 {
			cats = append(cats, &model.Category{Name: item})
		} else {
			cats = append(cats, &findCat)
		}
	}

	err := s.db.Create(&model.Post{
		Title:      title,
		Text:       text,
		Categories: cats,
	}).Error

	if err != nil {
		return err
	}
	return nil
}

func (s *PostRepository) AllPosts() ([]model.Post, error) {
	var posts []model.Post

	err := s.db.Preload("Categories").Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostRepository) PagePosts(pageNumber, pageSize int) ([]model.Post, error) {
	var posts []model.Post
	err := s.db.Order("updated_at desc").Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostRepository) UpdatePost(postID int, title, text string, categories []string) error {
	var postToUpdate model.Post
	err := s.db.Preload("Categories").Where("id = ?", postID).First(&postToUpdate).Error
	if err != nil {
		return err
	}

	var cats []*model.Category
	for _, item := range categories {
		var findCat model.Category
		s.db.Where("name = ?", item).First(&findCat)

		if findCat.ID == 0 {
			cats = append(cats, &model.Category{Name: item})
		} else {
			cats = append(cats, &findCat)
		}
	}

	postToUpdate.Title = title
	postToUpdate.Text = text
	postToUpdate.Categories = cats

	err = s.db.Save(&postToUpdate).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *PostRepository) DeletePost(postID int) error {
	var post model.Post
	err := s.db.First(&post, postID).Error
	if err != nil {
		return err
	}

	err = s.db.Delete(&post).Error
	if err != nil {
		return err
	}

	return nil
}
