package application

import (
	"fmt"
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

func (service *ConnectionService) Get(id primitive.ObjectID) (*domain.Connection, error) {
	return service.store.Get(id)
}

func (service *ConnectionService) GetAll() ([]*domain.Connection, error) {
	return service.store.GetAll()
}

func (service *ConnectionService) GetByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	return service.store.GetByUserId(userId)
}

func (service *ConnectionService) GetFollowingByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	return service.store.GetFollowingByUserId(userId)
}

func (service *ConnectionService) GetFollowersByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	return service.store.GetFollowersByUserId(userId)
}

func (service *ConnectionService) Create(connection *domain.Connection) (*domain.Connection, error) {
	existingConnection, _ := service.store.Get(connection.Id)
	connection.Id = primitive.NewObjectID()
	if existingConnection != nil {
		return nil, fmt.Errorf("connection with the same id already exists")
	}

	if connection.IssuerId == connection.SubjectId {
		return nil, fmt.Errorf("user cannot follow themselves")
	}

	allConnections, _ := service.store.GetAll()
	for _, c := range allConnections {
		if c.IssuerId == connection.IssuerId && c.SubjectId == connection.SubjectId {
			return nil, fmt.Errorf("same connection already exists")
		}
	}

	return service.store.Create(connection)
}

func (service *ConnectionService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *ConnectionService) Update(id primitive.ObjectID) (*domain.Connection, error) {
	connectionInDatabase, _ := service.store.Get(id)
	if connectionInDatabase == nil {
		return nil, fmt.Errorf("connection with given id does not exist")
	}

	connectionInDatabase.IsApproved = !connectionInDatabase.IsApproved

	return service.store.Update(connectionInDatabase)
}

func (service *ConnectionService) UpdatePrivacy(modifiedPrivacy *domain.ProfilePrivacy) (*domain.ProfilePrivacy, error) {
	privacyInDatabase, _ := service.store.GetPrivacy(modifiedPrivacy.Id)
	if privacyInDatabase == nil {
		return nil, fmt.Errorf("profile privacy with given id does not exist")
	}

	privacyInDatabase.IsPrivate = !privacyInDatabase.IsPrivate

	return service.store.UpdatePrivacy(modifiedPrivacy)
}

func (service *ConnectionService) CreateProfilePrivacy(privacy *domain.ProfilePrivacy) (*domain.ProfilePrivacy, error) {
	return service.store.CreateProfilePrivacy(privacy)
}

func (service *ConnectionService) DeleteProfilePrivacy(id primitive.ObjectID) error {
	return service.store.DeleteProfilePrivacy(id)
}
