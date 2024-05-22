package database

import (
	"context"
	"github.com/joho/godotenv"
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

	return "mongodb://" + user + ":" + password + "@" + host + ":" + port + "/" + database
}

func InitializeDatabase() *mongo.Database {

	godotenv.Load()
	database_uri := CreateDatabaseURI()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(database_uri))
	println("Connected to " + database_uri)

	if err != nil {
		panic(err)
	}

	db := client.Database("proxy")
	return db
}

var Db = InitializeDatabase()
