package main

import (
	"fmt"
	"homework/internal/app/gateways"
	"homework/internal/app/models"
	"homework/internal/app/repositories"
	"homework/internal/app/services"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := repositories.NewPostgresDb()
	fmt.Println("connect postgre")
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}

	AddUser := repositories.NewAddUser(gateways.GetJson(&models.AddUser{}))
	repositories.CreateUser(AddUser, db)
	repositories.NewListUser(repositories.ListUser(&models.ListUser{}, db))
	fmt.Println(services.New(gateways.GetJson(&models.AddUser{})))

}
