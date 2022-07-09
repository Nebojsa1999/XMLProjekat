package application

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InterviewService struct {
	store domain.InterviewStore
}

func NewInterviewService(store domain.InterviewStore) *InterviewService {
	return &InterviewService{
		store: store,
	}
}

func (service *InterviewService) Get(id primitive.ObjectID) (*domain.Interview, error) {
	return service.store.Get(id)
}

func (service *InterviewService) GetByCompanyId(companyId primitive.ObjectID) (*domain.Interview, error) {
	return service.store.GetByCompanyId(companyId)
}

func (service *InterviewService) GetAll() ([]*domain.Interview, error) {
	return service.store.GetAll()
}

func (service *InterviewService) CreateNewInterview(interview *domain.Interview) (string, error) {
	existingInterview, _ := service.store.Get(interview.Id)
	interview.Id = primitive.NewObjectID()
	if existingInterview != nil {
		return "Interview with the same id already exists.", nil
	}

	return service.store.CreateNewInterview(interview)
}

func (service *InterviewService) Update(modifiedInterview *domain.Interview) (string, *domain.Interview, error) {
	interviewInDatabase, _ := service.store.Get(modifiedInterview.Id)
	if interviewInDatabase == nil {
		return "Interview with given id does not exist.", nil, nil

	}

	interviewInDatabase.Position = modifiedInterview.Position
	interviewInDatabase.Title = modifiedInterview.Title
	interviewInDatabase.YearOfInterview = modifiedInterview.YearOfInterview
	interviewInDatabase.HRInterview = modifiedInterview.HRInterview
	interviewInDatabase.TechnicalInterview = modifiedInterview.TechnicalInterview

	return service.store.Update(interviewInDatabase)
}
