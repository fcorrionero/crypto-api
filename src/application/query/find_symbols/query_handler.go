package find_symbols

import (
	"encoding/json"
	"github.com/fcorrionero/crypto-api/src/infrastructure/binance"
	"log"
)

type QueryHandler struct {
	api binance.Api
}

func New(api binance.Api) QueryHandler {
	return QueryHandler{
		api: api,
	}
}

func (q QueryHandler) Handle() []string {
	var symbs []string
	body, err := q.api.Symbols()
	if nil != err {
		log.Println(err)
		return symbs
	}
	//fmt.Println(string(body))
	var parsed []map[string]interface{}
	err = json.Unmarshal(body, &parsed)
	for _, m := range parsed {
		symbs = append(symbs, m["symbol"].(string))
	}

	return symbs
}
