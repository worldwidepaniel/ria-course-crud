package db

import (
	"context"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddNote(noteData Note) string {
	connection := Connect()

	defer Close(connection)

	collection := connection.Database(dbName).Collection("notes")
	result, err := collection.InsertOne(context.TODO(), noteData)
	if err != nil {
		return fmt.Sprintf("%e", err)
	}
	notes, _ := GetNotes()
	AddToSearchEngine(notes)
	return fmt.Sprintf("%s", result.InsertedID)
}

func DeleteNote(noteID primitive.ObjectID, uid primitive.ObjectID) error {
	connection := Connect()

	defer Close(connection)

	collection := connection.Database(dbName).Collection("notes")

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", noteID}},
				bson.D{{"uid", uid}},
			}},
	}
	opts := options.Delete().SetHint(bson.D{{"_id", 1}})

	_, err := collection.DeleteMany(context.TODO(), filter, opts)
	if err != nil {
		return err
	}
	notes, _ := GetNotes()
	AddToSearchEngine(notes)
	return nil
}

func ModifyNote(noteData Note) error {
	connection := Connect()

	defer Close(connection)

	collection := connection.Database(dbName).Collection("notes")

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", noteData.Note_ID}},
				bson.D{{"uid", noteData.UID}},
			}},
	}

	update := bson.D{{"$set", bson.D{{"content", noteData.Content}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("error while modifying note")
	}
	notes, _ := GetNotes()
	AddToSearchEngine(notes)
	return nil
}

func CountNotes(userID primitive.ObjectID) (int64, error) {
	connection := Connect()
	defer Close(connection)

	collection := connection.Database(dbName).Collection("notes")
	filter := bson.D{{"uid", userID}}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetNote(userID primitive.ObjectID, noteID primitive.ObjectID) (Note, error) {
	connection := Connect()

	defer Close(connection)
	collection := connection.Database(dbName).Collection("notes")
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", noteID}},
				bson.D{{"uid", userID}},
			}},
	}
	var results Note
	if err := collection.FindOne(context.TODO(), filter).Decode(&results); err != nil {
		return Note{}, fmt.Errorf("error getting from cursor")
	}
	return results, nil
}

func GetUserNotes(limit string, offset string, userID primitive.ObjectID) ([]Note, error) {
	connection := Connect()

	defer Close(connection)

	collection := connection.Database(dbName).Collection("notes")

	filter := bson.D{{"uid", bson.D{{"$eq", userID}}}}
	projection := bson.D{}
	parsedOffset, err := strconv.ParseInt(offset, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error while parsing offset")
	}
	opts := options.Find().SetSkip(parsedOffset).SetLimit(50).SetProjection(projection)

	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("error while getting notes")
	}

	var results []Note
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, fmt.Errorf("error getting from cursor")
	}
	return results, nil
}

func GetNotes() ([]Note, error) {
	connection := Connect()

	defer Close(connection)

	collection := connection.Database(dbName).Collection("notes")

	filter := bson.D{}
	projection := bson.D{}
	opts := options.Find().SetProjection(projection)

	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("error while getting notes")
	}

	var results []Note
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, fmt.Errorf("error getting from cursor")
	}
	return results, nil
}
