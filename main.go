package main

import (
	"fmt"

	"./bitflyer"
	"./config"
	"./utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	order := &bitflyer.Order{
		ProductCode:     config.Config.ProdectCode,
		ChildOrderType:  "MARKET",
		Side:            "BUY",
		Size:            0.01,
		MinuteToExpires: 1,
		TimeInForce:     "GTC",
	}
	res, _ := apiClient.SendOrder(order)
	fmt.Println(res.ChildOrderAcceptanceID)

	// tickerChannel := make(chan bitflyer.Ticker)
	// go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	// for ticker := range tickerChannel {
	// 	fmt.Println(ticker)
	// 	fmt.Println(ticker.GetMidPrice())
	// 	fmt.Println(ticker.DateTime())
	// 	fmt.Println(ticker.TruncateDateTime(time.Second))
	// 	fmt.Println(ticker.TruncateDateTime(time.Minute))
	// 	fmt.Println(ticker.TruncateDateTime(time.Hour))
	// }
}
