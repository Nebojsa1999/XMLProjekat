package application

import (
	"fmt"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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

func (service *ConnectionService) GetConnectionOfFollowingType(id primitive.ObjectID) (*domain.Connection, error) {
	return service.store.GetConnectionOfFollowingType(id)
}

func (service *ConnectionService) GetAllConnectionsOfFollowingType() ([]*domain.Connection, error) {
	return service.store.GetAllConnectionsOfFollowingType()
}

func (service *ConnectionService) GetConnectionOfBlockingType(id primitive.ObjectID) (*domain.Connection, error) {
	return service.store.GetConnectionOfBlockingType(id)
}

func (service *ConnectionService) GetAllConnectionsOfBlockingType() ([]*domain.Connection, error) {
	return service.store.GetAllConnectionsOfBlockingType()
}

func (service *ConnectionService) GetByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	return service.store.GetByUserId(userId)
}

func (service *ConnectionService) GetConnectionsOfFollowingTypeByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	return service.store.GetConnectionsOfFollowingTypeByUserId(userId)
}

func (service *ConnectionService) GetConnectionsOfBlockingTypeByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	return service.store.GetConnectionsOfBlockingTypeByUserId(userId)
}

func (service *ConnectionService) GetFollowingByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	return service.store.GetFollowingByUserId(userId)
}

func (service *ConnectionService) GetFollowersByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	return service.store.GetFollowersByUserId(userId)
}

func (service *ConnectionService) GetConnectionsInWhichTheGivenUserIsBlocker(userId primitive.ObjectID) ([]*domain.Connection, error) {
	return service.store.GetConnectionsInWhichTheGivenUserIsBlocker(userId)
}

func (service *ConnectionService) GetConnectionsInWhichTheGivenUserIsBlockedOne(userId primitive.ObjectID) ([]*domain.Connection, error) {
	return service.store.GetConnectionsInWhichTheGivenUserIsBlockedOne(userId)
}

func (service *ConnectionService) GetFollowingUsersIds(userId primitive.ObjectID) ([]primitive.ObjectID, error) {
	followingUsers, err := service.store.GetFollowingByUserId(userId)
	if err != nil {
		return nil, err
	}

	var followingUsersIds []primitive.ObjectID
	for _, u := range followingUsers {
		followingUsersIds = append(followingUsersIds, u.Id)
	}

	return followingUsersIds, nil
}

func (service *ConnectionService) Create(connection *domain.Connection) (*domain.Connection, error) {
	existingConnection, _ := service.store.Get(connection.Id)
	connection.Id = primitive.NewObjectID()
	if existingConnection != nil {
		return nil, fmt.Errorf("connection with the same id already exists")
	}

	if connection.IssuerId == connection.SubjectId {
		return nil, fmt.Errorf("user cannot follow or block themselves")
	}

	allConnections, _ := service.store.GetAll()
	for _, c := range allConnections {
		if c.IssuerId == connection.IssuerId && c.SubjectId == connection.SubjectId &&
			c.Type == connection.Type {
			return nil, fmt.Errorf("same connection already exists")
		}
		if c.IssuerId == connection.IssuerId && c.SubjectId == connection.SubjectId && c.Type == domain.Blocking {
			return nil, fmt.Errorf("user is blocked")
		}
		if c.IssuerId == connection.SubjectId && c.SubjectId == connection.IssuerId && c.Type == domain.Blocking {
			return nil, fmt.Errorf("user is blocked")
		}
	}

	return service.store.Create(connection)
}

func (service *ConnectionService) Update(connectionUpdateDTO *domain.ConnectionUpdateDTO) (*domain.Connection, error) {
	connectionInDatabase, _ :=
		service.store.GetByIssuerIdAndSubjectId(connectionUpdateDTO)
	if connectionInDatabase == nil {
		return nil, fmt.Errorf("connection with given issuer id and subject id does not exist")
	}

	var updatedConnection *domain.Connection

	if connectionInDatabase.Type == domain.Following && connectionUpdateDTO.Type == domain.Blocking {
		err := service.Delete(connectionInDatabase.IssuerId, connectionInDatabase.SubjectId)
		if err != nil {
			return nil, err
		}
		var connectionDTO *domain.ConnectionUpdateDTO = &domain.ConnectionUpdateDTO{
			Type:       domain.Following,
			IssuerId:   connectionInDatabase.SubjectId,
			SubjectId:  connectionInDatabase.IssuerId,
			IsApproved: true,
		}
		possibleReverseConnection, err := service.store.GetByIssuerIdAndSubjectId(connectionDTO)
		if possibleReverseConnection != nil {
			err = service.Delete(connectionInDatabase.SubjectId, connectionInDatabase.IssuerId)
			if err != nil {
				return nil, err
			}
		}

		updatedConnection, err = service.Create(&domain.Connection{
			Id:         primitive.ObjectID{},
			Type:       domain.Blocking,
			IssuerId:   connectionUpdateDTO.IssuerId,
			SubjectId:  connectionUpdateDTO.SubjectId,
			Date:       time.Now(),
			IsApproved: true,
		})
		if err != nil {
			return nil, err
		}

		return updatedConnection, nil
	} else {
		connectionInDatabase.Type = connectionUpdateDTO.Type
		connectionInDatabase.IsApproved = connectionUpdateDTO.IsApproved
	}

	return service.store.Update(connectionInDatabase)
}

func (service *ConnectionService) Delete(issuerId, subjectId primitive.ObjectID) error {

	connectionUpdateDTO :=
		&domain.ConnectionUpdateDTO{IssuerId: issuerId, SubjectId: subjectId}

	connectionInDatabase, _ := service.store.GetByIssuerIdAndSubjectId(connectionUpdateDTO)
	if connectionInDatabase == nil {
		return fmt.Errorf("connection with given issuer id and subject id does not exist")
	}

	return service.store.Delete(connectionInDatabase.Id)
}
