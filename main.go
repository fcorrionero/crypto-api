package main

import (
	"encoding/json"
	"fmt"
	"github.com/fcorrionero/crypto-api/src/domain"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	url := "https://fapi.binance.com/fapi/v1/klines?symbol=VETUSDT&limit=1000&interval=5m"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))

	var parsed [][]interface{}
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		log.Println(err)
		return
	}
	var ks []domain.KlineData
	for _, r := range parsed {
		k := domain.KlineData{}
		oT := strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", r[0]), "0"), ".")
		oT = oT[0 : len(oT)-3]
		t, err := k.ParseTimeStamp(oT)
		if err != nil {
			log.Println(err)
		}
		cT := strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", r[6]), "0"), ".")
		t2, err := k.ParseTimeStamp(cT)
		if err != nil {
			log.Println(err)
		}
		k.OpenTime = t
		k.Open, _ = strconv.ParseFloat(r[1].(string), 64)
		k.High, _ = strconv.ParseFloat(r[2].(string), 64)
		k.Low, _ = strconv.ParseFloat(r[3].(string), 64)
		k.Close, _ = strconv.ParseFloat(r[4].(string), 64)
		k.Volume, _ = strconv.ParseFloat(r[5].(string), 64)
		k.CloseTime = t2
		k.QuoteAssetVolume, _ = strconv.ParseFloat(r[7].(string), 64)
		k.TradersNumber = r[8].(float64)
		k.TakerBuyBaseAssetVolume, _ = strconv.ParseFloat(r[9].(string), 64)
		k.TakerBuyQuoteAssetVolume, _ = strconv.ParseFloat(r[10].(string), 64)
		k.Ignore, _ = strconv.ParseFloat(r[11].(string), 64)
		ks = append(ks, k)
	}

	log.Println(ks[0].OpenTime)
	log.Println(ks[0].Open)

	log.Println(ks[len(ks)-1].OpenTime)
	log.Println(ks[len(ks)-1].Open)
}
