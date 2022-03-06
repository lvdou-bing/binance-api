package binance

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"
)

type SpotApi Client

func (spot *SpotApi) Ping() (*http.Response, error) {
	return spot.conn.Get(baseUrl + "/api/v3/ping")
}

func (spot *SpotApi) CreateLimitOrder(symbol string, side string, quantity string, price string, timeInForce string) (*http.Response, error) {
	values := url.Values{}
	values.Add("type", "LIMIT")
	values.Add("symbol", symbol)
	values.Add("side", side)
	values.Add("quantity", quantity)
	values.Add("price", price)
	values.Add("timeInForce", timeInForce)
	timestampStr := strconv.FormatInt(time.Now().UnixMilli(), 10)
	values.Add("timestamp", timestampStr)
	signature := spot.signer.Sign([]byte(values.Encode()))
	values.Add("signature", signature)

	body := bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", baseUrl+"/api/v3/order", body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-MBX-APIKEY", spot.apiKey)

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	log.Printf("\n%s\n", string(dump))
	return spot.conn.Do(req)
}

func (spot *SpotApi) CreateMarketBaseQuantityOrder(symbol string, side string, quantity string) (*http.Response, error) {
	values := url.Values{}
	values.Add("type", "MARKET")
	values.Add("symbol", symbol)
	values.Add("side", side)
	values.Add("quantity", quantity)
	timestampStr := strconv.FormatInt(time.Now().UnixMilli(), 10)
	values.Add("timestamp", timestampStr)
	signature := spot.signer.Sign([]byte(values.Encode()))
	values.Add("signature", signature)

	body := bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", baseUrl+"/api/v3/order", body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-MBX-APIKEY", spot.apiKey)

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	log.Printf("\n%s\n", string(dump))
	return spot.conn.Do(req)
}

func (spot *SpotApi) CreateMarketQuoteQuantityOrder(symbol string, side string, quoteOrderQty string) (*http.Response, error) {
	values := url.Values{}
	values.Add("type", "MARKET")
	values.Add("symbol", symbol)
	values.Add("side", side)
	values.Add("quoteOrderQty", quoteOrderQty)
	timestampStr := strconv.FormatInt(time.Now().UnixMilli(), 10)
	values.Add("timestamp", timestampStr)
	signature := spot.signer.Sign([]byte(values.Encode()))
	values.Add("signature", signature)

	body := bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", baseUrl+"/api/v3/order", body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-MBX-APIKEY", spot.apiKey)

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	log.Printf("\n%s\n", string(dump))
	return spot.conn.Do(req)
}

func (spot *SpotApi) CancelOrderById(symbol string, orderId string) (*http.Response, error) {
	values := url.Values{}
	values.Add("symbol", symbol)
	values.Add("orderId", orderId)
	timestampStr := strconv.FormatInt(time.Now().UnixMilli(), 10)
	values.Add("timestamp", timestampStr)
	signature := spot.signer.Sign([]byte(values.Encode()))
	values.Add("signature", signature)

	body := bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("DELETE", baseUrl+"/api/v3/order", body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-MBX-APIKEY", spot.apiKey)

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	log.Printf("\n%s\n", string(dump))
	return spot.conn.Do(req)
}

func (spot *SpotApi) CancelOrdersBySymbol(symbol string) (*http.Response, error) {
	values := url.Values{}
	values.Add("symbol", symbol)
	timestampStr := strconv.FormatInt(time.Now().UnixMilli(), 10)
	values.Add("timestamp", timestampStr)
	signature := spot.signer.Sign([]byte(values.Encode()))
	values.Add("signature", signature)

	body := bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("DELETE", baseUrl+"/api/v3/openOrders", body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-MBX-APIKEY", spot.apiKey)

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	log.Printf("\n%s\n", string(dump))
	return spot.conn.Do(req)
}

func (spot *SpotApi) QueryOrderById(symbol string, orderId string) (*http.Response, error) {
	values := url.Values{}
	values.Add("symbol", symbol)
	values.Add("orderId", orderId)
	timestampStr := strconv.FormatInt(time.Now().UnixMilli(), 10)
	values.Add("timestamp", timestampStr)
	signature := spot.signer.Sign([]byte(values.Encode()))
	values.Add("signature", signature)

	req, err := http.NewRequest("GET", baseUrl+"/api/v3/order", nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	req.Header.Add("X-MBX-APIKEY", spot.apiKey)

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	log.Printf("\n%s\n", string(dump))
	return spot.conn.Do(req)
}

func (spot *SpotApi) QueryOpenOrdersBySymbol(symbol string) (*http.Response, error) {
	values := url.Values{}
	values.Add("symbol", symbol)
	timestampStr := strconv.FormatInt(time.Now().UnixMilli(), 10)
	values.Add("timestamp", timestampStr)
	signature := spot.signer.Sign([]byte(values.Encode()))
	values.Add("signature", signature)

	req, err := http.NewRequest("GET", baseUrl+"/api/v3/openOrders", nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	req.Header.Add("X-MBX-APIKEY", spot.apiKey)

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	log.Printf("\n%s\n", string(dump))
	return spot.conn.Do(req)
}

func (spot *SpotApi) QueryAllOpenOrders() (*http.Response, error) {
	values := url.Values{}
	timestampStr := strconv.FormatInt(time.Now().UnixMilli(), 10)
	values.Add("timestamp", timestampStr)
	signature := spot.signer.Sign([]byte(values.Encode()))
	values.Add("signature", signature)

	req, err := http.NewRequest("GET", baseUrl+"/api/v3/openOrders", nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()
	req.Header.Add("X-MBX-APIKEY", spot.apiKey)

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}
	log.Printf("\n%s\n", string(dump))
	return spot.conn.Do(req)
}
