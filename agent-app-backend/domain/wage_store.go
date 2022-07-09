package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type WageStore interface {
	GetAll() ([]*Wage, error)
	Get(id primitive.ObjectID) (*Wage, error)
	GetByCompanyId(companyId primitive.ObjectID) (*Wage, error)
	CreateNewWage(wage *Wage) (string, error)
	Update(updatedWage *Wage) (string, *Wage, error)
	DeleteAll() (string, error)
}
