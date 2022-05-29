package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	Id           primitive.ObjectID `bson:"_id"`
	UserId       primitive.ObjectID `bson:"userId"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
	Position     string             `bson:"position"`
	Description  string             `bson:"description"`
	Requirements string             `bson:"requirements"`
}
