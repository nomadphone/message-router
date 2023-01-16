package users

import (
	"github.com/nomadphone/lib/database"
	"github.com/nomadphone/lib/models"
	"github.com/nomadphone/lib/phonenumbers"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserFromTwillioPhone(twillioPhone string) models.User {
	client, ctx, cancel := database.GetClient()
	defer client.Disconnect(ctx)
	defer cancel()
	var result models.User
	p := phonenumbers.NumbersOnly(twillioPhone)
	c := client.Database("nomadphone").Collection("users")
	filter := bson.M{"twilliophone": p}
	err := c.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}
