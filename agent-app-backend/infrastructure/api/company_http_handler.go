package api

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/application"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

type CompanyHandler struct {
	service *application.CompanyService
}

func NewCompanyHandler(service *application.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		service: service,
	}
}

func (handler *CompanyHandler) Get(writer http.ResponseWriter, request *http.Request) {
	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	company, err := handler.service.Get(objectId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(writer, company)
}

func (handler *CompanyHandler) GetByName(writer http.ResponseWriter, request *http.Request) {
	name := mux.Vars(request)["name"]
	if name == "" {
		http.Error(writer, "Company name is empty.", http.StatusBadRequest)
		return
	}

	company, err := handler.service.GetByName(name)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(writer, company)
}

func (handler *CompanyHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	companies, err := handler.service.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(writer, companies)
}

func (handler *CompanyHandler) RegisterANewCompany(writer http.ResponseWriter, request *http.Request) {
	if !isContentTypeValid(writer, request) {
		return
	}

	newCompany, err := decodeCompanyFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := handler.service.RegisterANewCompany(newCompany)
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

func (handler *CompanyHandler) Update(writer http.ResponseWriter, request *http.Request) {
	if !isContentTypeValid(writer, request) {
		return
	}

	modifiedCompany, err := decodeCompanyFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	} else if objectId != modifiedCompany.Id {
		http.Error(writer, "Id in path and id of modified company do not match!", http.StatusBadRequest)
		return
	}

	message, updatedCompany, err := handler.service.Update(modifiedCompany)
	if err != nil {
		http.Error(writer, message, http.StatusInternalServerError)
		return
	} else if updatedCompany == nil {
		http.Error(writer, message, http.StatusBadRequest)
		return
	}

	renderJSON(writer, updatedCompany)
}
