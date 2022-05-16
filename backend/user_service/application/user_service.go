package application

import (
	"fmt"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/domain"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
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
		return nil, "Error occurred during retrieval of possible user with same username from database.", err
	}

	if existingUser == nil {
		return nil, "There is no user with that username.", nil
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
		err = fmt.Errorf("Error occurred during signing of token: %s", err.Error())
		return nil, err
	}

	return &domain.JWTToken{Token: jwtTokenString}, nil
}

func (service *UserService) IsUserPrivate(id primitive.ObjectID) (bool, error) {
	return service.store.IsUserPrivate(id)
}

func (service *UserService) GetIdsOfAllPublicUsers() ([]primitive.ObjectID, error) {
	allPublicUsers, err := service.store.GetAllPublicUsers()
	if err != nil {
		return nil, err
	}

	var idsOfAllPublicUsers []primitive.ObjectID
	for _, user := range allPublicUsers {
		idsOfAllPublicUsers = append(idsOfAllPublicUsers, user.Id)
	}

	return idsOfAllPublicUsers, err
}

func (service *UserService) SearchPublicUsers(criteria string) ([]*domain.User, error) {
	var publicUsersMatchingCriteria []*domain.User

	criteria = strings.TrimSpace(criteria)
	splitSearch := strings.Split(criteria, " ")

	for _, splitSearchPart := range splitSearch {
		publicUsersWithMatchingUsername, err := service.store.SearchPublicUsersByUsername(splitSearchPart)
		if err != nil {
			return nil, err
		}

		for _, userOneSlice := range publicUsersWithMatchingUsername {
			publicUsersMatchingCriteria = appendIfMissing(publicUsersMatchingCriteria, userOneSlice)
		}

		publicUsersWithMatchingFirstName, err := service.store.SearchPublicUsersByFirstName(splitSearchPart)
		if err != nil {
			return nil, err
		}

		for _, userOneSlice := range publicUsersWithMatchingFirstName {
			publicUsersMatchingCriteria = appendIfMissing(publicUsersMatchingCriteria, userOneSlice)
		}

		publicUsersWithMatchingLastName, err := service.store.SearchPublicUsersByLastName(splitSearchPart)
		if err != nil {
			return nil, err
		}

		for _, userOneSlice := range publicUsersWithMatchingLastName {
			publicUsersMatchingCriteria = appendIfMissing(publicUsersMatchingCriteria, userOneSlice)
		}
	}

	return publicUsersMatchingCriteria, nil
}

func appendIfMissing(slice []*domain.User, i *domain.User) []*domain.User {
	for _, element := range slice {
		if element.Id == i.Id {
			return slice
		}
	}

	return append(slice, i)
}

func (service *UserService) UpdatePersonalInformation(user *domain.User) (string, error) {
	userInDatabase, _ := service.store.Get(user.Id)
	if userInDatabase == nil {
		return "User with given id does not exits.", nil
	}

	possibleUserWithSameUsername, _ := service.store.GetByUsername(user.Username)
	if possibleUserWithSameUsername != nil {
		if possibleUserWithSameUsername.Id != userInDatabase.Id {
			return "Given username is already taken by another user.", nil
		}
	}

	userInDatabase.FirstName = user.FirstName
	userInDatabase.Email = user.Email
	userInDatabase.Gender = user.Gender
	userInDatabase.DateOfBirth = user.DateOfBirth
	userInDatabase.Username = user.Username
	userInDatabase.Biography = user.Biography

	return service.store.UpdatePersonalInformation(userInDatabase)
}