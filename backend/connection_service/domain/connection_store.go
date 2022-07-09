package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConnectionStore interface {
	Get(id primitive.ObjectID) (*Connection, error)
	GetAll() ([]*Connection, error)
	GetConnectionOfFollowingType(id primitive.ObjectID) (*Connection, error)
	GetAllConnectionsOfFollowingType() ([]*Connection, error)
	GetConnectionOfBlockingType(id primitive.ObjectID) (*Connection, error)
	GetAllConnectionsOfBlockingType() ([]*Connection, error)
	GetByUserId(userId primitive.ObjectID) ([]*Connection, error)
	GetConnectionsOfFollowingTypeByUserId(userId primitive.ObjectID) ([]*Connection, error)
	GetConnectionsOfBlockingTypeByUserId(userId primitive.ObjectID) ([]*Connection, error)
	GetByTypeAndIssuerIdAndSubjectId(connectionUpdateDTO *ConnectionUpdateDTO) (*Connection, error)
	GetFollowingByUserId(userId primitive.ObjectID) ([]*Connection, error)
	GetFollowersByUserId(userId primitive.ObjectID) ([]*Connection, error)
	GetConnectionsInWhichTheGivenUserIsBlocker(userId primitive.ObjectID) ([]*Connection, error)
	GetConnectionsInWhichTheGivenUserIsBlockedOne(userId primitive.ObjectID) ([]*Connection, error)
	Create(connection *Connection) (*Connection, error)
	Update(updatedConnection *Connection) (*Connection, error)
	Delete(id primitive.ObjectID) error
	DeleteAll() error
}
