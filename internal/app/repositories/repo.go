package repositories

import (
	"database/sql"
	"fmt"
	"homework/internal/app/models"

	_ "github.com/lib/pq"
)

//type Repository interface {
//	CreateUser()
//	ListUser()
//}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func NewAddUserRepo(Id int, First_name string, Last_name string, Age int, Status string) *models.AddUser {
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

func (repo *UserRepository) CreateUser(a models.User) error {
	// Query для запросов, которые что-то возвращает
	// Exec для запросов, которые ничего не возвращают

	_, err := repo.db.Exec("insert into test (id, first_name, last_name, age, status) values ($1,$2,$3,$4,$5)",
		a.ID, a.FirstName, a.LastName, a.Age, a.Status,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) ListUser(l *models.ListUser) (int, string) {
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

// пишешь в телеге id, тебе выводятся в ТЕЛЕГЕ все данные по этому пользователю
// GetByID(id int)
