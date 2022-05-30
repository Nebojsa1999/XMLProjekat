package persistence

import (
	"context"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CompanyRegistrationRequestsCollection = "company_registration_requests"
)

type CompanyRegistrationRequestMongoDBStore struct {
	companyRegistrationRequests *mongo.Collection
}

func NewCompanyRegistrationRequestMongoDBStore(client *mongo.Client) domain.CompanyRegistrationRequestStore {
	companyRegistrationRequests := client.Database(DATABASE).Collection(CompanyRegistrationRequestsCollection)

	return &CompanyRegistrationRequestMongoDBStore{
		companyRegistrationRequests: companyRegistrationRequests,
	}
}

func (store *CompanyRegistrationRequestMongoDBStore) GetAll() ([]*domain.CompanyRegistrationRequest, error) {
	filter := bson.D{{}}

	return store.filter(filter)
}

func (store *CompanyRegistrationRequestMongoDBStore) Get(id primitive.ObjectID) (*domain.CompanyRegistrationRequest, error) {
	filter := bson.M{"_id": id}
	existingCompanyRegistrationRequest, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingCompanyRegistrationRequest, nil
}

func (store *CompanyRegistrationRequestMongoDBStore) GetByOwnerId(ownerId primitive.ObjectID) (*domain.CompanyRegistrationRequest, error) {
	filter := bson.M{"owner_id": ownerId}
	existingCompanyRegistrationRequest, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingCompanyRegistrationRequest, nil
}

func (store *CompanyRegistrationRequestMongoDBStore) GetByName(name string) (*domain.CompanyRegistrationRequest, error) {
	filter := bson.M{"name": name}
	existingCompanyRegistrationRequest, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingCompanyRegistrationRequest, nil
}

func (store *CompanyRegistrationRequestMongoDBStore) CreateCompanyRegistrationRequest(request *domain.CompanyRegistrationRequest) (string, error) {
	result, err := store.companyRegistrationRequests.InsertOne(context.TODO(), request)
	if err != nil {
		return "Error occurred while inserting new company registration request into database!", err
	}

	request.Id = result.InsertedID.(primitive.ObjectID)
	return "Success: company registration request has been created.", nil
}

func (store *CompanyRegistrationRequestMongoDBStore) Update(updatedCompanyRegistrationRequest *domain.CompanyRegistrationRequest) (string, *domain.CompanyRegistrationRequest, error) {
	filter := bson.M{"_id": updatedCompanyRegistrationRequest.Id}
	update := bson.M{"$set": updatedCompanyRegistrationRequest}

	_, err := store.companyRegistrationRequests.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "Error occurred during update of company registration request!", nil, err
	}

	return "Success: company registration request has been updated.", updatedCompanyRegistrationRequest, nil
}

func (store *CompanyRegistrationRequestMongoDBStore) DeleteAll() (string, error) {
	_, err := store.companyRegistrationRequests.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return "Error occurred during deleting all company registration requests!", err
	}

	return "Success: all company registration requests have been deleted.", nil
}

func (store *CompanyRegistrationRequestMongoDBStore) filter(filter interface{}) ([]*domain.CompanyRegistrationRequest, error) {
	cursor, err := store.companyRegistrationRequests.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeIntoCompanyRegistrationRequests(cursor)
}

func (store *CompanyRegistrationRequestMongoDBStore) filterOne(filter interface{}) (companyRegistrationRequest *domain.CompanyRegistrationRequest, err error) {
	result := store.companyRegistrationRequests.FindOne(context.TODO(), filter)
	err = result.Decode(&companyRegistrationRequest)
	if err != nil {
		return nil, err
	}

	return companyRegistrationRequest, nil
}

func decodeIntoCompanyRegistrationRequests(cursor *mongo.Cursor) (companyRegistrationRequests []*domain.CompanyRegistrationRequest, err error) {
	for cursor.Next(context.TODO()) {
		var companyRegistrationRequest domain.CompanyRegistrationRequest
		err = cursor.Decode(&companyRegistrationRequest)
		if err != nil {
			return
		}
		companyRegistrationRequests = append(companyRegistrationRequests, &companyRegistrationRequest)
	}
	err = cursor.Err()

	return
}
