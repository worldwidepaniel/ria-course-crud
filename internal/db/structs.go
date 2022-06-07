package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UID          primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password"`
}

type Note struct {
	Note_ID       primitive.ObjectID `json:",omitempty" bson:"_id, omitempty"`
	UID           primitive.ObjectID `json:",omitempty" bson:"uid, omitempty"`
	Categories    []string           `json:"categories" bson:"categories" binding:"required"`
	Creation_date int                `json:"creation_date" bson:"creation_date"`
	Content       string             `json:"content" bson:"content" binding:"required"`
}
