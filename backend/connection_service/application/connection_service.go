package application

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConnectionService struct {
	store domain.ConnectionStore
}

func NewConnectionService(store domain.ConnectionStore) *ConnectionService {
	return &ConnectionService{
		store: store,
	}
}

func (service *ConnectionService) Get(userId primitive.ObjectID) ([]*domain.Connection, error) {
	return service.store.Get(userId)
}

func (service *ConnectionService) CreateConnection(connection *domain.Connection) (*domain.Connection, error) {
	return service.store.CreateConnection(connection)
}

func (service *ConnectionService) UpdateConnection(id string) (*domain.Connection, error) {
	return service.store.UpdateConnection(id)
}

func (service *ConnectionService) DeleteConnection(id primitive.ObjectID) error {
	return service.store.DeleteConnection(id)
}
