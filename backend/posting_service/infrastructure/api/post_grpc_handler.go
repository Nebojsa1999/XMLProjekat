package api

import (
	"context"
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/domain"

	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	pb.UnimplementedPostingServiceServer
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (handler *PostHandler) GetPostFromUser(ctx context.Context, request *pb.GetPostRequest) (*pb.GetResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	objectPostId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.GetPostFromUser(objectId, objectPostId)
	if err != nil {
		return nil, err
	}
	postPb := mapPost(post)
	response := &pb.GetResponse{
		Post: postPb,
	}
	return response, nil
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	posts, err := handler.service.GetAllPosts()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) GetAllPostsFromUser(ctx context.Context, request *pb.GetRequest) (*pb.GetAllResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	posts, err := handler.service.GetAllPostsFromUser(objectId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) CreatePost(ctx context.Context, request *pb.NewPostRequest) (*pb.NewPostResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	newPost, err := handler.service.CreatePost(objectId, mapPostRequest(request.Post))
	response := &pb.NewPostResponse{
		Post: mapPost(newPost),
	}
	return response, err
}

func (handler *PostHandler) CreateComment(ctx context.Context, request *pb.CommentOnPostRequest) (*pb.CommentOnPostResponse, error) {

	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	objectPostId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}

	newComment, err := handler.service.CreateComment(objectId, objectPostId, mapCommentOnPostRequest(request.Comment))
	response := &pb.CommentOnPostResponse{
		Comment: mapComment(newComment),
	}
	return response, err
}

//todo: updateLikes, updateDislikes
func (handler *PostHandler) InsertLikeOrDislike(ctx context.Context, request *pb.LikeOrDislikePostRequest) (*pb.GetResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	objectPostId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}

	objectLikedOrDislikedBy, err := primitive.ObjectIDFromHex(request.LikedOrDislikedBy)
	if err != nil {
		return nil, err
	}
	liked_or_disliked_by := &domain.LikeOrDislike{
		Id:                objectId,
		PostId:            objectPostId,
		LikedOrDislikedBy: objectLikedOrDislikedBy,
	}

	var updatedPost *domain.Post
	var err1 error
	if request.Type == "like" {
		updatedPost, err1 = handler.service.UpdateLikes(liked_or_disliked_by)
	}
	if request.Type == "dislike" {
		updatedPost, err1 = handler.service.UpdateDislikes(liked_or_disliked_by)
	}
	if err1 != nil {
		return nil, err1
	}
	response := &pb.GetResponse{
		Post: mapPost(updatedPost),
	}
	return response, err
}
