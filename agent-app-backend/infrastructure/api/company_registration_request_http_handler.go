package api

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/application"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

type CompanyRegistrationRequestHandler struct {
	service *application.CompanyRegistrationRequestService
}

func NewCompanyRegistrationRequestHandler(service *application.CompanyRegistrationRequestService) *CompanyRegistrationRequestHandler {
	return &CompanyRegistrationRequestHandler{
		service: service,
	}
}

func (handler *CompanyRegistrationRequestHandler) Get(writer http.ResponseWriter, request *http.Request) {
	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	companyRegistrationRequest, err := handler.service.Get(objectId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(writer, companyRegistrationRequest)
}

func (handler *CompanyRegistrationRequestHandler) GetPendingOnes(writer http.ResponseWriter, request *http.Request) {
	pendingCompanyRegistrationRequests, err := handler.service.GetPendingOnes()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(writer, pendingCompanyRegistrationRequests)
}

func (handler *CompanyRegistrationRequestHandler) GetAcceptedOnes(writer http.ResponseWriter, request *http.Request) {
	acceptedCompanyRegistrationRequests, err := handler.service.GetAcceptedOnes()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(writer, acceptedCompanyRegistrationRequests)
}

func (handler *CompanyRegistrationRequestHandler) GetRejectedOnes(writer http.ResponseWriter, request *http.Request) {
	rejectedCompanyRegistrationRequests, err := handler.service.GetRejectedOnes()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(writer, rejectedCompanyRegistrationRequests)
}

func (handler *CompanyRegistrationRequestHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	companyRegistrationRequests, err := handler.service.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(writer, companyRegistrationRequests)
}

func (handler *CompanyRegistrationRequestHandler) CreateCompanyRegistrationRequest(writer http.ResponseWriter, request *http.Request) {
	if !isContentTypeValid(writer, request) {
		return
	}

	newCompanyRegistrationRequest, err := decodeCompanyRegistrationRequestFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := handler.service.CreateCompanyRegistrationRequest(newCompanyRegistrationRequest)
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

func updateCompanyRegistrationRequestMethodBasis(writer http.ResponseWriter, request *http.Request) *domain.CompanyRegistrationRequest {
	if !isContentTypeValid(writer, request) {
		return nil
	}

	modifiedRequest, err := decodeCompanyRegistrationRequestFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return nil
	}

	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return nil
	} else if objectId != modifiedRequest.Id {
		http.Error(writer, "Id in path and id of modified company registration request do not match!", http.StatusBadRequest)
		return nil
	}

	return modifiedRequest
}

func (handler *CompanyRegistrationRequestHandler) UpdateByOwner(writer http.ResponseWriter, request *http.Request) {
	modifiedRequest := updateCompanyRegistrationRequestMethodBasis(writer, request)
	if modifiedRequest == nil {
		return
	}

	message, updatedRequest, err := handler.service.UpdateByOwner(modifiedRequest)
	if err != nil {
		http.Error(writer, message, http.StatusInternalServerError)
		return
	} else if updatedRequest == nil {
		http.Error(writer, message, http.StatusBadRequest)
		return
	}

	renderJSON(writer, updatedRequest)
}

func (handler *CompanyRegistrationRequestHandler) UpdateByAdministrator(writer http.ResponseWriter, request *http.Request) {
	modifiedRequest := updateCompanyRegistrationRequestMethodBasis(writer, request)
	if modifiedRequest == nil {
		return
	}

	message, updatedRequest, err := handler.service.UpdateByAdministrator(modifiedRequest)
	if err != nil {
		http.Error(writer, message, http.StatusInternalServerError)
		return
	} else if updatedRequest == nil {
		http.Error(writer, message, http.StatusBadRequest)
		return
	}

	renderJSON(writer, updatedRequest)
}
