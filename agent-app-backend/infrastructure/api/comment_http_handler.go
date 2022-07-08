package api

import (
	"net/http"
	"strings"

	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/application"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentHandler struct {
	service *application.CommentService
}

func NewCommentHandler(service *application.CommentService) *CommentHandler {
	return &CommentHandler{
		service: service,
	}
}

func (handler *CommentHandler) Get(writer http.ResponseWriter, request *http.Request) {
	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	comment, err := handler.service.Get(objectId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(writer, comment)
}

func (handler *CommentHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	comments, err := handler.service.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(writer, comments)
}

func (handler *CommentHandler) CreateNewComment(writer http.ResponseWriter, request *http.Request) {
	if !isContentTypeValid(writer, request) {
		return
	}

	newComment, err := decodeCommentFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := handler.service.CreateNewComment(newComment)
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

func (handler *CommentHandler) Update(writer http.ResponseWriter, request *http.Request) {
	if !isContentTypeValid(writer, request) {
		return
	}

	modifiedComment, err := decodeCommentFromBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := mux.Vars(request)["id"]
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	} else if objectId != modifiedComment.Id {
		http.Error(writer, "Id in path and id of modified comment do not match!", http.StatusBadRequest)
		return
	}

	message, updatedComment, err := handler.service.Update(modifiedComment)
	if err != nil {
		http.Error(writer, message, http.StatusInternalServerError)
		return
	} else if updatedComment == nil {
		http.Error(writer, message, http.StatusBadRequest)
		return
	}

	renderJSON(writer, updatedComment)
}
