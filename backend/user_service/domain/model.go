package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Gender string
const (
	Undefined Gender = ""
	Male = "Male"
	Female = "Female"
)

type User struct {
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	Username       string             `json:"username" bson:"username"`
	Password       string             `json:"password" bson:"password"`
	IsPrivate      bool               `json:"is_private" bson:"is_private"`
	FirstName      string             `json:"first_name" bson:"first_name"`
	LastName       string             `json:"last_name" bson:"last_name"`
	Email          string             `json:"email" bson:"email"`
	Gender         Gender             `json:"gender" bson:"gender"`
	DateOfBirth    time.Time          `json:"date_of_birth" bson:"date_of_birth"`
	Biography      string             `json:"biography" bson:"biography"`
	WorkExperience string             `json:"work_experience" bson:"work_experience"`
	Education      string             `json:"education" bson:"education"`
	Skills         string             `json:"skills" bson:"skills"`
	Interests      string             `json:"interests" bson:"interests"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTToken struct {
	Token string `json:"token"`
}
