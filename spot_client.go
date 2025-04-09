package bingxgo

import (
	"encoding/json"
	"strconv"
	"time"
)

type SpotClient struct {
	client *Client
}

func NewSpotClient(client *Client) SpotClient {
	return SpotClient{client: client}
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

	var bingXResponse BingXResponse[map[string][]SpotBalance]
	err = json.Unmarshal(resp, &bingXResponse)
	if err != nil {
		return nil, err
	}
	if err := bingXResponse.Error(); err != nil {
		return nil, err
	}
	return bingXResponse.Data["balances"], err
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

	var bingXResponse BingXResponse[SpotOrderResponse]
	err = json.Unmarshal(resp, &bingXResponse)
	if err != nil {
		return nil, err
	}
	if err := bingXResponse.Error(); err != nil {
		return nil, err
	}
	return &bingXResponse.Data, err
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

	var bingXResponse BingXResponse[map[string][]SpotOrderResponse]
	err = json.Unmarshal(resp, &bingXResponse)
	if err != nil {
		return nil, err
	}
	if err := bingXResponse.Error(); err != nil {
		return nil, err
	}
	return bingXResponse.Data["orders"], err
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

	var bingXResponse BingXResponse[map[string][]SpotOrder]
	err = json.Unmarshal(resp, &bingXResponse)
	if err != nil {
		return nil, err
	}
	if err := bingXResponse.Error(); err != nil {
		return nil, err
	}
	return bingXResponse.Data["orders"], err
}

func (c *SpotClient) CancelOrder(symbol string, orderId string) error {
	endpoint := "/openApi/spot/v1/trade/cancel"
	params := map[string]interface{}{
		"symbol":  symbol,
		"orderId": orderId,
	}

	resp, err := c.client.sendRequest("POST", endpoint, params)
	if err != nil {
		return err
	}
	var bingXResponse BingXResponse[any]
	err = json.Unmarshal(resp, &bingXResponse)
	if err != nil {
		return err
	}
	if err := bingXResponse.Error(); err != nil {
		return err
	}
	return nil
}

func (c *SpotClient) CancelAllOpenOrders(symbol string) error {
	endpoint := "/openApi/spot/v1/trade/cancelOpenOrders"
	params := map[string]interface{}{
		"symbol": symbol,
	}

	resp, err := c.client.sendRequest("POST", endpoint, params)
	if err != nil {
		return err
	}
	var bingXResponse BingXResponse[any]
	err = json.Unmarshal(resp, &bingXResponse)
	if err != nil {
		return err
	}
	if err := bingXResponse.Error(); err != nil {
		return err
	}
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

	var bingXResponse BingXResponse[SpotOrder]
	err = json.Unmarshal(resp, &bingXResponse)
	if err != nil {
		return nil, err
	}
	if err := bingXResponse.Error(); err != nil {
		return nil, err
	}
	return &bingXResponse.Data, err
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

	var bingXResponse BingXResponse[map[string][]SpotOrder]
	err = json.Unmarshal(resp, &bingXResponse)
	if err != nil {
		return nil, err
	}
	if err := bingXResponse.Error(); err != nil {
		return nil, err
	}
	return bingXResponse.Data["orders"], err
}

func (c *SpotClient) OrderBook(symbol string, limit int) (*OrderBook, error) {
	endpoint := "/openApi/spot/v1/market/depth"
	params := map[string]interface{}{
		"symbol": symbol,
	}
	if limit > 0 {
		params["limit"] = limit
	}

	resp, err := c.client.sendRequest("GET", endpoint, params)
	if err != nil {
		return nil, err
	}

	var bingXResponse BingXResponse[OrderBook]
	err = json.Unmarshal(resp, &bingXResponse)
	if err != nil {
		return nil, err
	}
	if err := bingXResponse.Error(); err != nil {
		return nil, err
	}
	return &bingXResponse.Data, err
}

func (c *SpotClient) GetSymbolInfo(symbol string) (*SymbolInfo, error) {
	endpoint := "/openApi/spot/v1/common/symbols"
	params := map[string]interface{}{
		"symbol": symbol,
	}

	resp, err := c.client.sendRequest("GET", endpoint, params)
	if err != nil {
		return nil, err
	}

	var bingXResponse BingXResponse[SymbolInfos]
	err = json.Unmarshal(resp, &bingXResponse)
	if err != nil {
		return nil, err
	}
	if err := bingXResponse.Error(); err != nil {
		return nil, err
	}
	return &bingXResponse.Data.Symbols[0], err
}

func (c *SpotClient) GetTickers(symbol string) (*Ticker, error) {
	endpoint := "/openApi/spot/v1/ticker/price"
	params := map[string]interface{}{}
	if symbol != "" {
		params["symbol"] = symbol
	}

	resp, err := c.client.sendRequest("GET", endpoint, params)
	if err != nil {
		return nil, err
	}

	var bingXResponse BingXResponse[Ticker]
	err = json.Unmarshal(resp, &bingXResponse)
	if err != nil {
		return nil, err
	}
	if err := bingXResponse.Error(); err != nil {
		return nil, err
	}
	return &bingXResponse.Data, err
}
