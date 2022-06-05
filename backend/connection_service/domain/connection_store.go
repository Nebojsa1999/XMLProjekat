package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConnectionStore interface {
	Get(userId primitive.ObjectID) ([]*Connection, error)
	CreateConnection(connection *Connection) (*Connection, error)
	CreateProfilePrivacy(privacy *ProfilePrivacy) (*ProfilePrivacy, error)
	UpdateConnection(id string) (*Connection, error)
	DeleteConnection(id primitive.ObjectID) error
	DeleteAll() error
}
