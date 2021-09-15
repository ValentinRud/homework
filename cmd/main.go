package main

import (
	"fmt"
	"homework/config"
	"homework/internal/app/gateways"
	"homework/internal/app/repositories"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := repositories.NewPostgresDb(config.Config{
		User:     "postgres",
		Password: "Qweasdzxc1",
		Dbname:   "users",
		Sslmode:  "disable",
	})
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}

	repos := repositories.NewRepository(db)
	fmt.Println(repos)
	repositories.CreateUser(gateways.M, db)
	repositories.ListUser(repositories.NewDataBase(db))
}
