package repositories

import (
	"database/sql"
	"homework/internal/app/models"
	"log"

	"github.com/adshao/go-binance/v2"
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

func (r *UserRepository) CreatePrice(a binance.SymbolPrice) error {
	// Query для запросов, которые что-то возвращает
	// Exec для запросов, которые ничего не возвращают
	_, err := r.db.Exec("insert into SymbolPrice(symbol, price) values ($1,$2)",
		a.Symbol, a.Price,
	)
	if err != nil {
		return err
	}

	return nil
}

// func (r *UserRepository) ListUsers(m *models.User) models.User {
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
// 		return m
// 	}
// 	return models.User{}
// }
// выводить по пять, break
func (r *UserRepository) ListSymbol() []binance.SymbolPrice {
	var m binance.SymbolPrice
	rows, err := r.db.Query("SELECT * FROM SymbolPrice ORDER BY symbol")
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}
	sendUsers := []binance.SymbolPrice{}
	for rows.Next() {
		err := rows.Scan(&m.Symbol, &m.Price)
		if err != nil {
			log.Fatalf("ERROR %s", err)
			continue
		}
		sendUsers = append(sendUsers, m)
	}
	return sendUsers
}

func (r *UserRepository) FindById(id int) (*models.User, error) {
	u := &models.User{}
	if err := r.db.QueryRow("SELECT first_name, last_name FROM test WHERE id=1$",
		id).Scan(u.ID, u.LastName); err != nil {
		return nil, err
	}
	return u, nil
}

// func (r *UserRepository) GetByID(id int) {
// 	row, err := r.db.Query("SELECT id, last_name FROM test Where id=1$", id)
// 	if err != nil {
// 		log.Fatalf("ERROR %s", err)
// 	}
// }

// пишешь в телеге id, тебе выводятся в ТЕЛЕГЕ все данные по этому пользователю
// GetByID(id int)

// 	AddUser := repositories.NewAddUser(gateways.GetJson(&models.AddUser{}))
// 	repositories.CreateUser(AddUser, db)
// 	ListUser := repositories.NewListUser(repositories.ListUser(&models.ListUser{}, db))
// 	fmt.Println(ListUser)
// 	fmt.Println(services.New(gateways.GetJson(&models.AddUser{})))
// 	telegram.Telegram()
// }
