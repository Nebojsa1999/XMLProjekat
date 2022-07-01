package application

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobService struct {
	store domain.JobStore
}

func NewJobService(store domain.JobStore) *JobService {
	return &JobService{
		store: store,
	}
}

func (service *JobService) Get(id primitive.ObjectID) (*domain.Job, error) {
	return service.store.Get(id)
}

func (service *JobService) GetAll() ([]*domain.Job, error) {
	return service.store.GetAll()
}

func (service *JobService) CreateNewJob(job *domain.Job) (string, error) {
	existingJob, _ := service.store.Get(job.Id)
	job.Id = primitive.NewObjectID()
	if existingJob != nil {
		return "Job with the same id already exists.", nil
	}

	return service.store.CreateNewJob(job)
}

func (service *JobService) Update(modifiedJob *domain.Job) (string, *domain.Job, error) {
	jobInDatabase, _ := service.store.Get(modifiedJob.Id)
	if jobInDatabase == nil {
		return "Job with given id does not exist.", nil, nil

	}
	jobInDatabase.Position = modifiedJob.Position
	jobInDatabase.Description = modifiedJob.Description
	jobInDatabase.Requirements = modifiedJob.Requirements

	return service.store.Update(jobInDatabase)
}
