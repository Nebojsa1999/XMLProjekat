package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type InterviewStore interface {
	GetAll() ([]*Interview, error)
	GetByCompanyId(companyId primitive.ObjectID) ([]*Interview, error)
	Get(id primitive.ObjectID) (*Interview, error)
	CreateNewInterview(interview *Interview) (string, error)
	Update(updatedInterview *Interview) (string, *Interview, error)
	DeleteAll() (string, error)
}
