package mongo_project

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewConnectMongoClient() (*mongo.Collection, error) {
	connectURII := "mongodb://root:root@localhost:27017"
	clientOptions := options.Client().ApplyURI(connectURII)

	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	return client.Database("grpc").Collection("product"), nil
}
