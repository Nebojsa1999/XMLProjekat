package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentStore interface {
	GetAll() ([]*Comment, error)
	Get(id primitive.ObjectID) (*Comment, error)
	GetByCompanyId(companyId primitive.ObjectID) ([]*Comment, error)
	CreateNewComment(comment *Comment) (string, error)
	Update(updatedComment *Comment) (string, *Comment, error)
	DeleteAll() (string, error)
}
