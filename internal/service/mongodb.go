package service

import (
	"context"

	"github.com/fitraditya/hook-web/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDB() *MongoDB {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(config.GetMongoURL()).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.PrimaryPreferred())
	if err != nil {
		panic(err)
	}

	return &MongoDB{
		client: client,
	}
}

func (m *MongoDB) GetCollection(name string) *mongo.Collection {
	return m.client.Database(config.GetMongoDatabase()).Collection(name)
}
