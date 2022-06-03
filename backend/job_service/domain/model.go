package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Job struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	UserId       primitive.ObjectID `json:"userId" bson:"user_id"`
	CreatedAt    time.Time          `json:"createdAt" bson:"created_at"`
	Position     string             `json:"position" bson:"position"`
	Description  string             `json:"description" bson:"description"`
	Requirements string             `json:"requirements" bson:"requirements"`
}
