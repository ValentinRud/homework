package main

import (
	"fmt"
	"homework/internal/app/gateways"
	"homework/internal/app/repositories"
	"homework/internal/app/services"
)

func main() {
	gateways.GetJson()
	fmt.Println(services.New(gateways.M.Id, gateways.M.LastName, gateways.M.FirstName, gateways.M.Age, gateways.M.Status))
	repositories.CreateUser(gateways.M)
	repositories.SelectDb()
}
