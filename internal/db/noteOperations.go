package db

import (
	"context"
	"fmt"
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
