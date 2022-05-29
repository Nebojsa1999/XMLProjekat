package application

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyService struct {
	store domain.CompanyStore
}

func NewCompanyService(store domain.CompanyStore) *CompanyService {
	return &CompanyService{
		store: store,
	}
}

func (service *CompanyService) Get(id primitive.ObjectID) (*domain.Company, error) {
	return service.store.Get(id)
}

func (service *CompanyService) GetAll() ([]*domain.Company, error) {
	return service.store.GetAll()
}

func (service *CompanyService) RegisterANewCompany(company *domain.Company) (string, error) {
	existingCompany, _ := service.store.Get(company.Id)
	company.Id = primitive.NewObjectID()
	if existingCompany != nil {
		return "Company with the same id already exists.", nil
	}

	existingCompany, _ = service.store.GetByName(company.Name)
	if existingCompany != nil {
		return "Name is already taken, please choose another one.", nil
	}

	return service.store.RegisterANewCompany(company)
}

func (service *CompanyService) Update(modifiedCompany *domain.Company) (string, *domain.Company, error) {
	companyInDatabase, _ := service.store.Get(modifiedCompany.Id)
	if companyInDatabase == nil {
		return "Company with given id does not exist.", nil, nil
	}

	if companyInDatabase.OwnerId != modifiedCompany.OwnerId {
		return "Owner of company cannot be changed.", nil, nil
	}

	companyInDatabaseWithSameName, _ := service.store.GetByName(modifiedCompany.Name)
	if companyInDatabaseWithSameName != nil {
		if companyInDatabaseWithSameName.Id != companyInDatabase.Id {
			return "Given name is already taken by another company.", nil, nil
		}
	}

	companyInDatabase.Name = modifiedCompany.Name
	companyInDatabase.Address = modifiedCompany.Address
	companyInDatabase.Email = modifiedCompany.Email
	companyInDatabase.Phone = modifiedCompany.Phone
	companyInDatabase.AreaOfWork = modifiedCompany.AreaOfWork
	companyInDatabase.Description = modifiedCompany.Description
	companyInDatabase.WorkCulture = modifiedCompany.WorkCulture

	return service.store.Update(companyInDatabase)
}
