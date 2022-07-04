package persistence

import (
	"context"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE    = "connection_db"
	COLLECTION1 = "connections"
	COLLECTION2 = "profilesPrivacy"
)

type ConnectionMongoDBStore struct {
	connections     *mongo.Collection
	profilesPrivacy *mongo.Collection
}

func NewConnectionMongoDBStore(client *mongo.Client) domain.ConnectionStore {
	connections := client.Database(DATABASE).Collection(COLLECTION1)
	profilesPrivacy := client.Database(DATABASE).Collection(COLLECTION2)

	return &ConnectionMongoDBStore{
		connections:     connections,
		profilesPrivacy: profilesPrivacy,
	}
}

func (store *ConnectionMongoDBStore) Get(id primitive.ObjectID) (*domain.Connection, error) {
	filter := bson.M{"_id": id}

	return store.filterOneConnection(filter)
}

func (store *ConnectionMongoDBStore) GetByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	filter := bson.M{"$or": []bson.M{{"issuer_id": userId}, {"subject_id": userId}}}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) GetFollowingByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	filter := bson.M{"issuer_id": userId}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) GetFollowersByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	filter := bson.M{"subject_id": userId}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) Create(connection *domain.Connection) (*domain.Connection, error) {
	result, err := store.connections.InsertOne(context.TODO(), connection)
	if err != nil {
		return nil, err
	}

	connection.Id = result.InsertedID.(primitive.ObjectID)

	return connection, nil
}

func (store *ConnectionMongoDBStore) CreatePrivacy(privacy *domain.ProfilePrivacy) (*domain.ProfilePrivacy, error) {
	result, err := store.profilesPrivacy.InsertOne(context.TODO(), privacy)
	if err != nil {
		return nil, err
	}

	privacy.Id = result.InsertedID.(primitive.ObjectID)

	return privacy, nil
}

func (store *ConnectionMongoDBStore) DeleteAll() error {
	_, err := store.connections.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}

	_, err = store.profilesPrivacy.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}

	return nil
}

func (store *ConnectionMongoDBStore) Delete(id string) error {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": Id}
	_, err = store.connections.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (store *ConnectionMongoDBStore) Update(updatedConnection *domain.Connection) (*domain.Connection, error) {
	filter := bson.M{"_id": updatedConnection.Id}
	update := bson.M{"$set": updatedConnection}

	_, err := store.connections.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return updatedConnection, nil
}

func (store *ConnectionMongoDBStore) UpdatePrivacy(id primitive.ObjectID) error {
	filter := bson.M{"user_id": id}
	privacy, err := store.filterOnePrivacy(filter)
	if err != nil {
		return err
	}

	privacy.IsPrivate = !privacy.IsPrivate
	_, err = store.profilesPrivacy.UpdateOne(context.TODO(), filter, bson.D{{"$set",
		bson.M{"is_private": privacy.IsPrivate}}})
	if err != nil {
		return err
	}

	return nil
}

func (store *ConnectionMongoDBStore) CreateProfilePrivacy(privacy *domain.ProfilePrivacy) (*domain.ProfilePrivacy, error) {
	result, err := store.profilesPrivacy.InsertOne(context.TODO(), privacy)
	if err != nil {
		return nil, err
	}

	privacy.Id = result.InsertedID.(primitive.ObjectID)
	return privacy, nil
}

func (store *ConnectionMongoDBStore) DeleteProfilePrivacy(id primitive.ObjectID) error {
	filter := bson.M{"user_id": id}
	_, err := store.profilesPrivacy.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (store *ConnectionMongoDBStore) filterConnections(filter interface{}) ([]*domain.Connection, error) {
	cursor, err := store.connections.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cursor)
}

func (store *ConnectionMongoDBStore) filterOneConnection(filter interface{}) (connection *domain.Connection, err error) {
	result := store.connections.FindOne(context.TODO(), filter)
	err = result.Decode(&connection)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func (store *ConnectionMongoDBStore) filterOnePrivacy(filter interface{}) (privacy *domain.ProfilePrivacy, err error) {
	result := store.profilesPrivacy.FindOne(context.TODO(), filter)
	err = result.Decode(&privacy)
	if err != nil {
		return nil, nil
	}

	return privacy, nil
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
