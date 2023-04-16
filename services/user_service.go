package services

import (
	"context"
	"log"

	"github.com/wchopite/api-spark-v1/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	id			  string `bson:"id,omitempty" json:"id,omitempty"`
	shackName	  string `bson:"shackName,omitempty" json:"shackName,omitempty"`
	donatorStatus string `bson:"donator,omitempty" json:"donator,omitempty"`
}

//GetCurrencyBestPrice ...
func GetUser(id string) []User {
	log.Println(id)
	var filter bson.D = bson.D{}
	var result []User

	coll := database.GetCollection("users")

	opts := options.Find()

	if id != "" {
		filter = bson.D{{"id", id}}
		opts.SetLimit(1)
	}

	cursor, err := coll.Find(context.Background(), filter, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var object User
		err := cursor.Decode(&object)
		log.Println(object)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, object)
	}

	return result
}
