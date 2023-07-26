package sql

import (
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/repository"
	"blog-service-v3/internal/repository/sql/sqlmodel"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

var _ repository.PostRepository = (*PostRepository)(nil)

func NewPostRopo(db *gorm.DB, logger *zap.Logger) *PostRepository {
	logger.Info("NewPostRopo")
	return &PostRepository{
		db:     db,
		logger: logger,
	}
}

func (pr *PostRepository) Create(p model.Post) error {

	cats := make([]*sqlmodel.Category, len(p.Categories))

	for i, c := range p.Categories {
		var findCat sqlmodel.Category
		pr.db.Where("name = ?", c).First(&findCat)

		if findCat.ID == 0 {
			cats[i] = &sqlmodel.Category{Name: c}
		} else {
			cats[i] = &findCat
		}
	}

	err := pr.db.Create(&sqlmodel.Post{
		Title:      p.Title,
		Text:       p.Text,
		Categories: cats,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) All() ([]model.Post, error) {
	var posts []sqlmodel.Post

	err := pr.db.Preload("Categories").Find(&posts).Error
	if err != nil {
		return nil, err
	}

	result := make([]model.Post, len(posts))
	for i, p := range posts {
		result[i] = model.Post{
			ID:    model.ID(p.ID),
			Title: p.Title,
			Text:  p.Text,
			Categories: func(c []*sqlmodel.Category) []string {
				cstr := make([]string, len(c))
				for i2, c2 := range c {
					cstr[i2] = c2.Name
				}
				return cstr
			}(p.Categories),
		}
	}

	return result, nil
}

func (pr *PostRepository) Paginated(pageNumber, pageSize int) ([]model.Post, error) {
	var posts []sqlmodel.Post

	err := pr.db.Order("updated_at desc").Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	result := make([]model.Post, len(posts))
	for i, p := range posts {
		result[i] = model.Post{
			ID:    model.ID(p.ID),
			Title: p.Title,
			Text:  p.Text,
			Categories: func(c []*sqlmodel.Category) []string {
				cstr := make([]string, len(c))
				for i2, c2 := range c {
					cstr[i2] = c2.Name
				}
				return cstr
			}(p.Categories),
		}
	}

	return result, nil
}

func (pr *PostRepository) UpdateByID(p model.Post) error {
	var postToUpdate sqlmodel.Post
	err := pr.db.Preload("Categories").Where("id = ?", p.ID).First(&postToUpdate).Error
	if err != nil {
		return err
	}

	cats := make([]*sqlmodel.Category, len(p.Categories))

	for i, c := range p.Categories {
		var findCat sqlmodel.Category
		pr.db.Where("name = ?", c).First(&findCat)

		if findCat.ID == 0 {
			cats[i] = &sqlmodel.Category{Name: c}
		} else {
			cats[i] = &findCat
		}
	}

	postToUpdate.Title = p.Title
	postToUpdate.Text = p.Text
	postToUpdate.Categories = cats

	err = pr.db.Save(&postToUpdate).Error
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) DeleteByID(id model.ID) error {
	var post sqlmodel.Post
	err := pr.db.First(&post, id).Error
	if err != nil {
		return err
	}

	err = pr.db.Delete(&post).Error
	if err != nil {
		return err
	}

	return nil
}
