package main

import (
	"database/sql"
	"fmt"
	"homework/config"
	"homework/internal/app/api"
	"homework/internal/app/repositories"
	"homework/internal/app/telegram"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

var (
	apiKey    = os.Getenv("APIKEY")
	secretKey = os.Getenv("SECRETKEY")
)

func main() {
	db, err := initDb()
	fmt.Println("connect postgre")
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}

	Repo := repositories.NewSymbolRepository(db)

	Api := api.NewClientApi()
	futuresClient := binance.NewFuturesClient(apiKey, secretKey) // USDT-M Futures
	// deliveryClient := binance.NewDeliveryClient(apiKey, secretKey) // Coin-M Futures

	bot, err := tgbotapi.NewBotAPI(config.TeleToken)
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}
	bot.Debug = true

	telegramBot := telegram.NewBot(bot, Repo, Api)
	if err := telegramBot.Start(); err != nil {
		log.Fatalf("ERROR %s", err)
	}

}

func initDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		log.Fatalf("ERROR CREATING RECORD %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("ERROR CREATING RECORD %s", err)
	}
	return db, err
}
