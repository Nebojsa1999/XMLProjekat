package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Connection struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	IssuerId   primitive.ObjectID `json:"issuerId" bson:"issuer_id"`
	SubjectId  primitive.ObjectID `json:"subjectId" bson:"subject_id"`
	Date       time.Time          `json:"date" bson:"date"`
	IsApproved bool               `json:"isApproved" bson:"is_approved"`
}

type ProfilePrivacy struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	UserId    primitive.ObjectID `json:"userId" bson:"user_id"`
	IsPrivate bool               `json:"isPrivate" bson:"is_private"`
}
