package api

import (
	"context"
	"encoding/json"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/domain"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/infrastructure/services"
	userPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net/http"
)

type RegisterHandler struct {
	userClientAddress string
}

func NewRegisterHandler(userClientAddress string) Handler {
	return &RegisterHandler{
		userClientAddress: userClientAddress,
	}
}

func (handler *RegisterHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/user/reg-handler", handler.Register)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (handler *RegisterHandler) Register(writer http.ResponseWriter, request *http.Request, pathParams map[string]string) {
	var userRequest userPb.User

	err := json.NewDecoder(request.Body).Decode(&userRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	registrationRequest := &domain.UserRegistrationRequest{
		User: userPb.User{
			Id:                userRequest.Id,
			Username:          userRequest.Username,
			Password:          userRequest.Password,
			IsPrivate:         userRequest.IsPrivate,
			FirstName:         userRequest.FirstName,
			LastName:          userRequest.LastName,
			Email:             userRequest.Email,
			Gender:            userRequest.Gender,
			DateOfBirth:       userRequest.DateOfBirth,
			Biography:         userRequest.Biography,
			WorkExperience:    userRequest.WorkExperience,
			Education:         userRequest.Education,
			Skills:            userRequest.Skills,
			Interests:         userRequest.Interests,
			JobOffersAPIToken: userRequest.JobOffersAPIToken,
		},
	}

	err = handler.RegisterInUserService(registrationRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	} else {
		response, err := json.Marshal(registrationRequest.User)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(response)
	}
}

func (handler *RegisterHandler) RegisterInUserService(userRegistrationRequest *domain.UserRegistrationRequest) error {
	userClient := services.NewUserClient(handler.userClientAddress)

	_, err := userClient.RegisterANewUser(context.TODO(), &userPb.RegisterRequest{User: &userRegistrationRequest.User})
	if err != nil {
		return err
	}

	return nil
}
