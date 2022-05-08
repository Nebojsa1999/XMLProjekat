package domain

import (
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
