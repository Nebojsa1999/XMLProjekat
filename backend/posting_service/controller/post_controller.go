package controller

import (
	"context"
	"log"

	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/model"
	"github.com/gofiber/fiber/v2"

	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllPosts() *mongo.Collection {
	postsCollection := database.DB.Database("posts-db").Collection("posts")

	return postsCollection

}

func GetUserById(id uint) model.Post {
	var post model.Post
	postsCollection := GetAllPosts()
	filter := bson.D{{"id", id}}
	if err := postsCollection.FindOne(context.TODO(), filter).Decode(&post); err != nil {
		log.Fatal(err)
	}

	return post
}

func CreatePost(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	post := model.Post{
		Id:    0,
		Text:  data["text"],
		Link:  data["link"],
		Image: data["image"],
	}

	postsCollection := GetAllPosts()
	if _, err := postsCollection.InsertOne(context.TODO(), post); err != nil {
		log.Fatal(err)
	}

	return c.JSON(post)
}
