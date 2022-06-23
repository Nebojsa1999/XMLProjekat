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
	userInDatabase.IsPrivate = modifiedUser.IsPrivate
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

func (service *UserService) GenerateJobOffersAPIToken(userId primitive.ObjectID) (*domain.JobOffersAPIToken, error) {
	user, err := service.Get(userId)
	if err != nil {
		return nil, err
	}

	var tokenSigningKey = []byte(os.Getenv("SECRET_FOR_JOB_OFFERS_API_TOKEN"))
	jobOffersAPIToken := jwt.New(jwt.SigningMethodHS256)
	claims := jobOffersAPIToken.Claims.(jwt.MapClaims)

	claims["dislinktUserId"] = user.Id
	claims["exp"] = 0

	jobOffersAPITokenString, err := jobOffersAPIToken.SignedString(tokenSigningKey)
	if err != nil {
		err = fmt.Errorf("Error occurred during signing of token: %s", err.Error())
		return nil, err
	}

	return &domain.JobOffersAPIToken{Token: jobOffersAPITokenString}, nil
}
