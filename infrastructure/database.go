package infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	NoSQL *mongo.Database
}

func NewDatabase(env Env) (Database, error) {
	ctx := context.Background()

	clientOptions := options.Client()
	clientOptions.ApplyURI(env.MongoDSN)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return Database{}, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return Database{}, err
	}

	return Database{
		NoSQL: client.Database(env.MongoDatabase),
	}, nil

}
