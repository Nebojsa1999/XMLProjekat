package api


import (
	"context"
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/domain"
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	pb.UnimplementedPostingServiceServer
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler {cc
		service: service
	}
}

func (handler *PostHandler) GetPostFromUser(ctx context.Context, request *pb.GetPostRequest) (*pb.GetResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	ObjectPostId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.GetPostFromUser(objectId, objectPostId)
	if err != nil {
		return nil, err
	}
	postPb := mapPost(post)
	response := &pb.GetResponse{
		Post: postPb
	}
	return response, nil
}

//todo: GetAllPosts

func (handler *PostHandler) GetAllPostsFromUser(ctx context.Context, request *pb.GetRequest) (*pb.GetAllResponse, error) {
	id := request.PostId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	posts, err := handler.service.GetAllPostsFromUser(objectId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Posts: []*pb.Post{}
	}
	for _, post := range posts {
		current := mapPost(post),
		response.Posts = append(response.Posts, current)
	}
	return response
}

func (handler *PostHandler) CreatePost(ctx context.Context, request *pb.NewPostRequest) (*pb.NewPostResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	newPost, err := handler.service.CreatePost(objectId, mapPostRequest(request.Post))
	response := pb.NewPostRequest{
		Post: mapPost(newPost)
	}
	return response, err
}

//todo: updateLikes, updateDislikes, CreateComment