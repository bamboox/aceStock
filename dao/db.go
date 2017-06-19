package dao

import (
	"time"

	"github.com/bamboox/aceStock/common"
	"github.com/bamboox/aceStock/domains"
)

func SaveStocks(stocks []domains.StockDomainStruct) {
	for _, v := range stocks {
		var foundModel domains.StockDomainStruct
		has, _ := common.Engine.Where("code=?", v.Code).Get(&foundModel)

		if has {
			if _, err := common.Engine.Id(&foundModel.Id).Update(&v); err != nil {
				panic(err)
			}
		} else {
			if _, err := common.Engine.Insert(&v); err != nil {
				panic(err)
			}
		}
	}
}
func SaveStocksDayData(stocksDayDatas []domains.StockDayDomainStruct) {
	for _, v := range stocksDayDatas {
		t, _ := time.Parse(time.RubyDate, v.Time)
		v.TimeUnix = t.Unix()
		v.TimeFmt = t
		var foundModel domains.StockDayDomainStruct
		has, _ := common.Engine.Where("time_unix=?", v.TimeUnix).Get(&foundModel)
		if !has {
			if _, err := common.Engine.Insert(&v); err != nil {
				panic(err)
			}
		}
	}
}
