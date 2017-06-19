package domains

import "time"

type StockDomainStruct struct {
	Id            int64
	Symbol        string
	Code          string `xorm:"notnull unique"`
	Name          string
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
	Created       time.Time `xorm:"created timestamp 'gmt_created'"`
	Updated       time.Time `xorm:"updated timestamp 'gmt_updated'"`
}
