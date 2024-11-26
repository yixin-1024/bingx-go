package bingxgo

type BingXResponse[T any] struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	DebugMsg string `json:"debugMsg"`
	Data     T      `json:"data"`
}

type SpotOrderRequest struct {
	Symbol      string  `json:"symbol"`
	Side        string  `json:"side"` // BUY, SELL
	Type        string  `json:"type"` // LIMIT, MARKET
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price,omitempty"`
	TimeInForce string  `json:"timeInForce,omitempty"` // GTC, IOC, FOK
}

type SpotOrderResponse struct {
	OrderId       int    `json:"orderId"`
	Symbol        string `json:"symbol"`
	Status        string `json:"status"`
	ClientOrderId string `json:"clientOrderId"`
}

type SpotOrder struct {
	OrderId     int     `json:"orderId"`
	Symbol      string  `json:"symbol"`
	Price       string  `json:"price"`
	OrigQty     string  `json:"origQty"`
	ExecutedQty string  `json:"executedQty"`
	Status      string  `json:"status"`
	Type        string  `json:"type"`
	Side        string  `json:"side"`
	Time        int64   `json:"time"`
	Fee         float64 `json:"fee"`
	AvgPrice    float64 `json:"avgPrice"`
}

type SpotBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type Balance struct {
	Available float64 `json:"available"`
	Locked    float64 `json:"locked"`
}

type Kline struct {
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
	Time   string  `json:"time"`
}

type OrderRequest struct {
	Symbol       string  `json:"symbol"`
	Side         string  `json:"side"`         // BUY, SELL
	PositionSide string  `json:"positionSide"` // LONG, SHORT
	Type         string  `json:"type"`         // LIMIT, MARKET
	Quantity     float64 `json:"quantity"`
	Price        float64 `json:"price"`
}

type OrderResponse struct {
	OrderId       int    `json:"orderId"`
	Symbol        string `json:"symbol"`
	Status        string `json:"status"`
	ClientOrderId string `json:"clientOrderId"`
}
