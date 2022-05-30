package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type CompanyRegistrationRequestStore interface {
	GetAll() ([]*CompanyRegistrationRequest, error)
	Get(id primitive.ObjectID) (*CompanyRegistrationRequest, error)
	GetByOwnerId(ownerId primitive.ObjectID) (*CompanyRegistrationRequest, error)
	GetByName(name string) (*CompanyRegistrationRequest, error)
	CreateCompanyRegistrationRequest(request *CompanyRegistrationRequest) (string, error)
	Update(updatedCompany *CompanyRegistrationRequest) (string, *CompanyRegistrationRequest, error)
	DeleteAll() (string, error)
}
