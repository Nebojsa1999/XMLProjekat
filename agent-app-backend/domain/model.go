package domain

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	Role           enums.UserRole     `json:"role" bson:"role"`
	Username       string             `json:"username" bson:"username"`
	Password       string             `json:"password" bson:"password"`
	FirstName      string             `json:"firstName" bson:"first_name"`
	LastName       string             `json:"lastName" bson:"last_name"`
	Email          string             `json:"email" bson:"email"`
	Phone          string             `json:"phone" bson:"phone"`
	Gender         enums.Gender       `json:"gender" bson:"gender"`
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

type CompanyRegistrationRequest struct {
	Id                 primitive.ObjectID                     `json:"id" bson:"_id"`
	OwnerId            primitive.ObjectID                     `json:"ownerId" bson:"owner_id"`
	Status             enums.CompanyRegistrationRequestStatus `json:"status" bson:"status"`
	ReasonForRejection string                                 `json:"reasonForRejection" bson:"reason_for_rejection"`
	Name               string                                 `json:"name" bson:"name"`
	Address            string                                 `json:"address" bson:"address"`
	Email              string                                 `json:"email" bson:"email"`
	Phone              string                                 `json:"phone" bson:"phone"`
	AreaOfWork         string                                 `json:"areaOfWork" bson:"area_of_work"`
	Description        string                                 `json:"description" bson:"description"`
	WorkCulture        string                                 `json:"workCulture" bson:"work_culture"`
}
