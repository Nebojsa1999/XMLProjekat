package persistence

import (
	"context"

	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	InterviewsCollection = "interviews"
)

type InterviewMongoDBStore struct {
	interviews *mongo.Collection
}

func NewInterviewMongoDBStore(client *mongo.Client) domain.InterviewStore {
	interviews := client.Database(DATABASE).Collection(InterviewsCollection)

	return &InterviewMongoDBStore{
		interviews: interviews,
	}
}

func (store *InterviewMongoDBStore) GetAll() ([]*domain.Interview, error) {
	filter := bson.D{{}}

	return store.filter(filter)
}

func (store *InterviewMongoDBStore) Get(id primitive.ObjectID) (*domain.Interview, error) {
	filter := bson.M{"_id": id}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *InterviewMongoDBStore) GetByCompanyId(companyId primitive.ObjectID) (*domain.Interview, error) {
	filter := bson.M{"company_id": companyId}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *InterviewMongoDBStore) CreateNewInterview(interview *domain.Interview) (string, error) {
	result, err := store.interviews.InsertOne(context.TODO(), interview)
	if err != nil {
		return "Error occurred while inserting new interview into database!", err
	}

	interview.Id = result.InsertedID.(primitive.ObjectID)
	return "Success: interview has been created.", nil
}

func (store *InterviewMongoDBStore) Update(updatedInterview *domain.Interview) (string, *domain.Interview, error) {
	filter := bson.M{"_id": updatedInterview.Id}
	update := bson.M{"$set": updatedInterview}

	_, err := store.interviews.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "Error occurred during update of interview!", nil, err
	}

	return "Success: interview has been updated.", updatedInterview, nil
}

func (store *InterviewMongoDBStore) DeleteAll() (string, error) {
	_, err := store.interviews.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return "Error occurred during deleting all interviews!", err
	}

	return "Success: all interviews have been deleted.", nil
}

func (store *InterviewMongoDBStore) filter(filter interface{}) ([]*domain.Interview, error) {
	cursor, err := store.interviews.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeIntoInterviews(cursor)
}

func (store *InterviewMongoDBStore) filterOne(filter interface{}) (interview *domain.Interview, err error) {
	result := store.interviews.FindOne(context.TODO(), filter)
	err = result.Decode(&interview)
	if err != nil {
		return nil, err
	}

	return interview, nil
}

func decodeIntoInterviews(cursor *mongo.Cursor) (interviews []*domain.Interview, err error) {
	for cursor.Next(context.TODO()) {
		var interview domain.Interview
		err = cursor.Decode(&interview)
		if err != nil {
			return
		}
		interviews = append(interviews, &interview)
	}
	err = cursor.Err()

	return
}
