package application

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	store domain.PostStore
}

func NewPostService(store domain.PostStore) *PostService {
	return &PostService{
		store: store,
	}
}

func (service *PostService) GetPostFromUser(id, post_id primitive.ObjectID) (*domain.Post, error) {
	return service.store.GetPostFromUser(id, post_id)
}

func (service *PostService) GetAllPosts() ([]*domain.Post, error) {
	return service.store.GetAllPosts()
}

func (service *PostService) CreatePost(id primitive.ObjectID, post *domain.Post) (*domain.Post, error) {
	newPost, err := service.store.CreatePost(id, post)
	if err != nil {
		return nil, err
	}
	return newPost, nil
}
func (service *PostService) GetAllPostsFromUser(id primitive.ObjectID) ([]*domain.Post, error) {
	return service.store.GetAllPostsFromUser(id)
}

func (service *PostService) CreateComment(id primitive.ObjectID, post_id primitive.ObjectID, comment *domain.Comment) (*domain.Comment, error) {
	newComment, err := service.store.CreateComment(id, post_id, comment)
	if err != nil {
		return nil, err
	}
	return newComment, nil
}

func (service *PostService) UpdateLikes(likeordislike *domain.LikeOrDislike) (*domain.Post, error) {
	updatedPost, err := service.store.UpdateLikes(likeordislike)
	if err != nil {
		return nil, err
	}
	return updatedPost, nil
}

func (service *PostService) UpdateDislikes(likeordislike *domain.LikeOrDislike) (*domain.Post, error) {
	updatedPost, err := service.store.UpdateDislikes(likeordislike)
	if err != nil {
		return nil, err
	}
	return updatedPost, nil
}
