package persistence

import (
	"context"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE        = "agent_app_db"
	UsersCollection = "users"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(UsersCollection)

	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}

	return store.filter(filter)
}

func (store *UserMongoDBStore) Get(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *UserMongoDBStore) GetByUsername(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *UserMongoDBStore) GetByUsernameAndPassword(username string, password string) (*domain.User, error) {
	filter := bson.M{"username": username, "password": password}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *UserMongoDBStore) GetByEmail(email string) (*domain.User, error) {
	filter := bson.M{"email": email}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *UserMongoDBStore) RegisterANewUser(user *domain.User) (string, error) {
	result, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return "Error occurred while inserting new user into database!", err
	}

	user.Id = result.InsertedID.(primitive.ObjectID)
	return "Success: user has been registered.", nil
}

func (store *UserMongoDBStore) Update(updatedUser *domain.User) (string, *domain.User, error) {
	filter := bson.M{"_id": updatedUser.Id}
	update := bson.M{"$set": updatedUser}

	_, err := store.users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "Error occurred during update of user!", nil, err
	}

	return "Success: user has been updated.", updatedUser, nil
}

func (store *UserMongoDBStore) DeleteAll() (string, error) {
	_, err := store.users.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return "Error occurred during deleting all users!", err
	}

	return "Success: all users have been deleted.", nil
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeIntoUsers(cursor)
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (user *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func decodeIntoUsers(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var user domain.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()

	return
}
