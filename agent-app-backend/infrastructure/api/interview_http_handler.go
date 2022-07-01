package api

import (
	"net/http"
	"strings"

	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/application"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InterviewHandler struct {
	service *application.InterviewService
}

func NewInterviewHandler(service *application.InterviewService) *InterviewHandler {
	return &InterviewHandler{
		service: service,
	}
}

func (handler *InterviewHandler) Get(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	interview, err := handler.service.Get(objectId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(writer, interview)
}

func (handler *InterviewHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	interviews, err := handler.service.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(writer, interviews)
}

func (handler *InterviewHandler) CreateNewInterview(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	if !isContentTypeValid(writer, request) {
		return
	}

	newInterview, err := decodeInterviewFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := handler.service.CreateNewInterview(newInterview)
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

func (handler *InterviewHandler) Update(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	if !isContentTypeValid(writer, request) {
		return
	}

	modifiedInterview, err := decodeInterviewFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	} else if objectId != modifiedInterview.Id {
		http.Error(writer, "Id in path and id of modified interview do not match!", http.StatusBadRequest)
		return
	}

	message, updatedInterview, err := handler.service.Update(modifiedInterview)
	if err != nil {
		http.Error(writer, message, http.StatusInternalServerError)
		return
	} else if updatedInterview == nil {
		http.Error(writer, message, http.StatusBadRequest)
		return
	}

	renderJSON(writer, updatedInterview)
}
