package domain

type ConnectionStore interface {
	Get(userId string) ([]*Connection, error)
	CreateConnection(connection *Connection) (*Connection, error)
	CreateProfilePrivacy(privacy *ProfilePrivacy) (*ProfilePrivacy, error)
	UpdateConnection(id string) (*Connection, error)
	DeleteConnection(id string) error
	DeleteAll() error
}
