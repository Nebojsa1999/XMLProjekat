package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TypeOfConnection string
const (
	UndefinedTypeOfConnection TypeOfConnection = ""
	Following                                  = "Following"
	Blocking                                   = "Blocking"
)

type Connection struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	Type       TypeOfConnection   `json:"type" bson:"type"`
	IssuerId   primitive.ObjectID `json:"issuerId" bson:"issuer_id"`
	SubjectId  primitive.ObjectID `json:"subjectId" bson:"subject_id"`
	Date       time.Time          `json:"date" bson:"date"`
	IsApproved bool               `json:"isApproved" bson:"is_approved"`
}

type ConnectionUpdateDTO struct {
	Type       TypeOfConnection
	IssuerId   primitive.ObjectID
	SubjectId  primitive.ObjectID
	IsApproved bool
}
