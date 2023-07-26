package nosqlmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Title      string             `bson:"title"`
	Text       string             `bson:"text"`
	Categories []*Category        `bson:"categories"`
}
