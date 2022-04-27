package model

import "time"

type Gender string

const (
	Undefined = ""
	Male = "Male"
	Female = "Female"
)

type User struct {
	Id uint   `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"-"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Gender Gender `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Biography string `json:"biography"`
	WorkExperience string `json:"work_experience"`
	Education string `json:"education"`
	Skills string `json:"skills"`
	Interests string `json:"interests"`
}
