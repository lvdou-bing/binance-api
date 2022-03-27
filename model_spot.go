package binance

type ExchangeInfoFilter struct {
	FilterType string `json:"filterType"`

	// for PRICE_FILTER
	MinPrice string `json:"minPrice,omitempty"`
	MaxPrice string `json:"maxPrice,omitempty"`
	TickSize string `json:"tickSize,omitempty"`

	// for PERCENT_PRICE
	MultiplierUp   string `json:"multiplierUp,omitempty"`
	MultiplierDown string `json:"multiplierDown,omitempty"`
	AvgPriceMins   int    `json:"avgPriceMins,omitempty"` // also for MIN_NOTIONAL

	// for LOT_SIZE and MARKET_LOT_SIZE
	MinQty   string `json:"minQty,omitempty"`
	MaxQty   string `json:"maxQty,omitempty"`
	StepSize string `json:"stepSize,omitempty"`

	// for MIN_NOTIONAL
	MinNotional   string `json:"minNotional,omitempty"`
	ApplyToMarket bool   `json:"applyToMarket,omitempty"`

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
	ExchangeInfoList []*ExchangeInfo `json:"symbols"`
}
