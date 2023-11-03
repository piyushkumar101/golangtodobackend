package database

import (
	"context"
	"fmt"
	"log"

	// "log"

	// "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/options"
)

var collection *mongo.Collection

// func init() {
// 	loadTheEnv()
// }

// func loadTheEnv() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading the .env file")
// 	}
// }

func CreateDBInstance() *mongo.Client {
	// connectionString := os.Getenv("DB_URI")
	// dbName := os.Getenv("DB_NAME")
	// collName := os.Getenv("DB_COLLECTION_NAME")
	fmt.Printf("The connection string is : %v", "mongodb://localhost:27017/?directConnection=true")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/?directConnection=true")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb!")

	// collection = client.Database(dbName).Collection(collName)
	// fmt.Println("collection instance created")

	return client

}

var Client *mongo.Client = CreateDBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("test").Collection(collectionName)
	return collection
}
