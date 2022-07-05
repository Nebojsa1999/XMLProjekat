package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConnectionStore interface {
	Get(id primitive.ObjectID) (*Connection, error)
	GetAll() ([]*Connection, error)
	GetByIssuerIdAndSubjectId(issuerId, subjectId primitive.ObjectID) (*Connection, error)
	GetByUserId(userId primitive.ObjectID) ([]*Connection, error)
	GetFollowingByUserId(userId primitive.ObjectID) ([]*Connection, error)
	GetFollowersByUserId(userId primitive.ObjectID) ([]*Connection, error)
	Create(connection *Connection) (*Connection, error)
	Update(updatedConnection *Connection) (*Connection, error)
	Delete(id primitive.ObjectID) error
	DeleteAll() error
	GetPrivacy(id primitive.ObjectID) (*ProfilePrivacy, error)
	CreateProfilePrivacy(privacy *ProfilePrivacy) (*ProfilePrivacy, error)
	DeleteProfilePrivacy(id primitive.ObjectID) error
	UpdatePrivacy(updatedPrivacy *ProfilePrivacy) (*ProfilePrivacy, error)
}
