package domain

import (
	postingPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	userPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"

	"time"
)

type User struct {
	Id                string
	Username          string
	Password          string
	IsPrivate         bool
	FirstName         string
	LastName          string
	Email             string
	Gender            string
	DateOfBirth       time.Time
	Biography         string
	WorkExperience    string
	Education         string
	Skills            string
	Interests         string
	JobOffersAPIToken string
}

type UserRegistrationRequest struct {
	User userPb.User
}

type UserStatusRequest struct {
	Id        string
	IsPrivate bool
	Posts     []*postingPb.Post
}

type GetAllPostsRequest struct {
	UserIds []string
	Posts   []*postingPb.Post
}

type Job struct {
	Id           string
	UserId       string
	CreatedAt    time.Time
	Position     string
	Description  string
	Requirements string
}

type PostJobOfferRequest struct {
	JobOffersAPIToken string
	Job               *Job
}

type JobOffersAPITokenStatusRequest struct {
	UserId            string
	HasGeneratedToken bool
}
