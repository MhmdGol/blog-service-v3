package nosql

import (
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/repository"
	"blog-service-v3/internal/repository/nosql/nosqlmodel"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type PostRepository struct {
	db     *mongo.Database
	ctx    context.Context
	logger *zap.Logger
}

var _ repository.PostRepository = (*PostRepository)(nil)

func NewPostRepo(db *mongo.Database, ctx context.Context, logger *zap.Logger) *PostRepository {
	return &PostRepository{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (pr *PostRepository) Create(p model.Post) error {

	cats := make([]*nosqlmodel.Category, len(p.Categories))

	for i, c := range p.Categories {
		var findCat nosqlmodel.Category
		pr.db.Collection("categories").FindOne(pr.ctx, bson.M{"name": c}).Decode(&findCat)
		if findCat.Name == "" {
			cats[i] = &nosqlmodel.Category{Name: c}
		} else {
			cats[i] = &findCat
		}
	}

	_, err := pr.db.Collection("posts").InsertOne(pr.ctx, &nosqlmodel.Post{
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
	cursor, err := pr.db.Collection("posts").Find(pr.ctx, bson.D{})
	if err != nil {
		return []model.Post{}, err
	}

	var posts []nosqlmodel.Post
	if err := cursor.All(pr.ctx, &posts); err != nil {
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
	findOptions := options.Find()
	findOptions.SetSkip((int64)((pageNumber - 1) * pageSize))
	findOptions.SetLimit((int64)(pageSize))

	cursor, err := pr.db.Collection("posts").Find(pr.ctx, bson.D{}, findOptions)
	if err != nil {
		return []model.Post{}, err
	}

	var posts []nosqlmodel.Post
	if err := cursor.All(pr.ctx, &posts); err != nil {
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
	var postToUpdate nosqlmodel.Post
	err := pr.db.Collection("posts").FindOne(pr.ctx, bson.M{"_id": p.ID}).Decode(&postToUpdate)
	if err != nil {
		return err
	}

	cats := make([]*nosqlmodel.Category, len(p.Categories))

	for i, c := range p.Categories {
		var findCat nosqlmodel.Category
		pr.db.Collection("categories").FindOne(pr.ctx, bson.M{"name": c}).Decode(&findCat)
		if findCat.Name == "" {
			cats[i] = &nosqlmodel.Category{Name: c}
		} else {
			cats[i] = &findCat
		}
	}

	postToUpdate.Title = p.Title
	postToUpdate.Text = p.Text
	postToUpdate.Categories = cats

	_, err = pr.db.Collection("posts").UpdateOne(pr.ctx, bson.M{"_id": p.ID}, postToUpdate)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) DeleteByID(id model.ID) error {
	_, err := pr.db.Collection("posts").DeleteOne(pr.ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
