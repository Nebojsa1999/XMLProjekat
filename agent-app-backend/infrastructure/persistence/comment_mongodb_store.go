package persistence

import (
	"context"

	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CommentsCollection = "comments"
)

type CommentMongoDBStore struct {
	comments *mongo.Collection
}

func NewCommentMongoDBStore(client *mongo.Client) domain.CommentStore {
	comments := client.Database(DATABASE).Collection(CommentsCollection)

	return &CommentMongoDBStore{
		comments: comments,
	}
}

func (store *CommentMongoDBStore) GetAll() ([]*domain.Comment, error) {
	filter := bson.D{{}}

	return store.filter(filter)
}

func (store *CommentMongoDBStore) Get(id primitive.ObjectID) (*domain.Comment, error) {
	filter := bson.M{"_id": id}
	existingUser, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (store *CommentMongoDBStore) GetByCompanyId(companyId primitive.ObjectID) (*domain.Comment, error) {
	filter := bson.M{"company_id": companyId}
	existingComment, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}

	return existingComment, nil
}

func (store *CommentMongoDBStore) CreateNewComment(comment *domain.Comment) (string, error) {
	result, err := store.comments.InsertOne(context.TODO(), comment)
	if err != nil {
		return "Error occurred while inserting new comment into database!", err
	}

	comment.Id = result.InsertedID.(primitive.ObjectID)
	return "Success: comment has been created.", nil
}

func (store *CommentMongoDBStore) Update(updatedComment *domain.Comment) (string, *domain.Comment, error) {
	filter := bson.M{"_id": updatedComment.Id}
	update := bson.M{"$set": updatedComment}

	_, err := store.comments.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "Error occurred during update of comment!", nil, err
	}

	return "Success: comment has been updated.", updatedComment, nil
}

func (store *CommentMongoDBStore) DeleteAll() (string, error) {
	_, err := store.comments.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return "Error occurred during deleting all comments!", err
	}

	return "Success: all comments have been deleted.", nil
}

func (store *CommentMongoDBStore) filter(filter interface{}) ([]*domain.Comment, error) {
	cursor, err := store.comments.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeIntoComments(cursor)
}

func (store *CommentMongoDBStore) filterOne(filter interface{}) (comment *domain.Comment, err error) {
	result := store.comments.FindOne(context.TODO(), filter)
	err = result.Decode(&comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func decodeIntoComments(cursor *mongo.Cursor) (comments []*domain.Comment, err error) {
	for cursor.Next(context.TODO()) {
		var comment domain.Comment
		err = cursor.Decode(&comment)
		if err != nil {
			return
		}
		comments = append(comments, &comment)
	}
	err = cursor.Err()

	return
}
