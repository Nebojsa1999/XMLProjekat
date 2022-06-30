package application

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WageService struct {
	store domain.WageStore
}

func NewWageService(store domain.WageStore) *WageService {
	return &WageService{
		store: store,
	}
}

func (service *WageService) Get(id primitive.ObjectID) (*domain.Wage, error) {
	return service.store.Get(id)
}

func (service *WageService) GetAll() ([]*domain.Wage, error) {
	return service.store.GetAll()
}

func (service *WageService) CreateNewWage(wage *domain.Wage) (string, error) {
	existingWage, _ := service.store.Get(wage.Id)
	wage.Id = primitive.NewObjectID()
	if existingWage != nil {
		return "Wage with the same id already exists.", nil
	}

	return service.store.CreateNewWage(wage)
}

func (service *WageService) Update(modifiedWage *domain.Wage) (string, *domain.Wage, error) {
	wageInDatabase, _ := service.store.Get(modifiedWage.Id)
	if wageInDatabase == nil {
		return "Wage with given id does not exist.", nil, nil

	}

	wageInDatabase.Position = modifiedWage.Position
	wageInDatabase.Engagement = modifiedWage.Engagement
	wageInDatabase.ExperienceLevel = modifiedWage.ExperienceLevel
	wageInDatabase.NetoWage = modifiedWage.NetoWage

	return service.store.Update(wageInDatabase)
}
