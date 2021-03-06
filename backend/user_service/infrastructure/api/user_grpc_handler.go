package api

import (
	"context"
	"fmt"
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	user, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}

	userPb := mapDomainUserToPbUser(user)
	response := &pb.GetResponse{
		User: userPb,
	}

	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapDomainUserToPbUser(user)
		response.Users = append(response.Users, current)
	}

	return response, nil
}

func (handler *UserHandler) GetAllPublicUsers(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	users, err := handler.service.GetAllPublicUsers()
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapDomainUserToPbUser(user)
		response.Users = append(response.Users, current)
	}

	return response, nil
}

func (handler *UserHandler) RegisterANewUser(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	newUser := mapPbUserToDomainUser(request.User)

	message, err := handler.service.RegisterANewUser(newUser)
	response := &pb.RegisterResponse{
		Message: message,
	}

	return response, err
}

func (handler *UserHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	userCredentials := mapPbCredentialsToDomainCredentials(request.Credentials)

	jwtToken, message, err := handler.service.Login(userCredentials)
	response := &pb.LoginResponse{
		Token:   "",
		Message: message,
	}

	if jwtToken == nil {
		err = fmt.Errorf(message)
	}

	if jwtToken != nil {
		response.Token = jwtToken.Token
	}

	return response, err
}

func (handler *UserHandler) IsUserPrivate(ctx context.Context, request *pb.IsPrivateRequest) (*pb.IsPrivateResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	isUserPrivate, err := handler.service.IsUserPrivate(id)
	if err != nil {
		return nil, err
	}

	response := &pb.IsPrivateResponse{
		Private: isUserPrivate,
	}

	return response, nil
}

func (handler *UserHandler) GetIdsOfAllPublicUsers(ctx context.Context, request *pb.GetIdsOfAllPublicUsersRequest) (*pb.GetIdsOfAllPublicUsersResponse, error) {
	idsOfAllPublicUsers, err := handler.service.GetIdsOfAllPublicUsers()
	if err != nil {
		return nil, err
	}

	response := &pb.GetIdsOfAllPublicUsersResponse{
		Ids: []string{},
	}
	for _, id := range idsOfAllPublicUsers {
		currentId := id.Hex()
		response.Ids = append(response.Ids, currentId)
	}

	return response, nil
}

func (handler *UserHandler) SearchPublicUsers(ctx context.Context, request *pb.SearchPublicUsersRequest) (*pb.SearchPublicUsersResponse, error) {
	filteredUsers, err := handler.service.SearchPublicUsers(request.Criteria)
	if err != nil {
		return nil, err
	}

	response := &pb.SearchPublicUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range filteredUsers {
		currentUser := mapDomainUserToPbUser(user)
		response.Users = append(response.Users, currentUser)
	}

	return response, nil
}

func (handler *UserHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	if request.Id != request.ModifiedUser.Id {
		return &pb.UpdateResponse{
			Message:     "Id in path and id of modified user do not match!",
			UpdatedUser: nil,
		}, nil
	}

	modifiedUser := mapPbUserToDomainUser(request.ModifiedUser)

	message, updatedUser, err := handler.service.Update(modifiedUser)
	response := &pb.UpdateResponse{
		Message:     message,
		UpdatedUser: nil,
	}
	if updatedUser == nil {
		return response, fmt.Errorf(message)
	}
	if updatedUser != nil {
		response.UpdatedUser = mapDomainUserToPbUser(updatedUser)
	}

	return response, err
}

func (handler *UserHandler) GenerateJobOffersAPIToken(ctx context.Context, request *pb.GenerateJobOffersAPITokenRequest) (*pb.GenerateJobOffersAPITokenResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}

	message, jobOffersAPIToken, err := handler.service.GenerateJobOffersAPIToken(userId)
	response := &pb.GenerateJobOffersAPITokenResponse{
		Token:   "",
		Message: message,
	}

	if jobOffersAPIToken != nil {
		response.Token = jobOffersAPIToken.Token
	}

	return response, err
}
