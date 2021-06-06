package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection = "users"
var usersDatabase = "main"

// UpsertUser inserts UserDto into the database
func UpsertUser(user UserDto) {
	log.Println("Inserting new user...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client := getClient(ctx)
	collection := getUsersCollection(client)

	userDocument := UserDocument{FacebookUserId: user.FacebookUserId}

	r, err := collection.InsertOne(ctx, userDocument)
	if err != nil {
		log.Fatalf("Error inserting document: %v", err)
		return
	}

	log.Printf("Inserted; ID=%v", r.InsertedID)
}

// GetUser returns a user from the database with matching userId
func GetUser(userId string) *UserDto {
	log.Printf("Getting user %v", userId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client := getClient(ctx)
	collection := getUsersCollection(client)

	var filter interface{}
	filter = bson.D{{"facebookUserId", userId}}
	r := collection.FindOne(ctx, filter)
	var user UserDocument
	err := r.Decode(&user)

	if err == mongo.ErrNoDocuments {
		log.Println("User not found")
		return nil
	}
	log.Println("User found")
	return &UserDto{FacebookUserId: user.FacebookUserId}
}

func getUsersCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(usersDatabase).Collection(usersCollection)
}
