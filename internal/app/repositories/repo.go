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

type DataBase struct {
	db *sql.DB
}

func NewDataBase(db *sql.DB) *DataBase {
	return &DataBase{
		db: db,
	}
}

type AddUser interface {
}

type SeeUser interface {
}
type Repository struct {
	AddUser
	SeeUser
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}

func NewPostgresDb(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", cfg.User, cfg.Password, cfg.Dbname, cfg.Sslmode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}

func CreateUser(u *models.AddUser, d *sql.DB) {
	result, err := d.Query("insert into test (id, first_name, last_name, age, status) values ($1,$2,$3,$4,$5)", u.Id, u.FirstName, u.LastName, u.Age, u.Status)
	if err != nil {
		panic(err)
	}
	defer result.Close()

}

var s models.ListUser

func ListUser(d *DataBase) (int, string) {
	rows, err := d.db.Query("SELECT id, last_name FROM test ORDER BY id ASC")
	if err != nil {
		panic(err)
	}
	sendUsers := []models.ListUser{}
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
