package persistence

import (
	"context"

	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	WagesCollection = "wages"
)

type WageMongoDBStore struct {
	wages *mongo.Collection
}

func NewWageMongoDBStore(client *mongo.Client) domain.WageStore {
	wages := client.Database(DATABASE).Collection(WagesCollection)

	return &WageMongoDBStore{
		wages: wages,
	}
}

func (store *WageMongoDBStore) GetAll() ([]*domain.Wage, error) {
	filter := bson.D{{}}

	return store.filter(filter)
}

func (store *WageMongoDBStore) Get(id primitive.ObjectID) (*domain.Wage, error) {
	filter := bson.M{"_id": id}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *WageMongoDBStore) GetByCompanyId(companyId primitive.ObjectID) ([]*domain.Wage, error) {
	filter := bson.M{"company_id": companyId}
	//existingUser, err := store.filterOne(filter)
	//if err != nil {
	//	return nil, err
	//}

	return store.filter(filter)
}

func (store *WageMongoDBStore) CreateNewWage(wage *domain.Wage) (string, error) {
	result, err := store.wages.InsertOne(context.TODO(), wage)
	if err != nil {
		return "Error occurred while inserting new wage into database!", err
	}

	wage.Id = result.InsertedID.(primitive.ObjectID)
	return "Success: wage has been created.", nil
}

func (store *WageMongoDBStore) Update(updatedWage *domain.Wage) (string, *domain.Wage, error) {
	filter := bson.M{"_id": updatedWage.Id}
	update := bson.M{"$set": updatedWage}

	_, err := store.wages.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "Error occurred during update of wage!", nil, err
	}

	return "Success: wage has been updated.", updatedWage, nil
}

func (store *WageMongoDBStore) DeleteAll() (string, error) {
	_, err := store.wages.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return "Error occurred during deleting all wages!", err
	}

	return "Success: all wages have been deleted.", nil
}

func (store *WageMongoDBStore) filter(filter interface{}) ([]*domain.Wage, error) {
	cursor, err := store.wages.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeIntoWages(cursor)
}

func (store *WageMongoDBStore) filterOne(filter interface{}) (wage *domain.Wage, err error) {
	result := store.wages.FindOne(context.TODO(), filter)
	err = result.Decode(&wage)
	if err != nil {
		return nil, err
	}

	return wage, nil
}

func decodeIntoWages(cursor *mongo.Cursor) (wages []*domain.Wage, err error) {
	for cursor.Next(context.TODO()) {
		var wage domain.Wage
		err = cursor.Decode(&wage)
		if err != nil {
			return
		}
		wages = append(wages, &wage)
	}
	err = cursor.Err()

	return
}
