package nosql

import (
	"blog-service-v3/internal/model"
	"blog-service-v3/internal/repository"
	"blog-service-v3/internal/repository/nosql/nosqlmodel"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type CategoryRepository struct {
	db     *mongo.Database
	ctx    context.Context
	logger *zap.Logger
}

var _ repository.CategoryRepository = (*CategoryRepository)(nil)

func NewCategoryRepo(db *mongo.Database, ctx context.Context, logger *zap.Logger) *CategoryRepository {
	return &CategoryRepository{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (cr *CategoryRepository) Create(c model.Category) error {
	_, err := cr.db.Collection("categories").InsertOne(cr.ctx, &nosqlmodel.Category{
		Name: c.Name,
	})
	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoryRepository) All() ([]model.Category, error) {
	cursor, err := cr.db.Collection("categories").Find(cr.ctx, bson.D{})
	if err != nil {
		return []model.Category{}, err
	}
	defer cursor.Close(cr.ctx)

	var categories []nosqlmodel.Category
	if err := cursor.All(cr.ctx, &categories); err != nil {
		return []model.Category{}, err
	}

	result := make([]model.Category, len(categories))
	for i, c := range categories {
		result[i] = model.Category{
			ID:   model.ID(c.ID[:]),
			Name: c.Name,
		}
	}

	return result, nil
}

func (cr *CategoryRepository) UpdateByID(c model.Category) error {
	var categoryToUpdate nosqlmodel.Category
	err := cr.db.Collection("categories").FindOne(cr.ctx, bson.M{"_id": c.ID}).Decode(&categoryToUpdate)
	if err != nil {
		return err
	}

	categoryToUpdate.Name = c.Name

	_, err = cr.db.Collection("categories").UpdateOne(cr.ctx, bson.M{"_id": c.ID}, categoryToUpdate)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoryRepository) DeleteByID(id model.ID) error {
	_, err := cr.db.Collection("categories").DeleteOne(cr.ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
