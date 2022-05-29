package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserRole string

const (
	CommonUser    UserRole = "CommonUser"
	CompanyOwner           = "CompanyOwner"
	Administrator          = "Administrator"
)

type Gender string

const (
	Undefined Gender = ""
	Male             = "Male"
	Female           = "Female"
)

type User struct {
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	Role           UserRole           `json:"role" bson:"role"`
	Username       string             `json:"username" bson:"username"`
	Password       string             `json:"password" bson:"password"`
	FirstName      string             `json:"firstName" bson:"first_name"`
	LastName       string             `json:"lastName" bson:"last_name"`
	Email          string             `json:"email" bson:"email"`
	Phone          string             `json:"phone" bson:"phone"`
	Gender         Gender             `json:"gender" bson:"gender"`
	DateOfBirth    time.Time          `json:"dateOfBirth" bson:"date_of_birth"`
	Biography      string             `json:"biography" bson:"biography"`
	WorkExperience string             `json:"workExperience" bson:"work_experience"`
	Education      string             `json:"education" bson:"education"`
	Skills         string             `json:"skills" bson:"skills"`
	Interests      string             `json:"interests" bson:"interests"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AgentAppToken struct {
	Token string `json:"token"`
}

type Company struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	OwnerId     primitive.ObjectID `json:"ownerId" bson:"owner_id"`
	Name        string             `json:"name" bson:"name"`
	Address     string             `json:"address" bson:"address"`
	Email       string             `json:"email" bson:"email"`
	Phone       string             `json:"phone" bson:"phone"`
	AreaOfWork  string             `json:"areaOfWork" bson:"area_of_work"`
	Description string             `json:"description" bson:"description"`
	WorkCulture string             `json:"workCulture" bson:"work_culture"`
}

type CompanyRegistrationRequestStatus string

const (
	Pending  CompanyRegistrationRequestStatus = "Pending"
	Accepted                                  = "Accepted"
	Rejected                                  = "Rejected"
)

type CompanyRegistrationRequest struct {
	OwnerId primitive.ObjectID               `json:"ownerId" bson:"owner_id"`
	Status  CompanyRegistrationRequestStatus `json:"status" bson:"status"`
}
