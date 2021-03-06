package binance

type ExchangeInfoFilter struct {
	FilterType string `json:"filterType"`

	// for PRICE_FILTER
	MinPrice float64 `json:"minPrice,string,omitempty"`
	MaxPrice float64 `json:"maxPrice,string,omitempty"`
	TickSize float64 `json:"tickSize,string,omitempty"`

	// for PERCENT_PRICE
	MultiplierUp   float64 `json:"multiplierUp,string,omitempty"`
	MultiplierDown float64 `json:"multiplierDown,string,omitempty"`
	AvgPriceMins   int     `json:"avgPriceMins,omitempty"` // also for MIN_NOTIONAL

	// for LOT_SIZE and MARKET_LOT_SIZE
	MinQty   float64 `json:"minQty,string,omitempty"`
	MaxQty   float64 `json:"maxQty,string,omitempty"`
	StepSize float64 `json:"stepSize,string,omitempty"`

	// for MIN_NOTIONAL
	MinNotional   float64 `json:"minNotional,string,omitempty"`
	ApplyToMarket bool    `json:"applyToMarket,omitempty"`

	// for ICEBERG_PARTS
	Limit int `json:"limit,omitempty"`

	// for MAX_NUM_ORDERS
	MaxNumOrders int `json:"maxNumOrders,omitempty"`

	// for MAX_NUM_ALGO_ORDERS
	MaxNumAlgoOrders int `json:"maxNumAlgoOrders,omitempty"`
}

type ExchangeInfo struct {
	Symbol                     string                `json:"symbol"`
	Status                     string                `json:"status"`
	BaseAsset                  string                `json:"baseAsset"`
	BaseAssetPrecision         int                   `json:"baseAssetPrecision"`
	QuoteAsset                 string                `json:"quoteAsset"`
	QuotePrecision             int                   `json:"quotePrecision"`
	QuoteAssetPrecision        int                   `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int                   `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int                   `json:"quoteCommissionPrecision"`
	OrderTypes                 []string              `json:"orderTypes"`
	IcebergAllowed             bool                  `json:"icebergAllowed"`
	OcoAllowed                 bool                  `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed bool                  `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop          bool                  `json:"allowTrailingStop"`
	IsSpotTradingAllowed       bool                  `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool                  `json:"isMarginTradingAllowed"`
	Filters                    []*ExchangeInfoFilter `json:"filters"`
	Permissions                []string              `json:"permissions"`
}

type ExchangeInfoMsg struct {
	// other fields are not used, wait for later complemention
	ExchangeInfoList []*ExchangeInfo `json:"symbols"`
}

type OrderFill struct {
	Price           float64 `json:"price,string"`
	Qty             float64 `json:"qty,string"`
	Commission      float64 `json:"commission,string"`
	CommissionAsset string  `json:"commissionAsset"`
	TradeId         int64   `json:"tradeId"`
}

// response type: ACK RESULT FULL
type CreateOrderRespMsg struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	OrderListId   int64  `json:"orderListId"`
	ClientOrderId string `json:"clientOrderId"`
	TransactTime  int64  `json:"transactTime"`

	// ACK response does not contain the following fields
	Price               float64 `json:"price,string,omitempty"`
	OrigQty             float64 `json:"origQty,string,omitempty"`
	ExecutedQty         float64 `json:"executedQty,string,omitempty"`
	CummulativeQuoteQty float64 `json:"cummulativeQuoteQty,string,omitempty"`
	Status              string  `json:"status,omitempty"`
	TimeInForce         string  `json:"timeInForce,omitempty"`
	Type                string  `json:"type,omitempty"`
	Side                string  `json:"side,omitempty"`

	// the following fields are only for FULL response
	Fills []*OrderFill `json:"fills,omitempty"`
}
