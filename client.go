package binance

import (
	"net/http"
)

type Client struct {
	conn   *http.Client
	apiKey string
	signer *HmacSigner
	Spot   *SpotApi
}

func NewClient(conf *Config) *Client {
	c := &Client{
		conn:   &http.Client{},
		apiKey: conf.Key,
		signer: &HmacSigner{Key: []byte(conf.Secret)},
	}
	return &Client{
		Spot: (*SpotApi)(c),
	}
}
