package main

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/bamboox/aceStock/common"
	"github.com/bamboox/aceStock/dao"
	//	"github.com/bamboox/aceStock/domains"
	aceHttp "github.com/bamboox/aceStock/http"
	"github.com/bamboox/aceStock/models"
	"github.com/bamboox/aceStock/utils"
)

func main() {
	common.InitEngine(nil)
	client := &aceHttp.HttpClient{}
	client.Client = &http.Client{}
	loginRst := client.Login("https://xueqiu.com/user/login", map[string]string{
		"areacode":    "86",
		"remember_me": "on",
		"password":    utils.Md5("xx"),
		"telephone":   "xxx",
	})
	log.Printf(loginRst)
	client.Get("https://xueqiu.com/1637386964")

	// get all stocker data
	//	i := 1
	//	for i <= 59 {
	//		log.Printf(utils.Int2Str(i))
	//		fetchStockTargetUrl := utils.StrJion("https://xueqiu.com/stock/cata/stocklist.json?page=", utils.Int2Str(i), "&size=90&order=desc&orderby=percent&type=11%2C12&_=", utils.GetTimeStr())

	//		fetchStockRst := client.FetchStock(fetchStockTargetUrl)

	//		log.Printf(fetchStockRst)
	//		stockRecords := models.StockRecords{}
	//		err := json.Unmarshal([]byte(fetchStockRst), &stockRecords)
	//		if err != nil {
	//			log.Printf("err was %v", err)
	//		}
	//		log.Println(stockRecords)
	//		dao.SaveStocks(stockRecords.Stocks)
	//		time.Sleep(2000 * time.Millisecond)
	//		i = i + 1
	//	}

	// get StockDayRecords
	//	analysisRst := client.Analysis("https://xueqiu.com/stock/forchartk/stocklist.json?symbol=SH000001&period=1day&type=normal&begin=1466264872445&end=1497800872445&_=1497800872445")
	//	log.Printf(analysisRst)
	//	stockRecords := models.StockDayRecords{}
	//	err := json.Unmarshal([]byte(analysisRst), &stockRecords)
	//	if err != nil {
	//		log.Printf("err was %v", err)
	//	}
	//	log.Println(stockRecords)
	//	dao.SaveStocksDayData(stockRecords.Chartlist)

}
