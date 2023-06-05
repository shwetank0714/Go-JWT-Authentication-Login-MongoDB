package helpers

import (
	"context"
	"log"

	dbInstance "github.com/shwetank0714/jwtmodfile/database"
	model "github.com/shwetank0714/jwtmodfile/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	collection = dbInstance.CreateDatabaseInstance()
}

// Insert the User into Database
func CreateUser(user *model.CreateUser) model.User{
	
	var newUser model.User
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	newUser.Phone = user.Phone
	newUser.Password = user.Password
	newUser.Email = user.Email
	newUser.ID = primitive.NewObjectID()
	token, _ := GetJwtToken(newUser.ID.String())
	newUser.Token = token

	insertedUserResult, err := collection.InsertOne(context.Background(), newUser)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("INSERTED RESULT DATA", insertedUserResult)
	log.Println("Movie inserted succesfully in DB, ID: ", insertedUserResult.InsertedID)

	return newUser
}

// Get All users list --- 
func GetAllUsers() []primitive.M {

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var allUsersList []primitive.M

	for cursor.Next(context.Background()) {
		var user bson.M

		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}

		allUsersList = append(allUsersList, user)
	}

	defer cursor.Close(context.Background())

	return allUsersList
}
