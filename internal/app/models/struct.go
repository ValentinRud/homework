package models

import (
	"database/sql"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Status    string `json:"status"`
}

type DataBase struct {
	db *sql.DB
}

type BotMessage struct {
	Message struct {
		MessageId int
		From      struct {
			UserName string
			Id       int
		}
		Chat struct {
			Id int
		}
		Text string
	}
}
