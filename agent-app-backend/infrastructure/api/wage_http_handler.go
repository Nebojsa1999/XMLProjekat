package api

import (
	"net/http"
	"strings"

	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/application"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WageHandler struct {
	service *application.WageService
}

func NewWageHandler(service *application.WageService) *WageHandler {
	return &WageHandler{
		service: service,
	}
}

func (handler *WageHandler) Get(writer http.ResponseWriter, request *http.Request) {
	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	wage, err := handler.service.Get(objectId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(writer, wage)
}

func (handler *WageHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	wages, err := handler.service.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(writer, wages)
}

func (handler *WageHandler) CreateNewWage(writer http.ResponseWriter, request *http.Request) {
	if !isContentTypeValid(writer, request) {
		return
	}

	newWage, err := decodeWageFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := handler.service.CreateNewWage(newWage)
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

func (handler *WageHandler) Update(writer http.ResponseWriter, request *http.Request) {
	if !isContentTypeValid(writer, request) {
		return
	}

	modifiedWage, err := decodeWageFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	} else if objectId != modifiedWage.Id {
		http.Error(writer, "Id in path and id of modified wage do not match!", http.StatusBadRequest)
		return
	}

	message, updatedWage, err := handler.service.Update(modifiedWage)
	if err != nil {
		http.Error(writer, message, http.StatusInternalServerError)
		return
	} else if updatedWage == nil {
		http.Error(writer, message, http.StatusBadRequest)
		return
	}

	renderJSON(writer, updatedWage)
}
