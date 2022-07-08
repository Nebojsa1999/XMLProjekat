package persistence

import (
	"context"

	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	JobsCollection = "jobs"
)

type JobMongoDBStore struct {
	jobs *mongo.Collection
}

func NewJobMongoDBStore(client *mongo.Client) domain.JobStore {
	jobs := client.Database(DATABASE).Collection(JobsCollection)

	return &JobMongoDBStore{
		jobs: jobs,
	}
}

func (store *JobMongoDBStore) GetAll() ([]*domain.Job, error) {
	filter := bson.D{{}}

	return store.filter(filter)
}

func (store *JobMongoDBStore) Get(id primitive.ObjectID) (*domain.Job, error) {
	filter := bson.M{"_id": id}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *JobMongoDBStore) CreateNewJob(job *domain.Job) (string, error) {
	result, err := store.jobs.InsertOne(context.TODO(), job)
	if err != nil {
		return "Error occurred while inserting new job into database!", err
	}

	job.Id = result.InsertedID.(primitive.ObjectID)
	return "Success: job has been created.", nil
}

func (store *JobMongoDBStore) Update(updatedJob *domain.Job) (string, *domain.Job, error) {
	filter := bson.M{"_id": updatedJob.Id}
	update := bson.M{"$set": updatedJob}

	_, err := store.jobs.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "Error occurred during update of job!", nil, err
	}

	return "Success: job has been updated.", updatedJob, nil
}

func (store *JobMongoDBStore) UpdateReviews(updatedJob *domain.Job) (string, *domain.Job, error) {
	filter := bson.M{"_id": updatedJob.Id}
	update := bson.M{"$set": updatedJob}

	_, err := store.jobs.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "Error occurred during update of job!", nil, err
	}

	return "Success: job has been updated.", updatedJob, nil
}

func (store *JobMongoDBStore) DeleteAll() (string, error) {
	_, err := store.jobs.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return "Error occurred during deleting all jobs!", err
	}

	return "Success: all jobs have been deleted.", nil
}

func (store *JobMongoDBStore) filter(filter interface{}) ([]*domain.Job, error) {
	cursor, err := store.jobs.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeIntoJobs(cursor)
}

func (store *JobMongoDBStore) filterOne(filter interface{}) (job *domain.Job, err error) {
	result := store.jobs.FindOne(context.TODO(), filter)
	err = result.Decode(&job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func decodeIntoJobs(cursor *mongo.Cursor) (jobs []*domain.Job, err error) {
	for cursor.Next(context.TODO()) {
		var job domain.Job
		err = cursor.Decode(&job)
		if err != nil {
			return
		}
		jobs = append(jobs, &job)
	}
	err = cursor.Err()

	return
}
