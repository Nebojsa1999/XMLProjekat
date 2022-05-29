package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserRole string

const (
	COMMON_USER   UserRole = "COMMON__USER"
	COMPANY_OWNER          = "COMPANY_OWNER"
	ADMINISTRATOR          = "ADMINISTRATOR"
)

type Gender string

const (
	Undefined Gender = ""
	Male             = "Male"
	Female           = "Female"
)

type User struct {
	Id                 primitive.ObjectID `json:"id" bson:"_id"`
	Role               UserRole           `json:"role" bson:"role"`
	Username           string             `json:"username" bson:"username"`
	Password           string             `json:"password" bson:"password"`
	FirstName          string             `json:"firstName" bson:"first_name"`
	LastName           string             `json:"lastName" bson:"last_name"`
	Email              string             `json:"email" bson:"email"`
	Phone              string             `json:"phone" bson:"phone"`
	Gender             Gender             `json:"gender" bson:"gender"`
	DateOfBirth        time.Time          `json:"dateOfBirth" bson:"date_of_birth"`
	Biography          string             `json:"biography" bson:"biography"`
	WorkExperience     string             `json:"workExperience" bson:"work_experience"`
	Education          string             `json:"education" bson:"education"`
	Skills             string             `json:"skills" bson:"skills"`
	Interests          string             `json:"interests" bson:"interests"`
	CompanyDescription string             `json:"companyDescription" bson:"company_description"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AgentAppToken struct {
	Token string `json:"token"`
}
