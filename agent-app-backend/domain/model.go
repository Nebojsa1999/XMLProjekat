package domain

import (
	"time"

	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id                     primitive.ObjectID `json:"id" bson:"_id"`
	Role                   enums.UserRole     `json:"role" bson:"role"`
	OwnedCompanyId         primitive.ObjectID `json:"ownedCompanyId" bson:"owned_company_id"`
	IssuedCompanyRequestId primitive.ObjectID `json:"issuedCompanyRequestId" bson:"issued_company_request_id"`
	Username               string             `json:"username" bson:"username"`
	Password               string             `json:"password" bson:"password"`
	FirstName              string             `json:"firstName" bson:"first_name"`
	LastName               string             `json:"lastName" bson:"last_name"`
	Email                  string             `json:"email" bson:"email"`
	Phone                  string             `json:"phone" bson:"phone"`
	Gender                 enums.Gender       `json:"gender" bson:"gender"`
	DateOfBirth            time.Time          `json:"dateOfBirth" bson:"date_of_birth"`
	Biography              string             `json:"biography" bson:"biography"`
	WorkExperience         string             `json:"workExperience" bson:"work_experience"`
	Education              string             `json:"education" bson:"education"`
	Skills                 string             `json:"skills" bson:"skills"`
	Interests              string             `json:"interests" bson:"interests"`
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

type Job struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	CompanyId    primitive.ObjectID `json:"companyId" bson:"company_id"`
	CreatedAt    time.Time          `json:"createdAt" bson:"created_at"`
	Position     string             `json:"position" bson:"position"`
	Description  string             `json:"description" bson:"description"`
	Requirements string             `json:"requirements" bson:"requirements"`
	Comments     []Comment          `json:"comments" bson:"comments"`
	Wages        []Wage             `json:"wages" bson:"wages"`
	Interviews   []Interview        `json:"interviews" bson:"interviews"`
}

type Wage struct {
	Id              primitive.ObjectID    `json:"id" bson:"_id"`
	CompanyId       primitive.ObjectID    `json:"companyId" bson:"company_id"`
	Position        enums.Position        `json:"poition" bson:"position"`
	Engagement      enums.Engagement      `json:"engagement" bson:"engagement"`
	ExperienceLevel enums.ExperienceLevel `json:"experience_level" bson:"experience_level"`
	NetoWage        string                `json:"netoWage" bson:"neto_wage"`
}

type Interview struct {
	Id                 primitive.ObjectID `json:"id" bson:"_id"`
	CompanyId          primitive.ObjectID `json:"companyId" bson:"company_id"`
	Position           enums.Position     `json:"poition" bson:"position"`
	Title              string             `json:"title" bson:"title"`
	YearOfInterview    string             `json:"yearOfInterview" bson:"year_of_interview"`
	HRInterview        string             `json:"hrInterview" bson:"hr_interview"`
	TechnicalInterview string             `json:"technicalInterview" bson:"technical_interview"`
}

type Comment struct {
	Id              primitive.ObjectID    `json:"id" bson:"_id"`
	CompanyId       primitive.ObjectID    `json:"companyId" bson:"company_id"`
	Position        enums.Position        `json:"poition" bson:"position"`
	Engagement      enums.Engagement      `json:"engagement" bson:"engagement"`
	ExperienceLevel enums.ExperienceLevel `json:"experience_level" bson:"experience_level"`
	Content         string                `json:"content" bson:"content"`
}
