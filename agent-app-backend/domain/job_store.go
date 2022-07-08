package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type JobStore interface {
	GetAll() ([]*Job, error)
	Get(id primitive.ObjectID) (*Job, error)
	CreateNewJob(job *Job) (string, error)
	Update(updatedJob *Job) (string, *Job, error)
	UpdateReviews(updatedJob *Job) (string, *Job, error)
	DeleteAll() (string, error)
}
