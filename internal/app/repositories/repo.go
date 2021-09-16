package repositories

import (
	"database/sql"
	"fmt"
	"homework/config"
	"homework/internal/app/models"

	_ "github.com/lib/pq"
)

//type Repository interface {
//	CreateUser()
//	ListUser()
//}

func NewAddUser(Id int, First_name string, Last_name string, Age int, Status string) *models.AddUser {
	return &models.AddUser{
		Id:        Id,
		FirstName: First_name,
		LastName:  Last_name,
		Age:       Age,
		Status:    Status,
	}
}

func NewListUser(Id int, Last_name string) *models.ListUser {
	return &models.ListUser{
		Id:       Id,
		LastName: Last_name,
	}
}

func NewPostgresDb() (*sql.DB, error) {
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

func CreateUser(a *models.AddUser, s *sql.DB) {
	result, err := s.Query("insert into test (id, first_name, last_name, age, status) values ($1,$2,$3,$4,$5)", a.Id, a.FirstName, a.LastName, a.Age, a.Status)
	if err != nil {
		panic(err)
	}
	defer result.Close()
}

func ListUser(l *models.ListUser, s *sql.DB) (int, string) {
	rows, err := s.Query("SELECT id, last_name FROM test ORDER BY id ASC")
	if err != nil {
		panic(err)
	}
	sendUsers := []models.ListUser{}
	for rows.Next() {
		err := rows.Scan(&l.Id, &l.LastName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sendUsers = append(sendUsers, *l)
	}

	for _, l := range sendUsers {
		fmt.Println(l.Id, l.LastName)
	}
	return l.Id, l.LastName
}
