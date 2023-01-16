package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nomadphone/message-router/webhooks"
)

func main() {
	http.HandleFunc("/twillio/sms", webhooks.ReceiveSMS)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("Listening on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
