package domain

import "time"

type KlineData struct {
	OpenTime                 time.Time
	Open                     float64
	High                     float64
	Low                      float64
	Close                    float64
	Volume                   float64
	CloseTime                time.Time
	QuoteAssetVolume         float64
	TradersNumber            int
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
	Ignore                   float64
}
