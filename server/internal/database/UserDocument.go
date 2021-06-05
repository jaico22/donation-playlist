package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserDocument struct {
	ID             primitive.ObjectID `bson:"id,omitempty"`
	FacebookUserId string             `bson:"facebookUserId"`
}
