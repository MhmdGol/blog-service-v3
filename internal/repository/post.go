package repository

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string      `json:"title"`
	Text       string      `json:"text"`
	Categories []*Category `json:"cats" gorm:"many2many:post_categories;"`
}

func (s *postRepository) CreatePost(title, text string, categories []string) error {
	var cats []*Category
	for _, item := range categories {
		var findCat Category
		s.db.Where("name = ?", item).First(&findCat)

		if findCat.ID == 0 {
			cats = append(cats, &Category{Name: item})
		} else {
			cats = append(cats, &findCat)
		}
	}

	err := s.db.Create(&Post{
		Title:      title,
		Text:       text,
		Categories: cats,
	}).Error

	if err != nil {
		return err
	}
	return nil
}

func (s *postRepository) AllPosts() ([]Post, error) {
	var posts []Post

	err := s.db.Preload("Categories").Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *postRepository) PagePosts(pageNumber, pageSize int) ([]Post, error) {
	var posts []Post
	err := s.db.Order("updated_at desc").Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *postRepository) UpdatePost(postID int, title, text string, categories []string) error {
	var postToUpdate Post
	err := s.db.Preload("Categories").Where("id = ?", postID).First(&postToUpdate).Error
	if err != nil {
		return err
	}

	var cats []*Category
	for _, item := range categories {
		var findCat Category
		s.db.Where("name = ?", item).First(&findCat)

		if findCat.ID == 0 {
			cats = append(cats, &Category{Name: item})
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

func (s *postRepository) DeletePost(postID int) error {
	var post Post
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
