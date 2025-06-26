package main

import (
	"log"

	"telegram-bot/bot"
	"telegram-bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config.Load()

	if config.TelegramToken == "" {
		log.Fatal("TELEGRAM_TOKEN не задан")
	}

	botAPI, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	botAPI.Debug = true
	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := botAPI.GetUpdatesChan(u)

	for update := range updates {
		bot.HandleMessage(botAPI, update)
	}
}
