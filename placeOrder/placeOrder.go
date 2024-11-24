package placeorder

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/angel-one/smartapigo"
	"github.com/slahoty91/tradingBot/copyOrderAngleOne/db"
	gentoken "github.com/slahoty91/tradingBot/copyOrderAngleOne/genToken"
	"github.com/slahoty91/tradingBot/copyOrderAngleOne/model"
	"go.mongodb.org/mongo-driver/bson"
)

func PlaceOrder(transactionType string, tradingSymbol string, token string, quantity string, orderID string) {
	doc, client := gentoken.GetClient()
	fmt.Println(client, "clienttttt")
	accTkn := doc.AngleOne.AccToken
	fmt.Println(accTkn, "accTkn")
	client.SetAccessToken(accTkn)

	var ltpResp smartapigo.LTPResponse
	var err error
	lotSize := 0
	quat := ""
	angleOneCol := db.GetCollection("algoTrading", "angleOneOrders")
	if transactionType == "BUY" {

		ltpParms := smartapigo.LTPParams{
			Exchange:      "NFO",
			TradingSymbol: tradingSymbol,
			SymbolToken:   token,
		}
		ltpResp, err = client.GetLTP(ltpParms)
		if err != nil {
			fmt.Println(err, "err from LTP")
		}
		fmt.Println(ltpResp, lotSize)
		lotSize, quat = getLotSize(client, ltpResp.Ltp)
		fmt.Println(quat, "quat", lotSize, "lotSize")

	}
	var orderActive model.AngleOneOrder
	if transactionType == "SELL" {

		res := angleOneCol.FindOne(context.Background(), bson.M{"Status": "Active", "Symbol": tradingSymbol})
		err := res.Decode(orderActive)
		if err != nil {
			fmt.Println(err, "error from order decode")
		}
		lotSize = orderActive.LotSize
		quat = orderActive.Qty
		orderID = orderActive.OrderID
	}

	payLoad := smartapigo.OrderParams{
		Variety:         "NORMAL",
		TradingSymbol:   tradingSymbol,
		SymbolToken:     token,
		TransactionType: transactionType,
		Exchange:        "NFO",
		OrderType:       "MARKET",
		ProductType:     "INTRADAY",
		Duration:        "DAY",
		Price:           "0",
		SquareOff:       "0",
		StopLoss:        "0",
		Quantity:        quantity,
	}
	fmt.Println(payLoad, "PAYLOADDDDDDDDD")
	ordResp, err := client.PlaceOrder(payLoad)

	if err != nil {
		fmt.Println(err, "errrrr")
	}
	fmt.Println(ordResp, "ordResp")
	angleOneOrder := model.AngleOneOrder{
		Script:      ordResp.Script,
		OrderID:     ordResp.OrderID,
		Symbol:      tradingSymbol,
		SymbolToken: token,
		Qty:         quat,
		LotSize:     lotSize,
		TransType:   transactionType,
		ZeroOrderID: orderID,
	}
	fmt.Println(angleOneOrder)
	result, err := angleOneCol.InsertOne(context.Background(), angleOneOrder)
	if err != nil {
		fmt.Println(err, "err from save document order")
	}
	fmt.Println("insert result===>", result)
}

// type struct
func ConvertSymbol(symbol string) model.TokenData {
	if symbol == "" {
		return model.TokenData{}
	}
	var zerodha_mod model.OptionInstrument
	fmt.Println(symbol, "symbol")
	database_angleOne := db.GetCollection("algoTrading", "instrumentAngleOne")
	collection_zer_inst := db.GetCollection("algoTrading", "instrumentNFO")
	res := collection_zer_inst.FindOne(context.Background(), bson.M{"tradingsymbol": symbol})
	res.Decode(&zerodha_mod)
	fmt.Println(database_angleOne, "resssssss")
	strikePricePart := zerodha_mod.Strike    // 24400
	optionType := zerodha_mod.InstrumentType // CE

	// Print the results
	fmt.Println("Strike Price:", strikePricePart)
	fmt.Println("Option Type:", optionType)
	filter := bson.M{
		"strike_int":     strikePricePart,
		"inst_type":      optionType,
		"expiry_zer_fmt": zerodha_mod.Expiry,
	}
	// projection := bson.M{
	// 	"symbol": 1, // Include trading symbol
	// 	"token":  1,
	// 	"_id":    0, // Exclude _id by default
	// }
	fmt.Println(filter, "filterrrrr")
	res = database_angleOne.FindOne(context.Background(), filter)
	var result model.TokenData
	err := res.Decode(&result)
	if err != nil {
		fmt.Println("error in decoding", err)
	}
	fmt.Println("trading symbol", result)
	// tradingSymbol := result.Symbol

	return result
}

func getCapital(client *smartapigo.Client) smartapigo.RMS {

	rms, err := client.GetRMS()
	if err != nil {
		fmt.Println(err, "err from RMS")
	}

	return rms

}

func getLotSize(client *smartapigo.Client, ltp float64) (int, string) {

	rms := getCapital(client)

	cash := rms.AvailableCash

	cashFlot, err := strconv.ParseFloat(cash, 64)
	if err != nil {
		fmt.Println(err, "err from parse float")
	}
	quantity := int(math.Floor(cashFlot / ltp))
	numOfLots := quantity / 25
	return numOfLots, strconv.Itoa(quantity)
}
