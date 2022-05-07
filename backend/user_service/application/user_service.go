package application

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) RegisterANewUser(user *domain.User) (string, error) {
	existingUser, _ := service.store.Get(user.Id)
	user.Id = primitive.NewObjectID()
	if existingUser != nil {
		return "User with the same id already exists.", nil
	}

	existingUser, _ = service.store.GetByUsername(user.Username)
	if existingUser != nil {
		return "Username is already taken, please choose another one.", nil
	}

	existingUser, _ = service.store.GetByEmail(user.Email)
	if existingUser != nil {
		return "This email is already linked with another user.", nil
	}

	return service.store.RegisterANewUser(user)
}
