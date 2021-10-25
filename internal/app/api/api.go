package api

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

var (
	apiKey    = "qdAOdQVLhdJAiDuZFJETN0Nv8C5yFfycGePV8qw7oaPh5GyxIAf3qQVrOLO4EbHK"
	secretKey = "vht9WH2eSqACWiJa5xt5lojW03vLhGqEntetujbGCJCRb3mBoOLg9VIgDlfFrFSn"
)

type ClientApi struct {
	client *binance.Client
}

func NewClientApi() *ClientApi {
	//clients := binance.NewClient(apiKey, secretKey)
	// futuresClient := binance.NewFuturesClient(apiKey, secretKey)    // USDT-M Futures
	// deliveryClient := binance.NewDeliveryClient(apiKey, secretKey)  // Coin-M Futures
	return &ClientApi{
		client: binance.NewClient(apiKey, secretKey),
	}
}

func (c *ClientApi) ListOrders() []*binance.Order {
	orders, err := c.client.NewListOrdersService().Symbol("DTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return orders
	}
	return nil
}

func (c *ClientApi) ListTicket() []*binance.SymbolPrice {
	prices, err := c.client.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return prices
}

//client := binance.NewClient(apiKey, secretKey)
//futuresClient := binance.NewFuturesClient(apiKey, secretKey)    // USDT-M Futures
//deliveryClient := binance.NewDeliveryClient(apiKey, secretKey)  // Coin-M Futures
