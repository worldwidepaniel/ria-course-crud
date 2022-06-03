package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

const dbName = "todoDB"

func GetUser(userEmail string) bson.D {
	connection := Connect()

	defer Close(connection)

	collection := connection.Database(dbName).Collection("users")
	filter := bson.D{{
		"email", bson.D{{"$eq", userEmail}},
	}}

	var result bson.D
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil
	}

	return result
}
