package db

type User struct {
	Name         string `bson:"name"`
	Email        string `bson:"email"`
	PasswordHash string `bson:"password"`
}
