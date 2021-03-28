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
	TradersNumber            int
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
	Ignore                   float64
}

func (k *KlineData) ParseTimeStamp(utime string) (string, error) {
	i, err := strconv.ParseInt(utime, 10, 64)
	if err != nil {
		return "", err
	}
	t := time.Unix(i, 0)
	return t.Format(time.UnixDate), nil
}
