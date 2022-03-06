package main

import (
	"binance"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var config binance.KeySecret
	err = json.Unmarshal(f, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config.Key)
	fmt.Println(config.Secret)
	c := binance.NewClient(config)
	resp, err := c.Spot.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	resp, err = c.Spot.CreateLimitOrder("BTCUSDT", "BUY", "0.001", "37000", "GTC")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err = io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
