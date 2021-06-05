package database

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpsertUser(user UserDto) {
	log.Println("Connecting to database...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client := getClient(ctx)
	collection := client.Database("main").Collection("users")
	userDocument := UserDocument{FacebookUserId: user.FacebookUserId}
	userJson, _ := json.Marshal(userDocument)
	log.Printf("Upserting user %v", string(userJson))
	r, err := collection.InsertOne(ctx, userDocument)
	if err != nil {
		log.Fatalf("Error inserting document: %v", err)
	}
	log.Printf("Inserted; ID=%v", r.InsertedID)
}

func GetUser(userId string) *UserDto {
	log.Printf("Getting user %v", userId)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client := getClient(ctx)
	collection := client.Database("main").Collection("users")
	var filter interface{}
	filter = bson.D{{"facebookUserId", userId}}
	r := collection.FindOne(ctx, filter)
	var user UserDocument
	err := r.Decode(&user)
	if err == mongo.ErrNoDocuments {
		log.Println("User not found")
		return nil
	}
	u, _ := json.Marshal(user)
	log.Printf("User found: %v", string(u))
	return &UserDto{FacebookUserId: user.FacebookUserId}
}
