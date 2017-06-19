package models

import (
	"github.com/bamboox/aceStock/domains"
)

type StockRecords struct {
	Stocks  []domains.StockDomainStruct `json:"stocks"`
	Success string                      `json:"success"`
	count   CountStruct                 `json:"stock"`
}
type CountStruct struct {
	Count string `json:"count"`
}
type Stock struct {
	Symbol        string `json:"symbol"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	Current       string
	Percent       string
	Change        string
	High          string
	Low           string
	High52w       string
	Low52w        string
	Marketcapital string
	Amount        string
	Type          string
	Pettm         string
	Volume        string
	Hasexist      string
}
