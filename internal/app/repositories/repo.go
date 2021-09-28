package repositories

import (
	"database/sql"
	"fmt"
	"homework/internal/app/models"
	"log"

	_ "github.com/lib/pq"
)

type Repository interface {
	CreateUser(models.User) error
	ListUser(Id int) (string, error)
}

type UserRepository struct {
	db *sql.DB
}

// func (r *UserRepository) CreateUser(models.User) error

// func (r *UserRepository) ListUser(Id int) (string,error)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func NewAddUserRepo(Id int, First_name string, Last_name string, Age int, Status string) *models.User {
	return &models.User{
		ID:        Id,
		FirstName: First_name,
		LastName:  Last_name,
		Age:       Age,
		Status:    Status,
	}
}

func NewListUser(Id int, Last_name string) *models.User {
	return &models.User{
		ID:       Id,
		LastName: Last_name,
	}
}

func (r *UserRepository) CreateUser(a models.User) error {
	// Query для запросов, которые что-то возвращает
	// Exec для запросов, которые ничего не возвращают
	_, err := r.db.Exec("insert into test (id, first_name, last_name, age, status) values ($1,$2,$3,$4,$5)",
		a.ID, a.FirstName, a.LastName, a.Age, a.Status,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) ListUsers(m models.User) string {
	rows, err := r.db.Query("SELECT id, last_name FROM test ORDER BY id ASC")
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}
	sendUsers := []models.User{}
	for rows.Next() {
		err := rows.Scan(m.ID, m.LastName)
		if err != nil {
			log.Fatalf("ERROR %s", err)
			continue
		}
		sendUsers = append(sendUsers, m)
	}

	for _, m := range sendUsers {
		return fmt.Sprint(m)
	}
	return fmt.Sprint(m)
}

// func (r *UserRepository) ListUser(m *models.User) (int, string) {
// 	rows, err := r.db.Query("SELECT id, last_name FROM test ORDER BY id ASC")
// 	if err != nil {
// 		log.Fatalf("ERROR %s", err)
// 	}
// 	sendUsers := []models.User{}
// 	for rows.Next() {
// 		err := rows.Scan(&m.ID, &m.LastName)
// 		if err != nil {
// 			log.Fatalf("ERROR %s", err)
// 			continue
// 		}
// 		sendUsers = append(sendUsers, *m)
// 	}

// 	for _, m := range sendUsers {
// 		return m.ID, m.LastName
// 	}
// 	return m.ID, m.LastName
// }

func (r *UserRepository) FindById(ID int) (*models.User, error) {
	u := &models.User{}
	if err := r.db.QueryRow("SELECT first_name, last_name FROM test WHERE id=1$",
		ID).Scan(u.ID, u.LastName); err != nil {
		return nil, err
	}
	return u, nil
}

// пишешь в телеге id, тебе выводятся в ТЕЛЕГЕ все данные по этому пользователю
// GetByID(id int)

// 	AddUser := repositories.NewAddUser(gateways.GetJson(&models.AddUser{}))
// 	repositories.CreateUser(AddUser, db)
// 	ListUser := repositories.NewListUser(repositories.ListUser(&models.ListUser{}, db))
// 	fmt.Println(ListUser)
// 	fmt.Println(services.New(gateways.GetJson(&models.AddUser{})))
// 	telegram.Telegram()
// }
