package models

import (
	"github.com/bamboox/aceStock/domains"
)

type StockDayRecords struct {
	Chartlist []domains.StockDayDomainStruct `json:"chartlist"`
	Success   string                         `json:"success"`
	Stock     StockStruct                    `json:"stock"`
}
type StockStruct struct {
	Symbol string `json:"symbol"`
}
type chart struct {
	Volume   int64   `json:"volume"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Close    float64 `json:"close"`
	Low      float64 `json:"low"`
	Chg      float64 `json:"chg"`
	Percent  float64 `json:"percent"`
	Turnrate float64 `json:"turnrate"`
	Ma5      float64 `json:"ma5"`
	Ma10     float64 `json:"ma10"`
	Ma20     float64 `json:"ma20"`
	Ma30     float64 `json:"ma30"`
	Dif      float64 `json:"dif"`
	Dea      float64 `json:"dea"`
	Macd     float64 `json:"macd"`
	Time     string  `json:"time"`
}
