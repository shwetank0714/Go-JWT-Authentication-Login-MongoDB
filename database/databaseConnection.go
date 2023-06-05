package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoConnectionUrl = "mongodb+srv://dshwetank1:gakuccespif0@cluster0.mwqm8z6.mongodb.net/?retryWrites=true&w=majority"
const databaseName = "JwtDatabase"
const collectionName = "User"

// Taking the reference of MongoDB collection (IMP)

func CreateDatabaseInstance() *mongo.Collection {
	// client options
	clientOptions := options.Client().ApplyURI(mongoConnectionUrl)

	// connect to mongodb
	// can use the context.Background() ass well here
	// context is required every time, whenever the connection is need to be established to some reference machine
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mogno DB connection succesful")

	collection := client.Database(databaseName).Collection(collectionName)
	// reference of the collection

	fmt.Println("Collection reference is ready")

	return collection
}
