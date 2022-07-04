package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var (
// 	username = "prueba"
// 	password = "password"
// 	host     = "localhost"
// 	port     = 27017
// 	database = "tutorial"
// )

func GetCollection(collectionName string) *mongo.Collection {

	credential := options.Credential{
		AuthMechanism: os.Getenv("AUTH_MECHANISM"),
		Username:      os.Getenv("USERNAME_DB"),
		Password:      os.Getenv("PASSWORD_DB"),
		AuthSource:    os.Getenv("DATABASE_NAME"),
	}
	portDatabase := os.Getenv("PORT_DATABASE")
	hostDatabase := os.Getenv("HOST_DATABASE")
	databaseName := os.Getenv("DATABASE_NAME")
	uri := fmt.Sprintf("mongodb://%s:%s/?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000", hostDatabase, portDatabase)
	fmt.Println("Uri ", uri)
	// mongodb://prueba:password@localhost:27017/?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=tutorial&authMechanism=SCRAM-SHA-256
	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		panic(fmt.Sprintf("error create client %v", err.Error()))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(fmt.Sprintf("error connect %v", err.Error()))
	}
	return client.Database(databaseName).Collection((collectionName))
}
