package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func CreateDatabaseURI() string {
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	database := os.Getenv("MONGO_DATABASE")

	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, password, host, port, database)
}

func InitializeDatabase() *mongo.Database {

	database_uri := CreateDatabaseURI()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(database_uri))

	if err != nil {
		panic(fmt.Errorf("Error connecting to database: %v", err))
	}

	println(fmt.Sprintf("Connected to database: %v", database_uri))

	db := client.Database("proxy")
	return db
}

var Db = InitializeDatabase()
