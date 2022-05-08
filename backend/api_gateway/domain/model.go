package domain

import (
	postPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	userPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
	"time"
)

type User struct {
	Id             string
	Username       string
	Password       string
	IsPrivate      bool
	FirstName      string
	LastName       string
	Email          string
	Gender         string
	DateOfBirth    time.Time
	Biography      string
	WorkExperience string
	Education      string
	Skills         string
	Interests      string
}

type UserRegistrationRequest struct {
	User userPb.User
}

type UserStatusRequest struct {
	Id        string
	IsPrivate bool
	Posts     []*postPb.Post
}

type GetAllPostsRequest struct {
	UserIds []string
	Posts   []*postPb.Post
}
