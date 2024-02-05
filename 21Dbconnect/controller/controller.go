package controller

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://vishalvx:<password>@try-out.g5ikcxe.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const collectionName = "WatchList"

var collection *mongo.Collection

func init() {

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to DB")

	connectColletion := client.Database(dbName).Collection(collectionName)
	log.Println(connectColletion.Name() + " Connected succefully")
}
