package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddNote(NoteData Note) string {
	connection := Connect()

	defer Close(connection)

	collection := connection.Database(dbName).Collection("notes")
	result, err := collection.InsertOne(context.TODO(), NoteData)
	if err != nil {
		return fmt.Sprintf("%e", err)
	}
	return fmt.Sprintf("%s", result.InsertedID)
}

func DeleteNote(noteID primitive.ObjectID) error {
	connection := Connect()

	defer Close(connection)

	collection := connection.Database(dbName).Collection("notes")

	filter := bson.D{{"_id", bson.D{{"$eq", noteID}}}}
	opts := options.Delete().SetHint(bson.D{{"_id", 1}})

	_, err := collection.DeleteMany(context.TODO(), filter, opts)
	if err != nil {
		return err
	}
	return nil
}
