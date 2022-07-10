package persistence

import (
	"context"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "connection_db"
	COLLECTION = "connections"
)

type ConnectionMongoDBStore struct {
	connections     *mongo.Collection
}

func NewConnectionMongoDBStore(client *mongo.Client) domain.ConnectionStore {
	connections := client.Database(DATABASE).Collection(COLLECTION)

	return &ConnectionMongoDBStore{
		connections:     connections,
	}
}

func (store *ConnectionMongoDBStore) Get(id primitive.ObjectID) (*domain.Connection, error) {
	filter := bson.M{"_id": id}

	return store.filterOneConnection(filter)
}

func (store *ConnectionMongoDBStore) GetAll() ([]*domain.Connection, error) {
	filter := bson.D{{}}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) GetConnectionOfFollowingType(id primitive.ObjectID) (*domain.Connection, error) {
	filter := bson.M{"_id": id, "type": domain.Following}

	return store.filterOneConnection(filter)
}

func (store *ConnectionMongoDBStore) GetAllConnectionsOfFollowingType() ([]*domain.Connection, error) {
	filter := bson.D{{"type", domain.Following}}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) GetConnectionOfBlockingType(id primitive.ObjectID) (*domain.Connection, error) {
	filter := bson.M{"_id": id, "type": domain.Blocking}

	return store.filterOneConnection(filter)
}

func (store *ConnectionMongoDBStore) GetAllConnectionsOfBlockingType() ([]*domain.Connection, error) {
	filter := bson.D{{"type", domain.Blocking}}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) GetByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	filter := bson.M{"$or": []bson.M{{"issuer_id": userId}, {"subject_id": userId}}}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) GetConnectionsOfFollowingTypeByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	filter := bson.M{"$or": []bson.M{{"type": domain.Following}, {"issuer_id": userId}, {"subject_id": userId}}}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) GetConnectionsOfBlockingTypeByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	filter := bson.M{"$or": []bson.M{{"type": domain.Blocking}, {"issuer_id": userId}, {"subject_id": userId}}}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) GetByTypeAndIssuerIdAndSubjectId(connectionUpdateDTO *domain.ConnectionUpdateDTO) (*domain.Connection, error) {
	filter := bson.M{"type": connectionUpdateDTO.Type, "issuer_id": connectionUpdateDTO.IssuerId,
		"subject_id": connectionUpdateDTO.SubjectId}

	return store.filterOneConnection(filter)
}

func (store *ConnectionMongoDBStore) GetFollowingByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	filter := bson.M{"type": domain.Following, "issuer_id": userId}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) GetFollowersByUserId(userId primitive.ObjectID) ([]*domain.Connection, error) {
	filter := bson.M{"type": domain.Following, "subject_id": userId}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) GetConnectionsInWhichTheGivenUserIsBlocker(userId primitive.ObjectID) ([]*domain.Connection, error) {
	filter := bson.M{"type": domain.Blocking, "issuer_id": userId}

	return store.filterConnections(filter)
}

func (store *ConnectionMongoDBStore) GetConnectionsInWhichTheGivenUserIsBlockedOne(userId primitive.ObjectID) ([]*domain.Connection, error) {
	filter := bson.M{"type": domain.Blocking, "subject_id": userId}

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

func (store *ConnectionMongoDBStore) Update(updatedConnection *domain.Connection) (*domain.Connection, error) {
	filter := bson.M{"_id": updatedConnection.Id}
	update := bson.M{"$set": updatedConnection}

	_, err := store.connections.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return updatedConnection, nil
}

func (store *ConnectionMongoDBStore) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := store.connections.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (store *ConnectionMongoDBStore) DeleteAll() error {
	_, err := store.connections.DeleteMany(context.TODO(), bson.D{{}})
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

	return decodeIntoConnections(cursor)
}

func (store *ConnectionMongoDBStore) filterOneConnection(filter interface{}) (connection *domain.Connection, err error) {
	result := store.connections.FindOne(context.TODO(), filter)
	err = result.Decode(&connection)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func decodeIntoConnections(cursor *mongo.Cursor) (connections []*domain.Connection, err error) {
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
