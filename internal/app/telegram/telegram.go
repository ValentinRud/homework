package telegram

import (
	"encoding/json"
	"fmt"
	"homework/config"
	"homework/internal/app/models"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Telegram() *models.AddUser {
	bot, err := tgbotapi.NewBotAPI(config.TeleToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)

		in := []byte(update.Message.Text)
		err := json.Unmarshal(in, &models.AddUser)
		if err != nil {
			fmt.Println(err)
		}
	}
}
