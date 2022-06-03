package db

import (
	"context"
	"fmt"

	"github.com/worldwidepaniel/ria-course-crud/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	connection, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.AppConfig.Database.Address))
	if err != nil {
		fmt.Errorf("error while connecting to database: %s", err)
	}
	return connection
}

func Close(connection *mongo.Client) {
	if err := connection.Disconnect(context.TODO()); err != nil {
		fmt.Errorf("error while closing connection to database: %s", err)
	}
}
