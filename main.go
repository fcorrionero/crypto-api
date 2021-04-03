package main

import (
	"github.com/fcorrionero/crypto-api/src/application/query/find_kline_data"
	"github.com/fcorrionero/crypto-api/src/infrastructure/binance"
)

func main() {
	api := binance.New()

	//qSymb := find_symbols.New(api)
	kData := find_kline_data.New(api)

	query := find_kline_data.Query{
		Symbol:   "ONEUSDT",
		Limit:    "50",
		Interval: "1h",
	}
	kData.Handle(query)

	//symbols := qSymb.Handle()
	//for _, s := range symbols {
	//	fmt.Println(s)
	//	query := find_kline_data.Query{
	//		Symbol:   s,
	//		Limit:    "3",
	//		Interval: "1h",
	//	}
	//	kData.Handle(query)
	//}

	//symbols := find_symbols.New(api)
	//symbols.Handle()
}
