package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/domain"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/infrastructure/services"
	connectionPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/connection_service"
	postingPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type ConnectionHandler struct {
	postingClientAddress    string
	connectionClientAddress string
}

func NewConnectionHandler(postingClientAddress, connectionClientAddress string) Handler {
	return &ConnectionHandler{
		postingClientAddress:    postingClientAddress,
		connectionClientAddress: connectionClientAddress,
	}
}

func (handler *ConnectionHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/connection/{id}/post", handler.GetPostsByConnectedUsers)
	if err != nil {
		fmt.Println("Panika")
		panic(err)
	}
}

func (handler *ConnectionHandler) GetPostsByConnectedUsers(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	publicPost := &domain.GetAllPostsRequest{}

	err := handler.GetConnectedUsersIds(id, publicPost)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if publicPost.UserIds == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		err1 := handler.getAllPosts(publicPost)
		if err1 != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		response, err := json.Marshal(publicPost.Posts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (handler *ConnectionHandler) GetConnectedUsersIds(id string, publicPost *domain.GetAllPostsRequest) error {
	connectionClient := services.NewConnectionClient(handler.connectionClientAddress)
	response, err := connectionClient.Get(context.TODO(), &connectionPb.GetRequest{UserId: id})
	if err != nil {
		return err
	}
	for _, connection := range response.Connections {
		publicPost.UserIds = append(publicPost.UserIds, connection.UserBId)
	}
	return nil
}

func (handler *ConnectionHandler) getAllPosts(publicPost *domain.GetAllPostsRequest) error {
	postClient := services.NewPostingClient(handler.postingClientAddress)
	postCollection, err := postClient.GetAllPosts(context.TODO(), &postingPb.GetAllPublicPostsRequest{PostIds: publicPost.UserIds})
	if err != nil {
		return err
	}

	publicPost.Posts = postCollection.Posts
	return nil
}
