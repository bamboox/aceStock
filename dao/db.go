package dao

import (
	"time"

	//	log "github.com/Sirupsen/logrus"
	"github.com/bamboox/aceStock/common"
	"github.com/bamboox/aceStock/domains"
)

func SaveStocks(stocks []domains.StockDomainStruct) {
	for _, v := range stocks {
		var foundModel domains.StockDomainStruct
		has, _ := common.Engine.Where("symbol=?", v.Symbol).Get(&foundModel)

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
func FindStockList(foundModelsP *[]domains.StockDomainStruct) {

	//	foundModelsP := &foundModels
	var foundModel domains.StockDomainStruct
	rows, err := common.Engine.Desc("symbol").Rows(&foundModel)
	if err != nil {
		panic(err)
	}
	// SELECT * FROM user
	defer rows.Close()
	bean := new(domains.StockDomainStruct)
	for rows.Next() {
		err = rows.Scan(bean)
		*foundModelsP = append(*foundModelsP, *bean)
	}
}
func SaveStocksDayData(stocksDayDatas []domains.StockDayDomainStruct, symbol string) {
	for _, v := range stocksDayDatas {
		t, _ := time.Parse(time.RubyDate, v.Time)
		v.TimeUnix = t.Unix()
		v.TimeFmt = t
		v.Symbol = symbol
		var foundModel domains.StockDayDomainStruct
		has, _ := common.Engine.Where("time_unix=?", v.TimeUnix).And("symbol=?", symbol).Get(&foundModel)
		if !has {
			if _, err := common.Engine.Insert(&v); err != nil {
				panic(err)
			}
		}
	}
}
