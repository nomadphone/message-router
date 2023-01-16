//go:build exclude
// +build exclude

package main

import (
	"log"

	"github.com/nomadphone/message-router/users"
)

func main() {
	user := users.GetUserFromTwillioPhone("+14155296858")
	log.Printf("Found user with name: %s and TwillioPhone: %s", user.Name, user.TwillioPhone)
}
