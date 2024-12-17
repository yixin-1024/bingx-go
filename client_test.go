package bingxgo

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	client     = &Client{}
	symbol     = "THG-USDT"
	spotClient = &SpotClient{}
)

func init() {
	client = NewClient(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))
	client.Debug = true
	spotClient.client = client
}

func TestBalance(t *testing.T) {
	balances, err := spotClient.GetBalance()
	assert.Equal(t, err, nil)
	s, _ := json.MarshalIndent(balances, "", "\t")
	t.Log(string(s))
}

func TestBatchOrders(t *testing.T) {
	quantity := 100.0
	price := 0.01584

	orders, err := spotClient.CreateBatchOrders([]SpotOrderRequest{
		{
			Symbol:      symbol,
			Side:        "SELL",
			Type:        "LIMIT",
			Quantity:    quantity,
			Price:       price,
			TimeInForce: "GTC",
		},
		{
			Symbol:      symbol,
			Side:        "BUY",
			Type:        "LIMIT",
			Quantity:    quantity,
			Price:       price,
			TimeInForce: "FOK",
		},
	}, true)

	assert.Equal(t, err, nil)
	t.Log(orders)
}

func TestCreateOrder(t *testing.T) {
	order, err := spotClient.CreateOrder(SpotOrderRequest{
		Symbol:      symbol,
		Side:        "SELL",
		Type:        "LIMIT",
		Quantity:    50,
		Price:       0.05,
		TimeInForce: "GTC",
	})
	assert.Equal(t, err, nil)
	t.Log(order)
}

func TestOpenOrders(t *testing.T) {
	orders, err := spotClient.GetOpenOrders(symbol)
	assert.Equal(t, err, nil)

	t.Log(orders)
}
func TestCancelOrder(t *testing.T) {
	err := spotClient.CancelOrder(symbol, "1861312320206962688")
	assert.Equal(t, err, nil)
}

func TestCancelAllOpenOrders(t *testing.T) {
	err := spotClient.CancelAllOpenOrders(symbol)
	assert.Equal(t, err, nil)
}

func TestHistoryOrders(t *testing.T) {
	orders, err := spotClient.HistoryOrders(symbol)
	assert.Equal(t, err, nil)

	sort.Slice(orders, func(i, j int) bool {
		return orders[i].Time > orders[j].Time
	})
	s, _ := json.MarshalIndent(orders, "", "\t")
	t.Log(string(s))

}

func TestOderBook(t *testing.T) {
	book, err := spotClient.OrderBook(symbol, 10)
	assert.Equal(t, err, nil)
	t.Log(book)
	t.Log(book.Asks[len(book.Asks)-1][0])
	t.Log(book.Bids[0][0])
}

func TestGetSymbolInfo(t *testing.T) {
	symbolInfo, err := spotClient.GetSymbolInfo("SOL-USDT")
	assert.Equal(t, err, nil)
	t.Log(symbolInfo)
	log.Fatal(symbolInfo)
}
