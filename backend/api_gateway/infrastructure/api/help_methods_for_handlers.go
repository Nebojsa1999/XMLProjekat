package api

import (
	"encoding/json"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/domain"
	"io"
	"mime"
	"net/http"
)

func enableCors(writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
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

func decodePostJobOfferRequestFromBody(reader io.Reader) (*domain.PostJobOfferRequest, error) {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	var postJobOfferRequestInRequestBody domain.PostJobOfferRequest
	if err := decoder.Decode(&postJobOfferRequestInRequestBody); err != nil {
		return nil, err
	}

	return &postJobOfferRequestInRequestBody, nil
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
