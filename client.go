package binance

import (
	"log"
	"net/http"
)

type Client struct {
	conn   *http.Client
	apiKey string
	signer *HmacSigner
	logger *log.Logger
	debug  bool

	Spot *SpotApi
}

func NewClient(conf *Config) *Client {
	c := &Client{
		conn:   &http.Client{},
		apiKey: conf.key,
		signer: &HmacSigner{Key: []byte(conf.secret)},
		logger: conf.logger,
		debug:  conf.debug,
	}
	return &Client{
		Spot: (*SpotApi)(c),
	}
}
