package services

import "homework/internal/app/models"

func New(Id int, First_name string, Last_name string, Age int, Status string) *models.AddUser {
	return &models.AddUser{
		Id:        Id,
		FirstName: First_name,
		LastName:  Last_name,
		Age:       Age,
		Status:    Status,
	}
}
