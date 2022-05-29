package api

import (
	"encoding/json"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/application"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"mime"
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

func renderJSON(writer http.ResponseWriter, data interface{}) {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(marshalledData)
}

func decodeUserFromBody(reader io.Reader) (*domain.User, error) {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	var userInRequestBody domain.User
	if err := decoder.Decode(&userInRequestBody); err != nil {
		return nil, err
	}

	return &userInRequestBody, nil
}

func decodeCredentialsFromBody(reader io.Reader) (*domain.Credentials, error) {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	var credentialsInRequestBody domain.Credentials
	if err := decoder.Decode(&credentialsInRequestBody); err != nil {
		return nil, err
	}

	return &credentialsInRequestBody, nil
}

func isContentTypeValid(writer http.ResponseWriter, request *http.Request) bool {
	validity := true

	contentType := request.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		validity = false
	}
	if mediaType != "application/json" {
		http.Error(writer, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		validity = false
	}

	return validity
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
