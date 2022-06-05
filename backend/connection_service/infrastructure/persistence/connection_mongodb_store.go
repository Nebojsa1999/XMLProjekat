package persistence

import (
	"context"

	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE    = "connection_service"
	COLLECTION1 = "connections"
	COLLECTION2 = "privacy"
)

type ConnectionMongoDBStore struct {
	dbConnection *mongo.Collection
	dbPrivacy    *mongo.Collection
}

func NewConnectionMongoDBStore(client *mongo.Client) domain.ConnectionStore {
	connections := client.Database(DATABASE).Collection(COLLECTION1)
	privacy := client.Database(DATABASE).Collection(COLLECTION2)
	return &ConnectionMongoDBStore{
		connections: connections,
		privacy:     privacy,
	}
}

func (store *ConnectionMongoDBStore) Get(userId primitive.ObjectID) ([]*domain.Connection, error) {
	filter := bson.M{"$or": []bson.M{{"userAId": userId},
		{"userBId": userId}}}
	return store.filter(filter, id.Hex())
}

func (store *ConnectionMongoDBStore) CreateConnection(connection *domain.Connection) (*domain.Connection, error) {
	result, err := store.dbConnection.InsertOne(context.TODO(), connection)
	if err != nil {
		return nil, err
	}
	connection.Id = result.InsertedID.(primitive.ObjectID)
	return connection, nil
}

func (store *ConnectionMongoDBStore) CreateProfilePrivacy(privacy *domain.ProfilePrivacy) (*domain.ProfilePrivacy, error) {
	result, err := store.dbPrivacy.InsertOne(context.TODO(), privacy)
	if err != nil {
		return nil, err
	}
	privacy.Id = result.InsertedID.(primitive.ObjectID)
	return privacy, nil
}

func (store *ConnectionMongoDBStore) DeleteAll() error {
	_, err := store.dbConnection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	_, err = store.dbPrivacy.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionMongoDBStore) DeleteConnection(id primitive.ObjectID) error {
	filter := bson.M{"id": Id}
	_, err = store.dbConnection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionMongoDBStore) UpdateConnection(id string) (*domain.Connection, error) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"id": Id}
	connection, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	_, err = store.dbConnection.UpdateOne(context.TODO(), filter, bson.D{})
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (store *ConnectionMongoDBStore) filter(filter interface{}) ([]*domain.Connection, error) {
	connections := store.dbConnection.Collection(COLLECTION1 + id)
	cursor, err := connections.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *ConnectionMongoDBStore) filterOne(filter interface{}) (connection *domain.Connection, err error) {
	result := store.dbConnection.FindOne(context.TODO(), filter)
	err = result.Decode(&connection)
	return
}

func (store *ConnectionMongoDBStore) filterOnePrivacy(filter interface{}) (privacy *domain.ProfilePrivacy, err error) {
	result := store.dbPrivacy.FindOne(context.TODO(), filter)
	err = result.Decode(&privacy)
	return
}

func decode(cursor *mongo.Cursor) (connections []*domain.Connection, err error) {
	for cursor.Next(context.TODO()) {
		var Connection domain.Connection
		err = cursor.Decode(&Connection)
		if err != nil {
			return
		}
		connections = append(connections, &Connection)
	}
	err = cursor.Err()
	return
}
