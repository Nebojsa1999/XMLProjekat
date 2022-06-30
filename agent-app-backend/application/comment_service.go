package application

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentService struct {
	store domain.CommentStore
}

func NewCommentService(store domain.CommentStore) *CommentService {
	return &CommentService{
		store: store,
	}
}

func (service *CommentService) Get(id primitive.ObjectID) (*domain.Comment, error) {
	return service.store.Get(id)
}

func (service *CommentService) GetAll() ([]*domain.Comment, error) {
	return service.store.GetAll()
}

func (service *CommentService) CreateNewComment(comment *domain.Comment) (string, error) {
	existingComment, _ := service.store.Get(comment.Id)
	comment.Id = primitive.NewObjectID()
	if existingComment != nil {
		return "Comment with the same id already exists.", nil
	}

	return service.store.CreateNewComment(comment)
}

func (service *CommentService) Update(modifiedComment *domain.Comment) (string, *domain.Comment, error) {
	commentInDatabase, _ := service.store.Get(modifiedComment.Id)
	if commentInDatabase == nil {
		return "Comment with given id does not exist.", nil, nil

	}

	commentInDatabase.Position = modifiedComment.Position
	commentInDatabase.Engagement = modifiedComment.Engagement
	commentInDatabase.ExperienceLevel = modifiedComment.ExperienceLevel
	commentInDatabase.Content = modifiedComment.Content

	return service.store.Update(commentInDatabase)
}
