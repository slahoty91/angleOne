package copytrades

import (
	"context"
	"fmt"
	"log"

	"github.com/slahoty91/tradingBot/copyOrderAngleOne/db"
	placeorder "github.com/slahoty91/tradingBot/copyOrderAngleOne/placeOrder"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Remaning works:-
// Get live balance and quantity size cal
// Sell order
// Cloud function for account integration others
// Place order in other accounts

func WatchForChenges() {

	coll := db.GetCollection("algoTrading", "masterTable")
	changeStream, err := coll.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		log.Fatalln(err)
	}
	defer changeStream.Close(context.TODO())
	for changeStream.Next(context.TODO()) {
		var changeEvent bson.M
		err := changeStream.Decode(&changeEvent)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Change detected", changeEvent)
		operationType := changeEvent["operationType"].(string)
		documentKey := changeEvent["documentKey"].(bson.M)["_id"]
		updatedFields := changeEvent["updateDescription"].(bson.M)["updatedFields"]

		fmt.Printf("Operation: %s, Document ID: %v, Updated Fields: %v\n", operationType, documentKey, updatedFields)
		// updatedFields["orderActive"]
		if fieldsMap, ok := updatedFields.(bson.M); ok {
			// Safely access "orderActive"
			if activeOrderAsset, exists := fieldsMap["activeOrderAsset"]; exists {
				fmt.Printf("orderActive: %v\n", activeOrderAsset)
				if assetStr, ok := activeOrderAsset.(string); ok {
					// fmt.Printf("orderActive: %s\n", assetStr,len(assetStr))
					fmt.Println("Asset==>", assetStr, len(assetStr))
					if len(assetStr) > 1 {
						angleOneSymbol := placeorder.ConvertSymbol(assetStr)
						fmt.Println(angleOneSymbol, "angleOneSymbol")
						placeorder.PlaceOrder("BUY", angleOneSymbol.Symbol, angleOneSymbol.Token, angleOneSymbol.LotSize, fieldsMap["orderId"].(string))
					}

				}

			}
			if exitorder, exists := fieldsMap["exitorder"]; exists {
				fmt.Printf("orderActive: %v\n", exitorder)
				if exitOrderSymbol, ok := exitorder.(string); ok {
					// fmt.Printf("orderActive: %s\n", exitOrderSymbol,len(exitOrderSymbol))
					fmt.Println("Asset==>", exitOrderSymbol, len(exitOrderSymbol))
					if len(exitOrderSymbol) > 1 {
						angleOneSymbol := placeorder.ConvertSymbol(exitOrderSymbol)
						fmt.Println(angleOneSymbol, "angleOneSymbol")
						placeorder.PlaceOrder("SELL", angleOneSymbol.Symbol, angleOneSymbol.Token, angleOneSymbol.LotSize, "")
					}

				}

			} else {
				fmt.Println("Key 'orderActive' does not exist in updatedFields")
			}
		} else {
			fmt.Println("updatedFields is not a bson.M")
		}
	}
}
