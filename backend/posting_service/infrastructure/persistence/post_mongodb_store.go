package persistence

import (
	"context"
	"fmt"
	"log"

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
	posts *mongo.Database
}

func (store *PostMongoDBStore) GetAllPosts() ([]*domain.Post, error) {
	filter := bson.D{{}}
	x := []string{"000000000000000000001111", "111100000000000000000000"}
	posts := []*domain.Post{}
	for _, id := range x {
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
		log.Fatal(err)
		return nil, err
	}
	post.LikesCount = post.LikesCount + 1

	filter := bson.M{"_id": liked_or_disliked_by.PostId}
	update := bson.D{
		{"$set", bson.D{{"likes", post.LikesCount}}},
	}

	insertResult, err := store.posts.Collection(COLLECTION+liked_or_disliked_by.Id.Hex()).UpdateOne(context.TODO(), filter,
		update)
	if err != nil {
		return nil, err
	}
	if insertResult.MatchedCount != 1 {
		log.Fatal(err, "Greska")
		return nil, err
	}
	return post, err
}

func (store *PostMongoDBStore) UpdateDislikes(liked_or_disliked_by *domain.LikeOrDislike) (*domain.Post, error) {
	post, err := store.GetPostFromUser(liked_or_disliked_by.Id, liked_or_disliked_by.PostId)
	if err != nil {
		log.Fatal(err)
	}
	post.DislikesCount = post.DislikesCount + 1

	filter := bson.M{"_id": liked_or_disliked_by.PostId}
	update := bson.D{
		{"$set", bson.D{{"dislikes", post.DislikesCount}}},
	}

	insertResult, err := store.posts.Collection(COLLECTION+liked_or_disliked_by.Id.Hex()).UpdateOne(context.TODO(), filter,
		update)
	if err != nil {
		return nil, err
	}
	if insertResult.MatchedCount != 1 {
		log.Fatal(err, "Greska")
		return nil, err
	}
	return post, err
}

func NewPostMongoDBStore(client *mongo.Client) domain.PostStore {

	posts := client.Database(DATABASE)
	return &PostMongoDBStore{
		posts: posts,
	}
}

func (store *PostMongoDBStore) GetPostFromUser(id primitive.ObjectID, post_id primitive.ObjectID) (post *domain.Post, err error) {

	filter := bson.M{"_id": post_id}
	posts := store.posts.Collection(COLLECTION + id.Hex())
	result := posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func (store *PostMongoDBStore) GetAllPostsFromUser(id primitive.ObjectID) (post []*domain.Post, err error) {
	filter := bson.D{{}}
	return store.filter(filter, id.Hex())
}

func (store *PostMongoDBStore) CreatePost(id primitive.ObjectID, post *domain.Post) (*domain.Post, error) {

	result, err := store.posts.Collection(COLLECTION+id.Hex()).InsertOne(context.TODO(), &domain.Post{
		Id:      primitive.NewObjectID(),
		OwnerId: id,
		Content: post.Content,
		Image:   post.Image,
		Link:    post.Link,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a document: ", result.InsertedID)

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

	insertResult, err := store.posts.Collection(COLLECTION+id.Hex()).UpdateOne(context.TODO(), filter,
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
	result, err := store.posts.ListCollectionNames(
		context.TODO(),
		bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	for _, collection := range result {
		store.posts.Collection(collection).DeleteMany(context.TODO(), bson.D{{}})
	}

}

func (store *PostMongoDBStore) filter(filter interface{}, id string) ([]*domain.Post, error) {
	posts := store.posts.Collection(COLLECTION + id)
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
