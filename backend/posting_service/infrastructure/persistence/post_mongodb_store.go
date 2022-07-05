package persistence

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "post"
	COLLECTION = "posts"
)

type PostMongoDBStore struct {
	dbPost *mongo.Database
}

func NewPostMongoDBStore(client *mongo.Client) domain.PostStore {

	dbPost := client.Database(DATABASE)
	return &PostMongoDBStore{
		dbPost: dbPost,
	}
}

func (store *PostMongoDBStore) GetAllPosts(postIds []string) ([]*domain.Post, error) {
	filter := bson.D{{}}

	posts := []*domain.Post{}
	for _, id := range postIds {
		userPost, _ := store.filter(filter, id)
		for _, post := range userPost {
			posts = append(posts, post)
		}
	}
	return posts, nil
}

func (store *PostMongoDBStore) UpdateLikes(liked_or_disliked_by *domain.LikeOrDislike) (*domain.Post, error) {
	post, err := store.GetPostFromUser(liked_or_disliked_by.Id, liked_or_disliked_by.PostId)
	if err != nil {
		return nil, err
	}
	if has(post.WhoLiked, liked_or_disliked_by.LikedOrDislikedBy.Hex()) {
		return nil, errors.New("User already liked post")
	}
	post.WhoLiked = append(post.WhoLiked, liked_or_disliked_by.LikedOrDislikedBy.Hex())
	post.LikesCount = post.LikesCount + 1

	filter := bson.M{"_id": liked_or_disliked_by.PostId}
	update := bson.D{
		{"$set", bson.D{{"liked_by", post.WhoLiked}, {"likes", post.LikesCount}}},
	}

	insertResult, err := store.dbPost.Collection(COLLECTION+liked_or_disliked_by.Id.Hex()).UpdateOne(context.TODO(), filter,
		update)
	if err != nil {
		return nil, err
	}
	if insertResult.MatchedCount != 1 {
		return nil, fmt.Errorf("one document should've been updated")
	}
	return post, err
}

func (store *PostMongoDBStore) UpdateDislikes(liked_or_disliked_by *domain.LikeOrDislike) (*domain.Post, error) {
	post, err := store.GetPostFromUser(liked_or_disliked_by.Id, liked_or_disliked_by.PostId)
	if err != nil {
		return nil, err
	}
	if has(post.WhoDisliked, liked_or_disliked_by.LikedOrDislikedBy.Hex()) {
		return nil, errors.New("User already disliked post")
	}
	post.WhoDisliked = append(post.WhoDisliked, liked_or_disliked_by.LikedOrDislikedBy.Hex())
	post.DislikesCount = post.DislikesCount + 1

	filter := bson.M{"_id": liked_or_disliked_by.PostId}
	update := bson.D{
		{"$set", bson.D{{"disliked_by", post.WhoDisliked}, {"dislikes", post.DislikesCount}}},
	}

	insertResult, err := store.dbPost.Collection(COLLECTION+liked_or_disliked_by.Id.Hex()).UpdateOne(context.TODO(), filter,
		update)
	if err != nil {
		return nil, err
	}
	if insertResult.MatchedCount != 1 {
		return nil, fmt.Errorf("one document should've been updated")
	}
	return post, err
}

func (store *PostMongoDBStore) GetPostFromUser(id primitive.ObjectID, post_id primitive.ObjectID) (post *domain.Post, err error) {

	filter := bson.M{"_id": post_id}
	posts := store.dbPost.Collection(COLLECTION + id.Hex())
	result := posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func (store *PostMongoDBStore) GetAllPostsFromUser(id primitive.ObjectID) (post []*domain.Post, err error) {
	filter := bson.D{{}}
	return store.filter(filter, id.Hex())
}

func (store *PostMongoDBStore) CreatePost(id primitive.ObjectID, post *domain.Post) (*domain.Post, error) {

	insertResult, err := store.dbPost.Collection(COLLECTION+id.Hex()).InsertOne(context.TODO(), &domain.Post{
		Id:            primitive.NewObjectID(),
		OwnerId:       id,
		Content:       post.Content,
		Image:         post.Image,
		LikesCount:    post.LikesCount,
		DislikesCount: post.DislikesCount,
		Comments:      post.Comments,
		Link:          post.Link,
		WhoLiked:      post.WhoLiked,
		WhoDisliked:   post.WhoDisliked,
		PostedAt:      primitive.NewDateTimeFromTime(time.Now()),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a document: ", insertResult.InsertedID)

	return post, nil
}

func (store *PostMongoDBStore) CreateComment(id primitive.ObjectID, post_id primitive.ObjectID, comment *domain.Comment) (*domain.Comment, error) {
	post, err := store.GetPostFromUser(id, post_id)
	if err != nil {
		log.Fatal(err)
	}
	post.Comments = append(post.Comments, *comment)

	filter := bson.M{"_id": post_id}
	update := bson.D{
		{"$set", bson.D{{"comments", post.Comments}}},
	}

	insertResult, err := store.dbPost.Collection(COLLECTION+id.Hex()).UpdateOne(context.TODO(), filter,
		update)
	if err != nil {
		return nil, err
	}
	if insertResult.MatchedCount != 1 {
		log.Fatal(err, "one document should've been updated")
		return nil, err
	}
	return comment, err
}

func (store *PostMongoDBStore) DeleteAll() {
	result, err := store.dbPost.ListCollectionNames(
		context.TODO(),
		bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	for _, collection := range result {
		store.dbPost.Collection(collection).DeleteMany(context.TODO(), bson.D{{}})
	}

}

func (store *PostMongoDBStore) filter(filter interface{}, id string) ([]*domain.Post, error) {
	posts := store.dbPost.Collection(COLLECTION + id)
	cursor, err := posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (posts []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var post domain.Post
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}

func has(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
