package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nomadphone/lib/database"
	"github.com/nomadphone/message-router/webhooks"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func ready(w http.ResponseWriter, r *http.Request) {
	client, ctx, cancel := database.GetClient()
	defer client.Disconnect(ctx)
	defer cancel()
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/health-check", healthCheck)
	http.HandleFunc("/ready-check", ready)
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
