package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type CompanyStore interface {
	GetAll() ([]*Company, error)
	Get(id primitive.ObjectID) (*Company, error)
	GetByOwnerId(ownerId primitive.ObjectID) (*Company, error)
	GetByName(name string) (*Company, error)
	RegisterANewCompany(company *Company) (string, error)
	Update(updatedCompany *Company) (string, *Company, error)
	DeleteAll() (string, error)
}
