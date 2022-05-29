package api

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/application"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

type UserHandler struct {
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(writer http.ResponseWriter, request *http.Request) {
	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := handler.service.Get(objectId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(writer, user)
}

func (handler *UserHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	users, err := handler.service.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(writer, users)
}

func (handler *UserHandler) RegisterANewUser(writer http.ResponseWriter, request *http.Request) {
	if !isContentTypeValid(writer, request) {
		return
	}

	newUser, err := decodeUserFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := handler.service.RegisterANewUser(newUser)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	} else if !strings.HasPrefix(message, "Success:") {
		http.Error(writer, message, http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	renderJSON(writer, message)
}

func (handler *UserHandler) Login(writer http.ResponseWriter, request *http.Request) {
	if !isContentTypeValid(writer, request) {
		return
	}

	userCredentials, err := decodeCredentialsFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	jwtToken, message, err := handler.service.Login(userCredentials)
	if err != nil {
		http.Error(writer, message, http.StatusInternalServerError)
		return
	} else if jwtToken == nil {
		http.Error(writer, message, http.StatusBadRequest)
		return
	}

	renderJSON(writer, jwtToken)
}

func (handler *UserHandler) Update(writer http.ResponseWriter, request *http.Request) {
	if !isContentTypeValid(writer, request) {
		return
	}

	modifiedUser, err := decodeUserFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	} else if objectId != modifiedUser.Id {
		http.Error(writer, "Id in path and id of modified user do not match!", http.StatusBadRequest)
		return
	}

	message, updatedUser, err := handler.service.Update(modifiedUser)
	if err != nil {
		http.Error(writer, message, http.StatusInternalServerError)
		return
	} else if updatedUser == nil {
		http.Error(writer, message, http.StatusBadRequest)
		return
	}

	renderJSON(writer, updatedUser)
}
