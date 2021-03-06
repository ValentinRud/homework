package telegram

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart     = "start"
	commandList      = "list"
	commandOrders    = "orders"
	commandPrices    = "prices"
	commandSyncOrder = "syncorder"
	// commandListUser = "list/{id:[0-9]+}"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	b.bot.Send(msg)

}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")
	limit := 25
	var err error
	switch message.Command() {
	case commandStart:
		msg.Text = "Ты ввел комманду СТАРТ"
		_, err := b.bot.Send(msg)
		return err

	case commandPrices:
		orders := b.binanceClient.ListTicket()
		for _, price := range orders {
			b.repositories.CreatePrice(*price)
		}

		msg.Text = "Курсы валют записаны в БД"
		_, err = b.bot.Send(msg)
		return err

	case commandList:

		list := b.repositories.ListSymbol()
		for i := 0; i < len(list); i += limit {
			minValue := 0

			if i+limit <= len(list) {
				minValue = i + limit
			} else {
				minValue = len(list)
			}

			batch := list[i:minValue]
			bytes, err := json.Marshal(batch)
			if err != nil {
				log.Fatalf("ERROR %s", err)
			}

			msg.Text = string(bytes)
			_, err = b.bot.Send(msg)
		}
		return err

	case commandSyncOrder:
		orders := b.binanceClient.ListOrders()
		for _, price := range orders {
			b.repositories.CreateOrder(*price)
		}

		msg.Text = "Ордера записаны в БД"
		_, err = b.bot.Send(msg)
		return err

	case commandOrders:
		list := b.binanceClient.ListOrders()
		for i := 0; i < len(list); i += limit {
			minValue := 0

			if i+limit <= len(list) {
				minValue = i + limit
			} else {
				minValue = len(list)
			}

			batch := list[i:minValue]
			bytes, err := json.Marshal(batch)
			if err != nil {
				log.Fatalf("ERROR %s", err)
			}

			msg.Text = string(bytes)
			_, err = b.bot.Send(msg)
		}
		return err
	// case commandListUser:
	// 	msg.Text = fmt.Sprint(b.repositories.FindById())
	// 	_, err := b.bot.Send(msg)
	// 	return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

// func (b *Bot) UnmarshalBody(message *tgbotapi.Message) bool {

// 	type botText models.User

// 	if err := json.Unmarshal([]byte(message.Text), &botText); err != nil {
// 		return false
// 	}
// 	return true
// }
