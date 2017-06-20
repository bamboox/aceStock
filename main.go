package main

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/bamboox/aceStock/common"
	"github.com/bamboox/aceStock/dao"
	"github.com/bamboox/aceStock/domains"
	aceHttp "github.com/bamboox/aceStock/http"
	"github.com/bamboox/aceStock/models"
	"github.com/bamboox/aceStock/utils"
)

// get all stocker data
func getAllStockerData(client *aceHttp.HttpClient) {
	i := 1
	for i <= 59 {
		log.Printf(utils.Int2Str(i))
		fetchStockTargetUrl := utils.StrJion("https://xueqiu.com/stock/cata/stocklist.json?page=", utils.Int2Str(i), "&size=90&order=desc&orderby=percent&type=11%2C12&_=", utils.GetTimeStr())

		fetchStockRst := client.FetchStock(fetchStockTargetUrl)

		stockRecords := models.StockRecords{}
		err := json.Unmarshal([]byte(fetchStockRst), &stockRecords)
		if err != nil {
			log.Printf("err was %v", err)
		}
		log.Println(len(stockRecords.Stocks))
		dao.SaveStocks(stockRecords.Stocks)
		time.Sleep(2000 * time.Millisecond)
		i = i + 1
	}
}
func StockDayRecords(client *aceHttp.HttpClient, symbol string) {
	beginStr := utils.GetTimeStrByIn("2010-01-01 00:00:00")
	endStr := utils.GetTimeStrByIn("2017-06-21 00:00:00")
	//	before normal after
	analysisTargetUrl := "https://xueqiu.com/stock/forchartk/stocklist.json?symbol=" + symbol + "&period=1day&type=before&begin=" + beginStr + "&end=" + endStr + "&_=" + utils.GetTimeStr()
	log.Println(symbol)
	analysisRst, err := client.Analysis(analysisTargetUrl)
	if err != nil { //retry
		//
		log.Printf("err retry %v ", symbol)
		time.Sleep(10000 * time.Millisecond)
		login(client)
		StockDayRecords(client, symbol)
	}
	stockRecords := models.StockDayRecords{}
	err = json.Unmarshal([]byte(analysisRst), &stockRecords)
	if err != nil {
		log.Printf("err was %v", err)
		log.Printf("err was %v", analysisRst)
		time.Sleep(10000 * time.Millisecond)

		login(client)
		StockDayRecords(client, symbol)
	}
	log.Println(len(stockRecords.Chartlist))
	dao.SaveStocksDayData(stockRecords.Chartlist, symbol)
}
func login(client *aceHttp.HttpClient) {
	loginRst := client.Login("https://xueqiu.com/user/login", map[string]string{
		"areacode":    "86",
		"remember_me": "on",
		"password":    utils.Md5("1qaz2wsx"),
		"telephone":   "18818280053",
	})
	log.Printf(loginRst)
	client.Get("https://xueqiu.com/1637386964")
}
func main() {
	common.InitEngine(nil)
	client := &aceHttp.HttpClient{}
	client.Client = &http.Client{}
	//login
	login(client)

	// get all stocker data
	//	getAllStockerData(client)
	//getAllStockerDataFormDB
	foundModels := make([]domains.StockDomainStruct, 0)
	dao.FindStockList(&foundModels)
	//	for _, v := range foundModels {
	//		StockDayRecords(client, v.Symbol)
	//	}
	//error
	var foundModelsCopy []domains.StockDomainStruct
	for i, v := range foundModels {
		if v.Symbol == "SZ300047" {
			foundModelsCopy = foundModels[i:]
			break
		}

	}
	for _, v := range foundModelsCopy {
		StockDayRecords(client, v.Symbol)
	}

	// get StockDayRecords
	//	StockDayRecords(client, "SH603909")
}
