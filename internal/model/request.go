package model

import (
	"context"
	"fmt"
	"time"

	"github.com/fitraditya/hook-web/internal/constant"
	"github.com/fitraditya/hook-web/internal/schema"
	"github.com/fitraditya/hook-web/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Request struct {
	mongo *service.MongoDB
}

func NewRequest(mongo *service.MongoDB) *Request {
	return &Request{
		mongo: mongo,
	}
}

func (r *Request) getCollection() *mongo.Collection {
	c := r.mongo.GetCollection(constant.CollectionRequest)
	return c
}

func (r *Request) List(ctx context.Context, slug string) ([]*schema.Request, error) {
	c := r.getCollection()

	option := options.Find()
	option.SetLimit(100)
	option.SetSkip(0)
	option.SetSort(bson.D{
		bson.E{
			Key:   "_id",
			Value: -1,
		},
	})

	filter := bson.M{"slug": slug}

	cursor, err := c.Find(ctx, filter, option)
	if err != nil {
		return nil, err
	}

	fmt.Println("AYAM")

	result := []*schema.Request{}

	if cursor != nil {
		err = cursor.All(ctx, &result)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println("SAPI")

	return result, err
}

func (r *Request) Create(ctx context.Context, slug string, info *schema.RequestInfo) (*schema.Request, error) {
	c := r.getCollection()

	now := time.Now()

	request := &schema.Request{
		Request: info,
	}

	request.ID = primitive.NewObjectID()
	request.Slug = slug

	request.Created = now

	_, err := c.InsertOne(ctx, request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
