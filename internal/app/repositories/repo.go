package repositories

import (
	"database/sql"
	"log"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
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

func (r *SymbolRepository) CreateOrder(a futures.Order) error {

	_, err := r.db.Exec("insert into FutureOrder(Symbol, OrderID, ClientOrderID, Price, ReduceOnly, OrigQuantity,ExecutedQuantity, CumQuantity, CumQuote, Status, TimeInForce, Type, Side, StopPrice,	Time, UpdateTime, WorkingType, ActivatePrice, PriceRate, AvgPrice, OrigType,PositionSide, PriceProtect, ClosePosition) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24)",
		a.Symbol, a.OrderID, a.ClientOrderID, a.Price, a.ReduceOnly, a.OrigQuantity,
		a.ExecutedQuantity, a.CumQuantity, a.CumQuote, a.Status, a.TimeInForce, a.Type, a.Side, a.StopPrice,
		a.Time, a.UpdateTime, a.WorkingType, a.ActivatePrice, a.PriceRate, a.AvgPrice, a.OrigType,
		a.PositionSide, a.PriceProtect, a.ClosePosition,
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
	sendPrices := []binance.SymbolPrice{}
	for rows.Next() {
		err := rows.Scan(&m.Symbol, &m.Price)
		if err != nil {
			log.Fatalf("ERROR %s", err)
			continue
		}
		sendPrices = append(sendPrices, m)
	}
	return sendPrices
}

func (r *SymbolRepository) ListOrder() []futures.Order {
	var m futures.Order
	rows, err := r.db.Query("SELECT * FROM FutureOrder ORDER BY OrderID")
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}
	sendOrders := []futures.Order{}
	for rows.Next() {
		err := rows.Scan(&m)
		if err != nil {
			log.Fatalf("ERROR %s", err)
			continue
		}
		sendOrders = append(sendOrders, m)
	}
	return sendOrders
}
