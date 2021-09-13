package repositories

import (
	"database/sql"
	"fmt"
	"homework/config"
	"homework/internal/app/models"

	_ "github.com/lib/pq"
)

func CreateUser(u *models.AddUser) {
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	result, err := db.Query("insert into test (id, first_name, last_name, age, status) values ($1,$2,$3,$4,$5)", u.Id, u.FirstName, u.LastName, u.Age, u.Status)
	if err != nil {
		panic(err)
	}
	defer result.Close()

}

var s models.SeeUser

func SelectDb() (int, string) {
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, last_name FROM test ORDER BY id ASC")
	if err != nil {
		panic(err)
	}
	sendUsers := []models.SeeUser{}

	for rows.Next() {
		err := rows.Scan(&s.Id, &s.LastName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sendUsers = append(sendUsers, s)
	}

	for _, s := range sendUsers {
		fmt.Println(s.Id, s.LastName)
	}
	return s.Id, s.LastName
}
