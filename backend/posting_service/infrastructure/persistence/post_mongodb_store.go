package persistence

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "post"
	COLLECTION = "posts"
)

type PostMongoDBStore struct {
	posts *mongo.Database
}
