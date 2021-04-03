package domain

import (
	"errors"
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

func (k KlineData) IsTIUPattern(ks []KlineData) (bool, error) {
	if len(ks) != 3 {
		return false, errors.New("only three points needed for TIU pattern")
	}
	// c1 := ks[0].Close < ks[0].Open // 1 + Downtrend
	c2 := ks[1].Close > ks[1].Open                               // 2
	c3 := ks[0].Open >= ks[1].Close && ks[1].Close > ks[0].Close // 2
	c4 := ks[0].Open >= ks[1].Open && ks[1].Open > ks[0].Close   // 2
	c5 := ks[2].Close > ks[2].Open && ks[2].Close > ks[0].Open
	return c2 && c3 && c4 && c5, nil
}

func (k KlineData) IsTIDPattern(ks []KlineData) (bool, error) {
	if len(ks) != 3 {
		return false, errors.New("only three points needed for TID pattern")
	}
	// 1) Uptrend + ks[0].Open > ks[0].Close
	c1 := ks[1].Close < ks[1].Open && (ks[0].Close >= ks[1].Close && ks[1].Close > ks[0].Open)
	c2 := ks[0].Close >= ks[1].Open && ks[1].Open > ks[0].Open
	c3 := ks[2].Close < ks[2].Open && ks[2].Close < ks[0].Open

	return c1 && c2 && c3, nil
}
