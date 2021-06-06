package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type document struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

func getClient(ctx context.Context) *mongo.Client {
	connectionString := "mongodb://localhost:C2y6yDjf5%2FR%2Bob0N8A7Cgv30VRDJIWEHLM%2B4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw%2FJw%3D%3D@localhost:10255/admin?ssl=true"
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.NewClient(clientOptions)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Unable to initialize mongo connection %v", err)
	}
	return client
}
