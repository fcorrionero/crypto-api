package binance

import (
	"io/ioutil"
	"net/http"
)

type Api struct {
	klinesUrl  string
	symbolsUrl string
}

func New() Api {
	return Api{
		klinesUrl:  "https://fapi.binance.com/fapi/v1/klines",
		symbolsUrl: "https://fapi.binance.com/fapi/v1/ticker/24hr",
	}
}

func (a Api) Klines(symbol string, limit string, interval string) ([]byte, error) {
	if len(symbol) == 0 {
		symbol = "BTCUSDT"
	}
	if len(limit) == 0 {
		limit = "1000"
	}
	if len(interval) == 0 {
		interval = "1h"
	}
	query := "?symbol=" + symbol + "&limit=" + limit + "&interval=" + interval
	a.klinesUrl = a.klinesUrl + query

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, a.klinesUrl, nil)

	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func (a Api) Symbols() ([]byte, error) {

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, a.symbolsUrl, nil)

	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
