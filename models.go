package bingxgo

import "fmt"

type BingXResponse[T any] struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	DebugMsg string `json:"debugMsg"`
	Data     T      `json:"data"`
}

func (resp BingXResponse[T]) Error() error {
	if resp.Code != 0 {
		return fmt.Errorf("code: %d, msg: %s, debugMsg: %s", resp.Code, resp.Msg, resp.DebugMsg)
	}
	return nil
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
	Symbol              string `json:"symbol"`
	OrderId             int64  `json:"orderId"`
	TransactTime        int64  `json:"transactTime"`
	Price               string `json:"price"`
	StopPrice           string `json:"stopPrice"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	ClientOrderID       string `json:"clientOrderID"`
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

type OrderBook struct {
	// The timestamp of when the orderbook last changed (in milliseconds)
	Timestamp int64 `json:"ts,omitempty"`
	// Asks order depth
	Asks [][]string `json:"asks"`
	// Bids order depth
	Bids [][]string `json:"bids"`
}

type SymbolInfos struct {
	Symbols []SymbolInfo `json:"symbols"`
}

type SymbolInfo struct {
	Symbol       string  `json:"symbol"`       // 交易对符号
	TickSize     float64 `json:"tickSize"`     // 最小价格变动单位
	StepSize     float64 `json:"stepSize"`     // 最小交易单位
	MinQty       float64 `json:"minQty"`       // 最小下单量
	MaxQty       float64 `json:"maxQty"`       // 最大下单量
	MinNotional  float64 `json:"minNotional"`  // 最小名义价值
	MaxNotional  float64 `json:"maxNotional"`  // 最大名义价值
	Status       int     `json:"status"`       // 状态标识
	ApiStateBuy  bool    `json:"apiStateBuy"`  // 买入API状态
	ApiStateSell bool    `json:"apiStateSell"` // 卖出API状态
	TimeOnline   int64   `json:"timeOnline"`   // 上线时间 (Unix时间戳)
	OffTime      int64   `json:"offTime"`      // 下线时间 (Unix时间戳)
	MaintainTime int64   `json:"maintainTime"` // 维护时间 (秒)
}

type Ticker struct {
	Symbol string  `json:"symbol"`
	Trades []Trade `json:"trades"`
}

type Trade struct {
	Timestamp int64  `json:"timestamp"`
	TradeId   string `json:"tradeId"`
	Price     string `json:"price"`
	Amount    string `json:"amount"`
	Type      int    `json:"type"`
	Volume    string `json:"volume"`
}

//	{
//	    "amount": "49999.00000000000000000000",
//	    "coin": "USDTTRC20",
//	    "network": "TRC20",
//	    "status": 1,
//	    "address": "TP******B4v",
//	    "addressTag": "",
//	    "txId": "60*****1d",
//	    "insertTime": 1701557778000,
//	    "unlockConfirm": "2/2",
//	    "confirmTimes": "2/2"
//	  }
type DepositRecord struct {
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	TxId          string `json:"txId"`
	InsertTime    int64  `json:"insertTime"`
	UnlockConfirm string `json:"unlockConfirm"`
	ConfirmTimes  string `json:"confirmTimes"`
}

// [
//
//	{
//	  "address": "TR****zc",
//	  "amount": "3500.00000000000000000000",
//	  "applyTime": "2023-12-14T04:05:02.000+08:00",
//	  "coin": "USDTTRC20",
//	  "id": "125*****98",
//	  "network": "TRC20",
//	  "transferType": 1,
//	  "transactionFee": "1.00000000000000000000",
//	  "confirmNo": 2,
//	  "info": "",
//	  "txId": "b9***********b67"
//	}
//
// ]
type WithdrawRecord struct {
	Address        string `json:"address"`
	Amount         string `json:"amount"`
	ApplyTime      string `json:"applyTime"`
	Coin           string `json:"coin"`
	Id             string `json:"id"`
	Network        string `json:"network"`
	TransferType   int    `json:"transferType"`
	TransactionFee string `json:"transactionFee"`
	ConfirmNo      int    `json:"confirmNo"`
	Info           string `json:"info"`
	TxId           string `json:"txId"`
}
