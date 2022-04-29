package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func Connect() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	postingDBHost := os.Getenv("POSTING_DB_HOST")
	postingDBPort := os.Getenv("POSTING_DB_PORT")
	clientOptions := options.Client().
		ApplyURI("mongodb://" + postingDBHost + ":" + postingDBPort + "/?connect=direct")

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
		fmt.Println("Connection to MongoDB \"posting_db\" closed.")
	}
}
