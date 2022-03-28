package binance

import (
	"log"
)

const (
	baseUrl = "https://api.binance.com"
)

type Config struct {
	key    string
	secret string
	logger *log.Logger
	debug  bool
}

func NewConfig(key, secret string, logger *log.Logger, debug bool) (*Config, error) {
	if logger == nil {
		logger = log.Default()
	}
	config := &Config{
		key:    key,
		secret: secret,
		logger: logger,
		debug:  debug,
	}
	return config, nil
}
