package bingxgo

import (
	"encoding/json"
	"strconv"
	"time"
)

type SpotClient struct {
	client *Client
}

func (c *SpotClient) GetBalance() ([]SpotBalance, error) {
	endpoint := "/openApi/spot/v1/account/balance"
	params := map[string]interface{}{
		"timestamp": time.Now().UnixMilli(),
	}

	resp, err := c.client.sendRequest("GET", endpoint, params)
	if err != nil {
		return nil, err
	}

	var balance BingXResponse[map[string][]SpotBalance]
	err = json.Unmarshal(resp, &balance)
	return balance.Data["balances"], err
}

func (c *SpotClient) CreateOrder(order SpotOrderRequest) (*SpotOrderResponse, error) {
	endpoint := "/openApi/spot/v1/trade/order"
	params := map[string]interface{}{
		"symbol":   order.Symbol,
		"side":     string(order.Side),
		"type":     string(order.Type),
		"quantity": strconv.FormatFloat(order.Quantity, 'f', -1, 64),
		"price":    strconv.FormatFloat(order.Price, 'f', -1, 64),
	}

	resp, err := c.client.sendRequest("POST", endpoint, params)
	if err != nil {
		return nil, err
	}

	var orderResp BingXResponse[SpotOrderResponse]
	err = json.Unmarshal(resp, &orderResp)
	return &orderResp.Data, err
}

func (c *SpotClient) CreateBatchOrders(orders []SpotOrderRequest, isSync bool) ([]SpotOrderResponse, error) {
	endpoint := "/openApi/spot/v1/trade/batchOrders"

	ordersJSON, err := json.Marshal(orders)
	if err != nil {
		return nil, err
	}
	params := map[string]interface{}{
		"data": string(ordersJSON),
		"sync": isSync,
	}

	resp, err := c.client.sendRequest("POST", endpoint, params)
	if err != nil {
		return nil, err
	}

	var batchResp BingXResponse[map[string][]SpotOrderResponse]
	err = json.Unmarshal(resp, &batchResp)
	return batchResp.Data["orders"], err
}

func (c *SpotClient) GetOpenOrders(symbol string) ([]SpotOrder, error) {
	endpoint := "/openApi/spot/v1/trade/openOrders"
	params := map[string]interface{}{
		"symbol": symbol,
	}

	resp, err := c.client.sendRequest("GET", endpoint, params)
	if err != nil {
		return nil, err
	}

	var orders BingXResponse[map[string][]SpotOrder]
	err = json.Unmarshal(resp, &orders)
	return orders.Data["orders"], err
}

func (c *SpotClient) CancelOrder(symbol string, orderId string) error {
	endpoint := "/openApi/spot/v1/trade/cancel"
	params := map[string]interface{}{
		"symbol":  symbol,
		"orderId": orderId,
	}

	_, err := c.client.sendRequest("POST", endpoint, params)
	return err
}

func (c *SpotClient) CancelAllOpenOrders(symbol string) error {
	endpoint := "/openApi/spot/v1/trade/cancelOpenOrders"
	params := map[string]interface{}{
		"symbol": symbol,
	}

	_, err := c.client.sendRequest("POST", endpoint, params)
	return err
}

func (c *SpotClient) GetOrder(symbol string, orderId string) (*SpotOrder, error) {
	endpoint := "/openApi/spot/v1/trade/order"
	params := map[string]interface{}{
		"symbol":  symbol,
		"orderId": orderId,
	}

	resp, err := c.client.sendRequest("GET", endpoint, params)
	if err != nil {
		return nil, err
	}

	var order BingXResponse[SpotOrder]
	err = json.Unmarshal(resp, &order)
	return &order.Data, err
}

func (c *SpotClient) HistoryOrders(symbol string) ([]SpotOrder, error) {
	endpoint := "/openApi/spot/v1/trade/historyOrders"
	params := map[string]interface{}{
		"symbol": symbol,
	}

	resp, err := c.client.sendRequest("GET", endpoint, params)
	if err != nil {
		return nil, err
	}

	var orders BingXResponse[map[string][]SpotOrder]
	err = json.Unmarshal(resp, &orders)
	return orders.Data["orders"], err
}
