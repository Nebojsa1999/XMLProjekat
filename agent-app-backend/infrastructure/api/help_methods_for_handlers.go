package api

import (
	"encoding/json"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"io"
	"mime"
	"net/http"
)

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

func decodeCompanyFromBody(reader io.Reader) (*domain.Company, error) {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	var companyInRequestBody domain.Company
	if err := decoder.Decode(&companyInRequestBody); err != nil {
		return nil, err
	}

	return &companyInRequestBody, nil
}

func decodeCompanyRegistrationRequestFromBody(reader io.Reader) (*domain.CompanyRegistrationRequest, error) {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	var companyRegistrationRequestInRequestBody domain.CompanyRegistrationRequest
	if err := decoder.Decode(&companyRegistrationRequestInRequestBody); err != nil {
		return nil, err
	}

	return &companyRegistrationRequestInRequestBody, nil
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
