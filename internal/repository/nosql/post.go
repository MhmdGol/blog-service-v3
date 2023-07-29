package nosql

import (
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/repository"
	"blog-service-v3/internal/repository/nosql/nosqlmodel"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type PostRepository struct {
	db     *mongo.Database
	logger *zap.Logger
}

var _ repository.PostRepository = (*PostRepository)(nil)

func NewPostRepo(db *mongo.Database, logger *zap.Logger) *PostRepository {
	logger.Info("repository, post, NewPostRepo")
	return &PostRepository{
		db:     db,
		logger: logger,
	}
}

func (pr *PostRepository) Create(p model.Post) error {
	pr.logger.Info("repository, post, Create")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cats := make([]*nosqlmodel.Category, len(p.Categories))

	for i, c := range p.Categories {
		var findCat nosqlmodel.Category
		pr.db.Collection("categories").FindOne(ctx, bson.M{"name": c}).Decode(&findCat)
		if findCat.Name == "" {
			cats[i] = &nosqlmodel.Category{Name: c}
		} else {
			cats[i] = &findCat
		}
	}

	_, err := pr.db.Collection("posts").InsertOne(ctx, &nosqlmodel.Post{
		Title:      p.Title,
		Text:       p.Text,
		Categories: cats,
	})
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) All() ([]model.Post, error) {
	pr.logger.Info("repository, post, All")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := pr.db.Collection("posts").Find(ctx, bson.D{})
	if err != nil {
		return []model.Post{}, err
	}

	var posts []nosqlmodel.Post
	if err := cursor.All(ctx, &posts); err != nil {
		return []model.Post{}, err
	}

	result := make([]model.Post, len(posts))
	for i, p := range posts {
		result[i] = model.Post{
			ID:    model.ID(p.ID[:]),
			Title: p.Title,
			Text:  p.Text,
			Categories: func(c []*nosqlmodel.Category) []string {
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
	pr.logger.Info("repository, post, Paginated")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSkip((int64)((pageNumber - 1) * pageSize))
	findOptions.SetLimit((int64)(pageSize))

	cursor, err := pr.db.Collection("posts").Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return []model.Post{}, err
	}

	var posts []nosqlmodel.Post
	if err := cursor.All(ctx, &posts); err != nil {
		return []model.Post{}, err
	}

	result := make([]model.Post, len(posts))
	for i, p := range posts {
		result[i] = model.Post{
			ID:    model.ID(p.ID[:]),
			Title: p.Title,
			Text:  p.Text,
			Categories: func(c []*nosqlmodel.Category) []string {
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
	pr.logger.Info("repository, post, UpdateByID")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var postToUpdate nosqlmodel.Post
	err := pr.db.Collection("posts").FindOne(ctx, bson.M{"_id": p.ID}).Decode(&postToUpdate)
	if err != nil {
		return err
	}

	cats := make([]*nosqlmodel.Category, len(p.Categories))

	for i, c := range p.Categories {
		var findCat nosqlmodel.Category
		pr.db.Collection("categories").FindOne(ctx, bson.M{"name": c}).Decode(&findCat)
		if findCat.Name == "" {
			cats[i] = &nosqlmodel.Category{Name: c}
		} else {
			cats[i] = &findCat
		}
	}

	postToUpdate.Title = p.Title
	postToUpdate.Text = p.Text
	postToUpdate.Categories = cats

	_, err = pr.db.Collection("posts").UpdateOne(ctx, bson.M{"_id": p.ID}, postToUpdate)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) DeleteByID(id model.ID) error {
	pr.logger.Info("repository, post, DeleteByID")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := pr.db.Collection("posts").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
