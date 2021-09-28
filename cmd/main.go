package main

import (
	"database/sql"
	"fmt"
	"homework/config"
	"homework/internal/app/repositories"
	"homework/internal/app/telegram"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

func main() {

	db, err := initDb()
	fmt.Println("connect postgre")
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}

	Repo := repositories.NewUserRepository(db)

	bot, err := tgbotapi.NewBotAPI(config.TeleToken)
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}
	bot.Debug = true

	telegramBot := telegram.NewBot(bot, Repo)

	// go func() {
	if err := telegramBot.Start(); err != nil {
		log.Fatalf("ERROR %s", err)
	}
	// }()

	// var u models.User

	// err = Repo.CreateUser(u)
	// if err != nil {
	// 	log.Fatalf("ERROR CREATING USER %s", err)
	// }
}

// 	a, err := userRepo.FindById(5)
// 	if err != nil {
// 		log.Fatalf("ERROR %s", err)
// 	}
// 	fmt.Println(a)
// }

func initDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		log.Fatalf("ERROR CREATING USER %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("ERROR CREATING USER %s", err)
	}
	return db, err
}

// u.SetName("name", "surname")
// firstName := u.GetFirstName()
// // чтение из json
