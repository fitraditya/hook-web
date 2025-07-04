package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type M map[string]interface{}

type RequestInfo struct {
	Method        string `json:"method" bson:"method"`
	Path          string `json:"path" bson:"path"`
	ContentType   string `json:"contentType,omitempty" bson:"contentType"`
	ContentLength int64  `json:"contentLength,omitempty" bson:"contentLength"`
	IP            string `json:"ip" bson:"ip"`
	Headers       M      `json:"headers" bson:"headers"`
	Query         M      `json:"query,omitempty" bson:"query"`
	Data          string `json:"data,omitempty" bson:"data"`
	Form          M      `json:"form,omitempty" bson:"form"`
	Body          M      `json:"body,omitempty" bson:"body"`
}

type Request struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Slug    string             `json:"-" bson:"slug"`
	Request *RequestInfo       `json:"request" bson:"request"`
	Created time.Time          `json:"created" bson:"created"`
}
