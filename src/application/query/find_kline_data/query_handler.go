package find_kline_data

import (
	"encoding/json"
	"fmt"
	"github.com/fcorrionero/crypto-api/src/domain"
	"github.com/fcorrionero/crypto-api/src/infrastructure/binance"
	"gonum.org/v1/gonum/stat"
	"log"
	"strconv"
	"strings"
	"time"
)

type QueryHandler struct {
	api binance.Api
}

func New(api binance.Api) QueryHandler {
	return QueryHandler{
		api: api,
	}
}

func (q QueryHandler) Handle(query Query) {
	body, err := q.api.Klines(query.Symbol, query.Limit, query.Interval)
	if nil != err {
		log.Println(err)
		return
	}
	var parsed [][]interface{}
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		log.Println(err)
		fmt.Println(string(body))
		return
	}
	ks := q.parseKlines(parsed)

	var x, y []float64
	var patterns []domain.KlineData
	i := 0
	for _, k := range ks {
		x = append(x, float64(k.OpenTime.Unix()))
		y = append(y, k.Close)
		patterns = append(patterns, k)
		i++
		if i%3 == 0 { // 3 days/intervals moving average https://www.hindawi.com/journals/mpe/2017/3096917/
			b, a := stat.LinearRegression(x, y, nil, false)
			m, d := stat.MeanStdDev(y, nil)
			TIU, err := k.IsTIUPattern(patterns)
			if err != nil {
				log.Println(err)
			}
			TID, err := k.IsTIDPattern(patterns)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s : %.4v  |   %.4v  | %.4v  | %.4v | %t | %t\n", k.OpenTime.Format(time.Stamp), a, b, m, d, TIU, TID)
			x = nil
			y = nil
			patterns = nil
		}
	}
}

func (q QueryHandler) parseKlines(parsed [][]interface{}) []domain.KlineData {
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
	return ks
}
