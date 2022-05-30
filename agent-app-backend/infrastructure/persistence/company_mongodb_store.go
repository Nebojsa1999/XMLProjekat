package persistence

import (
	"context"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CompaniesCollection = "companies"
)

type CompanyMongoDBStore struct {
	companies *mongo.Collection
}

func NewCompanyMongoDBStore(client *mongo.Client) domain.CompanyStore {
	companies := client.Database(DATABASE).Collection(CompaniesCollection)

	return &CompanyMongoDBStore{
		companies: companies,
	}
}

func (store *CompanyMongoDBStore) GetAll() ([]*domain.Company, error) {
	filter := bson.D{{}}

	return store.filter(filter)
}

func (store *CompanyMongoDBStore) Get(id primitive.ObjectID) (*domain.Company, error) {
	filter := bson.M{"_id": id}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *CompanyMongoDBStore) GetByOwnerId(ownerId primitive.ObjectID) (*domain.Company, error) {
	filter := bson.M{"owner_id": ownerId}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *CompanyMongoDBStore) GetByName(name string) (*domain.Company, error) {
	filter := bson.M{"name": name}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *CompanyMongoDBStore) RegisterANewCompany(company *domain.Company) (string, error) {
	result, err := store.companies.InsertOne(context.TODO(), company)
	if err != nil {
		return "Error occurred while inserting new company into database!", err
	}

	company.Id = result.InsertedID.(primitive.ObjectID)
	return "Success: company has been registered.", nil
}

func (store *CompanyMongoDBStore) Update(updatedCompany *domain.Company) (string, *domain.Company, error) {
	filter := bson.M{"_id": updatedCompany.Id}
	update := bson.M{"$set": updatedCompany}

	_, err := store.companies.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "Error occurred during update of company!", nil, err
	}

	return "Success: company has been updated.", updatedCompany, nil
}

func (store *CompanyMongoDBStore) DeleteAll() (string, error) {
	_, err := store.companies.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return "Error occurred during deleting all companies!", err
	}

	return "Success: all companies have been deleted.", nil
}

func (store *CompanyMongoDBStore) filter(filter interface{}) ([]*domain.Company, error) {
	cursor, err := store.companies.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeIntoCompanies(cursor)
}

func (store *CompanyMongoDBStore) filterOne(filter interface{}) (company *domain.Company, err error) {
	result := store.companies.FindOne(context.TODO(), filter)
	err = result.Decode(&company)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func decodeIntoCompanies(cursor *mongo.Cursor) (companies []*domain.Company, err error) {
	for cursor.Next(context.TODO()) {
		var company domain.Company
		err = cursor.Decode(&company)
		if err != nil {
			return
		}
		companies = append(companies, &company)
	}
	err = cursor.Err()

	return
}
