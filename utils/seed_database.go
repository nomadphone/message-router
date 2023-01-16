//go:build exclude
// +build exclude

package main

import (
	"log"

	"github.com/nomadphone/message-router/database"
	"github.com/nomadphone/message-router/users"
)

func main() {
	client, ctx, cancel := database.GetClient()
	defer client.Disconnect(ctx)
	defer cancel()

	collection := client.Database("nomadphone").Collection("users")
	user := users.User{
		TelegramUsername: "bia_rm",
		Name:             "Bianca Rosa",
		TwillioPhone:     "14155296858",
	}
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		panic(err)
	}
	log.Printf("Inserted user with ID: %s", result.InsertedID)
}
