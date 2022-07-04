package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConnectionStore interface {
	Get(id primitive.ObjectID) (*Connection, error)
	GetAll() ([]*Connection, error)
	GetByUserId(userId primitive.ObjectID) ([]*Connection, error)
	GetFollowingByUserId(userId primitive.ObjectID) ([]*Connection, error)
	GetFollowersByUserId(userId primitive.ObjectID) ([]*Connection, error)
	Create(connection *Connection) (*Connection, error)
	CreatePrivacy(privacy *ProfilePrivacy) (*ProfilePrivacy, error)
	Delete(id string) error
	DeleteAll() error
	Update(updatedConnection *Connection) (*Connection, error)
	CreateProfilePrivacy(privacy *ProfilePrivacy) (*ProfilePrivacy, error)
	DeleteProfilePrivacy(id primitive.ObjectID) error
	UpdatePrivacy(id primitive.ObjectID) error
}
