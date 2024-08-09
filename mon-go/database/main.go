package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db,_  = ConnectToMongoDB();


func ConnectToMongoDB()(*mongo.Client, error) {
	// Connect to MongoDB
	uri := "mongodb://localhost:27017";
	if uri == "" {
		return nil, fmt.Errorf("URI is empty")
	}

	clientOptions := options.Client().ApplyURI(uri);
	client, err := mongo.Connect(context.Background(), clientOptions);
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil);
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil

}

func GetDB() *mongo.Client {
	return db
}