package router

import (
	"log"
	"os"

	"github.com/nomadphone/lib/models"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramMessageRouter struct {
	api *telegram.BotAPI
}

func NewTelegramMessageRouter() TelegramMessageRouter {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	bot, err := telegram.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return TelegramMessageRouter{api: bot}
}

func (t TelegramMessageRouter) RouteMessage(incomingPhone string, user models.User, body string) {
	log.Printf("Redirecting message with body: %s to telegram user: %s", body, user.TelegramUsername)
	body = "Heyoo!\nYou have a new message from " + incomingPhone + "\n\n-----\n" + body
	message := telegram.NewMessage(user.TelegramChatID, body)
	t.api.Send(message)
}
