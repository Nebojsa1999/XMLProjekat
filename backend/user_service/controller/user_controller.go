package controller

import (
	"context"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/database"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetInfoForLoggedInUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var existingUser model.User
	usersCollection := GetAllUsers()
	filter := bson.D{{"id", claims.Issuer}}
	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&existingUser); err != nil {
		log.Fatal(err)
	}

	return c.JSON(existingUser)
}

func GetAllUsers() *mongo.Collection {
	usersCollection := database.DB.Database("users-db").Collection("users")

	return usersCollection
}

func GetUserByUsername(username string) model.User {
	var user model.User
	usersCollection := GetAllUsers()
	filter := bson.D{{"username", username}}
	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		log.Fatal(err)
	}

	return user
}

func getAllUsernames() []string {
	var usersCollection = GetAllUsers()
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var allUsernames []string
	for cursor.Next(context.TODO()) {
		var elem model.User
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		allUsernames = append(allUsernames, elem.Username)
	}

	return allUsernames
}

func isUsernameAlreadyTaken(username string) bool {
	allUsernames := getAllUsernames()
	for i := 0; i < len(allUsernames); i++ {
		if allUsernames[i] == username {
			return true
		}
	}

	return false
}

func getAllEmails() []string {
	var usersCollection = GetAllUsers()
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var allEmails []string
	for cursor.Next(context.TODO()) {
		var elem model.User
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		allEmails = append(allEmails, elem.Email)
	}

	return allEmails
}

func isEmailAlreadyTaken(email string) bool {
	allEmails := getAllEmails()
	for i := 0; i < len(allEmails); i++ {
		if allEmails[i] == email {
			return true
		}
	}

	return false
}
