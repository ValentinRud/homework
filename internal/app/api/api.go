package api

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/delivery"
	"github.com/adshao/go-binance/v2/futures"
)

type ClientApi struct {
	client         *binance.Client
	futuresClient  *futures.Client
	deliveryClient *delivery.Client
}

func NewClientApi(apiKey string, secretKey string) *ClientApi {
	//clients := binance.NewClient(apiKey, secretKey)
	//  futuresClient := binance.NewFuturesClient(apiKey, secretKey)    // USDT-M Futures
	//  deliveryClient := binance.NewDeliveryClient(apiKey, secretKey)  // Coin-M Futures
	return &ClientApi{
		client:         binance.NewClient(apiKey, secretKey),
		futuresClient:  binance.NewFuturesClient(apiKey, secretKey),
		deliveryClient: binance.NewDeliveryClient(apiKey, secretKey),
	}
}

func (c *ClientApi) ListOrders() []*futures.Order {

	/*xxx, err := c.client.NewListOrdersService().Symbol("IOTAUSDT").Do(context.Background())
	yyy, err := c.deliveryClient.NewListOrdersService().Symbol("IOTAUSDT").Do(context.Background())

	fmt.Sprintln(xxx)
	fmt.Sprintln(yyy)*/

	orders, err := c.futuresClient.NewListOrdersService(). /*Symbol("IOTAUSDT").*/
								Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return orders
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
