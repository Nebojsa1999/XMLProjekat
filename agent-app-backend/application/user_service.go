package application

import (
	"fmt"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	jwt "github.com/dgrijalwa/jwt-go"
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

func (service *UserService) Login(credentials *domain.Credentials) (*domain.AgentAppToken, string, error) {
	existingUser, err := service.store.GetByUsername(credentials.Username)
	if err != nil {
		return nil, "Error occurred during retrieval of possible user with same username from database.", err
	}

	if existingUser == nil {
		return nil, "There is no user with that username.", nil
	}

	if credentials.Password != existingUser.Password {
		return nil, "Password is incorrect.", nil
	}

	jwtToken, err := service.GenerateAgentAppJWTToken(existingUser)
	if err != nil {
		return nil, "Error occurred during generating JWT token, login failed!", err
	}

	return jwtToken, "Success: user has been logged in.", nil
}

func (service *UserService) GenerateAgentAppJWTToken(user *domain.User) (*domain.AgentAppToken, error) {
	var tokenSigningKey = []byte(os.Getenv("SECRET_FOR_AGENT_APP_JWT"))
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = user.Id
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	jwtTokenString, err := jwtToken.SignedString(tokenSigningKey)
	if err != nil {
		err = fmt.Errorf("error occurred during signing of token: %s", err.Error())
		return nil, err
	}

	return &domain.AgentAppToken{Token: jwtTokenString}, nil
}

func (service *UserService) Update(modifiedUser *domain.User) (string, *domain.User, error) {
	userInDatabase, _ := service.store.Get(modifiedUser.Id)
	if userInDatabase == nil {
		return "User with given id does not exist.", nil, nil
	}

	userInDatabaseWithSameUsername, _ := service.store.GetByUsername(modifiedUser.Username)
	if userInDatabaseWithSameUsername != nil {
		if userInDatabaseWithSameUsername.Id != userInDatabase.Id {
			return "Given username is already taken by another user.", nil, nil
		}
	}

	userInDatabaseWithSameEmail, _ := service.store.GetByEmail(modifiedUser.Email)
	if userInDatabaseWithSameEmail != nil {
		if userInDatabaseWithSameEmail.Id != userInDatabase.Id {
			return "Given email is already linked to another user.", nil, nil
		}
	}

	userInDatabase.Username = modifiedUser.Username
	userInDatabase.Password = modifiedUser.Password
	userInDatabase.FirstName = modifiedUser.FirstName
	userInDatabase.LastName = modifiedUser.LastName
	userInDatabase.Email = modifiedUser.Email
	userInDatabase.Phone = modifiedUser.Phone
	userInDatabase.Gender = modifiedUser.Gender
	userInDatabase.DateOfBirth = modifiedUser.DateOfBirth
	userInDatabase.Biography = modifiedUser.Biography
	userInDatabase.WorkExperience = modifiedUser.WorkExperience
	userInDatabase.Education = modifiedUser.Education
	userInDatabase.Skills = modifiedUser.Skills
	userInDatabase.Interests = modifiedUser.Interests

	return service.store.Update(userInDatabase)
}
