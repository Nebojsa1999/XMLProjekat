package application

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyRegistrationRequestService struct {
	store domain.CompanyRegistrationRequestStore
}

func NewCompanyRegistrationRequestService(store domain.CompanyRegistrationRequestStore) *CompanyRegistrationRequestService {
	return &CompanyRegistrationRequestService{
		store: store,
	}
}

func (service *CompanyRegistrationRequestService) Get(id primitive.ObjectID) (*domain.CompanyRegistrationRequest, error) {
	return service.store.Get(id)
}

func (service *CompanyRegistrationRequestService) GetAll() ([]*domain.CompanyRegistrationRequest, error) {
	return service.store.GetAll()
}

func (service *CompanyRegistrationRequestService) CreateCompanyRegistrationRequest(request *domain.CompanyRegistrationRequest) (string, error) {
	existingCompanyRegistrationRequest, _ := service.store.Get(request.Id)
	request.Id = primitive.NewObjectID()
	if existingCompanyRegistrationRequest != nil {
		return "Company registration request with the same id already exists.", nil
	}

	existingCompanyRegistrationRequest, _ = service.store.GetByName(request.Name)
	if existingCompanyRegistrationRequest != nil {
		return "Name is already taken, please choose another one.", nil
	}

	return service.store.CreateCompanyRegistrationRequest(request)
}

func (service *CompanyRegistrationRequestService) updateCompanyRegistrationRequestMethodBasis(modifiedRequest *domain.CompanyRegistrationRequest) (string, *domain.CompanyRegistrationRequest, error) {
	companyRegistrationRequestInDatabase, _ := service.store.Get(modifiedRequest.Id)
	if companyRegistrationRequestInDatabase == nil {
		return "Company registration request with given id does not exist.", nil, nil
	}

	if companyRegistrationRequestInDatabase.OwnerId != modifiedRequest.OwnerId {
		return "Owner of company in registration request cannot be changed.", nil, nil
	}

	if companyRegistrationRequestInDatabase.Status == enums.Accepted {
		return "Company registration request is already accepted and cannot be changed.", nil, nil
	}

	return "", companyRegistrationRequestInDatabase, nil
}

func (service *CompanyRegistrationRequestService) UpdateByOwner(modifiedRequest *domain.CompanyRegistrationRequest) (string, *domain.CompanyRegistrationRequest, error) {
	message, companyRegistrationRequestInDatabase, _ := service.updateCompanyRegistrationRequestMethodBasis(modifiedRequest)
	if companyRegistrationRequestInDatabase == nil {
		return message, nil, nil
	}

	companyRegistrationRequestInDatabaseWithSameName, _ := service.store.GetByName(modifiedRequest.Name)
	if companyRegistrationRequestInDatabaseWithSameName != nil {
		if companyRegistrationRequestInDatabaseWithSameName.Id != companyRegistrationRequestInDatabase.Id {
			return "Given name of company is already taken by another registration request.", nil, nil
		}
	}

	companyRegistrationRequestInDatabase.Status = enums.Pending
	companyRegistrationRequestInDatabase.ReasonForRejection = ""
	companyRegistrationRequestInDatabase.Name = modifiedRequest.Name
	companyRegistrationRequestInDatabase.Address = modifiedRequest.Address
	companyRegistrationRequestInDatabase.Email = modifiedRequest.Email
	companyRegistrationRequestInDatabase.Phone = modifiedRequest.Phone
	companyRegistrationRequestInDatabase.AreaOfWork = modifiedRequest.AreaOfWork
	companyRegistrationRequestInDatabase.Description = modifiedRequest.Description
	companyRegistrationRequestInDatabase.WorkCulture = modifiedRequest.WorkCulture

	return service.store.Update(companyRegistrationRequestInDatabase)
}

func (service *CompanyRegistrationRequestService) UpdateByAdministrator(modifiedRequest *domain.CompanyRegistrationRequest) (string, *domain.CompanyRegistrationRequest, error) {
	message, companyRegistrationRequestInDatabase, _ := service.updateCompanyRegistrationRequestMethodBasis(modifiedRequest)
	if companyRegistrationRequestInDatabase == nil {
		return message, nil, nil
	}

	if modifiedRequest.ReasonForRejection == "" {
		if modifiedRequest.Status == enums.Rejected {
			return "Reason for rejection must be written.", nil, nil
		} else {
			return "Registration request is not rejected and therefore cannot have a filled reason.", nil, nil
		}
	}

	companyRegistrationRequestInDatabase.Status = modifiedRequest.Status
	companyRegistrationRequestInDatabase.ReasonForRejection = modifiedRequest.ReasonForRejection

	return service.store.Update(companyRegistrationRequestInDatabase)
}
