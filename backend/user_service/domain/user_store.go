package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserStore interface {
	GetAll() ([]*User, error)
	Get(id primitive.ObjectID) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByUsernameAndPassword(username string, password string) (*User, error)
	GetByEmail(email string) (*User, error)
	RegisterANewUser(user *User) (string, error)
	Update(updatedUser *User) (string, *User, error)
	DeleteAll()
	IsUserPrivate(id primitive.ObjectID) (bool, error)
	GetAllPublicUsers() ([]*User, error)
	SearchPublicUsersByUsername(criteria string) ([]*User, error)
	SearchPublicUsersByFirstName(criteria string) ([]*User, error)
	SearchPublicUsersByLastName(criteria string) ([]*User, error)
}
