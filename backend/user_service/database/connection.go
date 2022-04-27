package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var DB *mongo.Client

func Connect() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	userDBHost := os.Getenv("USER_DB_HOST")
	userDBPort := os.Getenv("USER_DB_PORT")
	clientOptions := options.Client().
		ApplyURI("mongodb://" + userDBHost + ":" + userDBPort + "/?connect=direct")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client

	<-quit

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB \"user_db\" closed.")
	}
}
