package application

import (
	"fmt"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/domain"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
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

func (service *UserService) Login(credentials *domain.Credentials) (*domain.JWTToken, string, error) {
	existingUser, err := service.store.GetByUsername(credentials.Username)
	if err != nil {
		return nil, "There is no user with that username.", err
	}

	if credentials.Password != existingUser.Password {
		return nil, "Password is incorrect.", nil
	}

	jwtToken, err := service.GenerateJWTToken(credentials.Username)
	if err != nil {
		return nil, "Error occurred during generating JWT token, login failed!", err
	}

	return jwtToken, "Success: user has been logged in.", nil
}

func (service *UserService) GenerateJWTToken(username string) (*domain.JWTToken, error) {
	var tokenSigningKey = []byte(os.Getenv("SECRET_FOR_JWT"))
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	jwtTokenString, err := jwtToken.SignedString(tokenSigningKey)
	if err != nil {
		err = fmt.Errorf("Error occured during signing of token: %s", err)
		return nil, err
	}

	return &domain.JWTToken{Token: jwtTokenString}, nil
}
