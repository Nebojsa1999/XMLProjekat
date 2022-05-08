package api

import (
	"context"
	"encoding/json"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/domain"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/infrastructure/services"
	postingPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	userPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net/http"
)

type PublicPostHandler struct {
	userClientAddress    string
	postingClientAddress string
}

func NewPublicPostHandler(userClientAddress, postingClientAddress string) Handler {
	return &PublicPostHandler{
		userClientAddress:    userClientAddress,
		postingClientAddress: postingClientAddress,
	}
}

func (handler *PublicPostHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/user/{id}/public-posts", handler.GetPublicPostsOfOneUser)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = mux.HandlePath("GET", "/post/public", handler.GetAllPublicPosts)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (handler *PublicPostHandler) GetPublicPostsOfOneUser(writer http.ResponseWriter, request *http.Request, pathParams map[string]string) {
	userId := pathParams["id"]
	if userId == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userStatusRequest := &domain.UserStatusRequest{Id: userId}
	err := handler.isUserPrivate(userStatusRequest)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if userStatusRequest.IsPrivate {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		err := handler.getPosts(userStatusRequest)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		response, err := json.Marshal(userStatusRequest.Posts)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(response)
	}
}

func (handler *PublicPostHandler) isUserPrivate(userStatusRequest *domain.UserStatusRequest) error {
	userClient := services.NewUserClient(handler.userClientAddress)

	isUserPrivateResponse, err := userClient.IsUserPrivate(context.TODO(), &userPb.IsPrivateRequest{Id: userStatusRequest.Id})
	if err != nil {
		return err
	}

	userStatusRequest.IsPrivate = isUserPrivateResponse.Private
	return nil
}

func (handler *PublicPostHandler) getPosts(userStatusRequest *domain.UserStatusRequest) error {
	postingClient := services.NewPostingClient(handler.postingClientAddress)

	postsCollection, err := postingClient.GetAllPostsFromUser(context.TODO(), &postingPb.GetRequest{Id: userStatusRequest.Id})
	if err != nil {
		return err
	}

	userStatusRequest.Posts = postsCollection.Posts
	return nil
}

func (handler *PublicPostHandler) GetAllPublicPosts(writer http.ResponseWriter, request *http.Request, pathParams map[string]string) {
	allPublicPosts := &domain.GetAllPostsRequest{}

	err := handler.getIdsOfAllPublicUsers(allPublicPosts)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if allPublicPosts.UserIds == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else {
		err := handler.getAllPosts(allPublicPosts)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		response, err := json.Marshal(allPublicPosts.Posts)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(response)
	}
}

func (handler *PublicPostHandler) getIdsOfAllPublicUsers(getAllPostsRequest *domain.GetAllPostsRequest) error {
	userClient := services.NewUserClient(handler.userClientAddress)

	getIdsOfAllPublicUsersIds, err := userClient.GetIdsOfAllPublicUsers(context.TODO(), &userPb.GetIdsOfAllPublicUsersRequest{})
	if err != nil {
		return err
	}

	getAllPostsRequest.UserIds = getIdsOfAllPublicUsersIds.Ids
	return nil
}

func (handler *PublicPostHandler) getAllPosts(publicPost *domain.GetAllPostsRequest) error {
	postingClient := services.NewPostingClient(handler.postingClientAddress)

	postsCollection, err := postingClient.GetAllPublicPosts(context.TODO(), &postingPb.GetAllPublicPostsRequest{PostIds: publicPost.UserIds})
	if err != nil {
		return err
	}

	publicPost.Posts = postsCollection.Posts
	return nil
}
