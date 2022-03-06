package binance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	baseUrl = "https://api.binance.com"
)

type Config struct {
	Key    string
	Secret string
}

func NewConfig(configFile string) (*Config, error) {
	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var config Config
	err = json.Unmarshal(f, &config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &config, nil
}
