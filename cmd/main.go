package main

import (
	"database/sql"
	"fmt"
	"homework/config"
	"homework/internal/app/gateways"
	"homework/internal/app/models"
	"homework/internal/app/repositories"
	"homework/internal/app/services"
	"homework/internal/app/telegram"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.TeleToken)
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}
	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatalf("ERROR %s", err)
	}

	db, err := initDb()
	fmt.Println("connect postgre")
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}

	var u models.User

	// u.SetName("name", "surname")
	// firstName := u.GetFirstName()
	// // чтение из json

	userRepo := repositories.NewUserRepository(db)
	err = userRepo.CreateUser(u)
	if err != nil {
		log.Fatalf("ERROR CREATING USER %s", err)
	}

// 	AddUser := repositories.NewAddUser(gateways.GetJson(&models.AddUser{}))
// 	repositories.CreateUser(AddUser, db)
// 	ListUser := repositories.NewListUser(repositories.ListUser(&models.ListUser{}, db))
// 	fmt.Println(ListUser)
// 	fmt.Println(services.New(gateways.GetJson(&models.AddUser{})))
// 	telegram.Telegram()
// }

func initDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
