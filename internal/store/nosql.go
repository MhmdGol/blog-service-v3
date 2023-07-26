package store

import (
	"blog-service-v3/internal/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewNosqlStorage(ctx context.Context, conf config.NoSQLDtabaseConfig) (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", conf.Host, conf.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(conf.Name), nil
}

//defer store.CloseNosql(...)
func CloseNosql(ctx context.Context, client *mongo.Client) error {
	if err := client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}
