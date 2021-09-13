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

type DataBase struct {
	db *sql.DB
}

func (d *DataBase) Open() error {
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	d.db = db

	return err
}

func CreateUser(u *models.AddUser, d DataBase) {
	d.Open()
	result, err := d.db.Query("insert into test (id, first_name, last_name, age, status) values ($1,$2,$3,$4,$5)", u.Id, u.FirstName, u.LastName, u.Age, u.Status)
	if err != nil {
		panic(err)
	}
	defer result.Close()

}

var s models.SeeUser

func ListUser(d *DataBase) (int, string) {
	d.Open()
	rows, err := d.db.Query("SELECT id, last_name FROM test ORDER BY id ASC")
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
