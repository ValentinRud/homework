package repositories

import (
	"database/sql"
	"log"

	"github.com/adshao/go-binance/v2"
	_ "github.com/lib/pq"
)

type SymbolRepository struct {
	db *sql.DB
}

func NewSymbolRepository(db *sql.DB) *SymbolRepository {
	return &SymbolRepository{
		db: db,
	}
}

func (r *SymbolRepository) CreatePrice(a binance.SymbolPrice) error {
	_, err := r.db.Exec("insert into SymbolPrice(symbol, price) values ($1,$2)",
		a.Symbol, a.Price,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *SymbolRepository) ListSymbol() []binance.SymbolPrice {
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
