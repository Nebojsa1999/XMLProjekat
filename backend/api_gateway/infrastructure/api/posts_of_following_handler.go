package api

import (
	"context"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/domain"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/infrastructure/services"
	connectionPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/connection_service"
	postingPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net/http"
)

type PostsOfFollowingHandler struct {
	connectionClientAddress string
	postingClientAddress    string
}

func NewPostsOfFollowingHandler(connectionClientAddress, postingClientAddress string) Handler {
	return &PostsOfFollowingHandler{
		connectionClientAddress: connectionClientAddress,
		postingClientAddress:    postingClientAddress,
	}
}

func (handler *PostsOfFollowingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/user/{id}/posts-of-following", handler.GetPostsOfFollowingUsers)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (handler *PostsOfFollowingHandler) GetPostsOfFollowingUsers(writer http.ResponseWriter, request *http.Request, pathParams map[string]string) {
	if !isContentTypeValid(writer, request) {
		return
	}

	userId := pathParams["id"]
	if userId == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	getPostsOfFollowingUsersRequest := &domain.GetPostsOfFollowingUsersRequest{UserId: userId}

	err := handler.getFollowingUsersIds(getPostsOfFollowingUsersRequest)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if getPostsOfFollowingUsersRequest.FollowingIds == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		err := handler.getAllPostsOfFollowingUsers(getPostsOfFollowingUsersRequest)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		renderJSON(writer, getPostsOfFollowingUsersRequest.PostsOfFollowingUsers)
	}
}

func (handler *PostsOfFollowingHandler) getFollowingUsersIds(request *domain.GetPostsOfFollowingUsersRequest) error {
	connectionClient := services.NewConnectionClient(handler.connectionClientAddress)

	followingUsersIdsResponse, err := connectionClient.GetFollowingUsersIds(context.TODO(),
		&connectionPb.GetByUserIdRequest{UserId: request.UserId})
	if err != nil {
		return err
	}

	request.FollowingIds = followingUsersIdsResponse.Ids

	return nil
}

func (handler *PostsOfFollowingHandler) getAllPostsOfFollowingUsers(request *domain.GetPostsOfFollowingUsersRequest) error {
	postingClient := services.NewPostingClient(handler.postingClientAddress)

	var allPostsCollection []*postingPb.Post

	for _, id := range request.FollowingIds {
		allPostsOfOneUserResponse, err := postingClient.GetAllPostsFromUser(context.TODO(),
			&postingPb.GetRequest{Id: id})
		if err != nil {
			return err
		}

		if allPostsOfOneUserResponse.Posts != nil {
			allPostsCollection = append(allPostsCollection, allPostsOfOneUserResponse.Posts)
		}
	}

	request.PostsOfFollowingUsers = allPostsCollection

	return nil
}
