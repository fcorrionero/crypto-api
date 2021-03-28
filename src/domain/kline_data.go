package domain

import (
	"strconv"
	"time"
)

type KlineData struct {
	OpenTime                 time.Time
	Open                     float64
	High                     float64
	Low                      float64
	Close                    float64
	Volume                   float64
	CloseTime                time.Time
	QuoteAssetVolume         float64
	TradersNumber            float64
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
	Ignore                   float64
}

func (k KlineData) ParseTimeStamp(utime string) (time.Time, error) {
	i, err := strconv.ParseInt(utime, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	t := time.Unix(i, 0)
	//t.Format(time.RFC3339)
	return t, nil
}
