package webhooks

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nomadphone/message-router/router"
	"github.com/nomadphone/message-router/users"
	"github.com/twilio/twilio-go/twiml"
)

type TwilioSMS twiml.MessagingMessage

func redirectSMS(from, to, body string) {
	user := users.GetUserFromTwillioPhone(to)

	telegramRouter := router.NewTelegramMessageRouter()
	telegramRouter.RouteMessage(from, user, body)
}

// Receives SMS from Twilio Webhook
func ReceiveSMS(w http.ResponseWriter, req *http.Request) {
	t := TwilioSMS{
		From: req.FormValue("From"),
		Body: req.FormValue("Body"),
		To:   req.FormValue("To"),
	}
	log.Printf("Received message from: %s with body: %s to phone number: %s\n", t.From, t.Body, t.To)
	message := &twiml.MessagingMessage{}
	body := &twiml.MessagingBody{}
	message.InnerElements = []twiml.Element{body}
	verbList := []twiml.Element{message}
	twimlResult, err := twiml.Messages(verbList)
	if err == nil {
		fmt.Println(twimlResult)
	} else {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(twimlResult))
	w.Header().Set("Content-Type", "application/xml")

	go redirectSMS(t.From, t.To, t.Body)
}
