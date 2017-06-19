package domains

import "time"

type StockDayDomainStruct struct {
	Id       int64
	Symbol   string    `xorm:"Varchar(10)"`
	Volume   int64     `xorm:"BigInt"`
	Open     float64   `xorm:"Double"`
	High     float64   `xorm:"Double"`
	Close    float64   `xorm:"Double"`
	Low      float64   `xorm:"Double"`
	Chg      float64   `xorm:"Double"`
	Percent  float64   `xorm:"Double"`
	Turnrate float64   `xorm:"Double"`
	Ma5      float64   `xorm:"Double"`
	Ma10     float64   `xorm:"Double"`
	Ma20     float64   `xorm:"Double"`
	Ma30     float64   `xorm:"Double"`
	Dif      float64   `xorm:"Double"`
	Dea      float64   `xorm:"Double"`
	Macd     float64   `xorm:"Double"`
	Time     string    `xorm:"-"`
	TimeUnix int64     `xorm:"notnull unique"`
	TimeFmt  time.Time `xorm:"timestamp"`
	Created  time.Time `xorm:"created timestamp 'gmt_created'"`
}
