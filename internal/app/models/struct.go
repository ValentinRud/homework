package models

import (
	"database/sql"
)

type AddUser struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Status    string `json:"status"`
}

type ListUser struct {
	Id       int    `json:"id"`
	LastName string `json:"lastName"`
}

type DataBase struct {
	db *sql.DB
}
