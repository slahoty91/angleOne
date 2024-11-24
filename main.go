package main

import (
	"fmt"

	copytrades "github.com/slahoty91/tradingBot/copyOrderAngleOne/copyTrades"
	mongo "github.com/slahoty91/tradingBot/copyOrderAngleOne/db"
	gentoken "github.com/slahoty91/tradingBot/copyOrderAngleOne/genToken"
	instrumentdata "github.com/slahoty91/tradingBot/copyOrderAngleOne/instrumentData"
	"github.com/slahoty91/tradingBot/copyOrderAngleOne/test"
	// placeorder "github.com/slahoty91/tradingBot/copyOrderAngleOne/placeOrder"
)

func main() {
	fmt.Println("HIII NEW GO PROJECT CREATED")
	test.Test()
	mongo.ConnectDB()
	gentoken.GenerateToken()
	// placeorder.PlaceOrder("BUY", "", "", "")
	instrumentdata.InsertData()
	copytrades.WatchForChenges()

}
