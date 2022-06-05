package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Connection struct {
	Id      primitive.ObjectID `bson:"id"`
	UserAId primitive.ObjectID `bson:"userAId"`
	UserBId primitive.ObjectID `bson:"userBId"`
}

type ProfilePrivacy struct {
	Id        primitive.ObjectID `bson:"id"`
	UserId    primitive.ObjectID `bson:"userId"`
	IsPrivate bool               `bson:"isPrivate"`
}
