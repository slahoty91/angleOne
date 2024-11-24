package instrumentdata

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/slahoty91/tradingBot/copyOrderAngleOne/date"
	"github.com/slahoty91/tradingBot/copyOrderAngleOne/db"
	"github.com/slahoty91/tradingBot/copyOrderAngleOne/model"
)

func InsertData() {

	// fmt.Println("Equal date, session in progress")
	var newSession string
	fmt.Println("Enter instrument data:- ")
	fmt.Scanln(&newSession)
	var optionsTrue = [7]string{"Y", "y", "yes", "Yes", "YEs", "yES", "yEs"}
	// var optionsFalse = [7]string{"N", "n", "No", "no", "nO", "n0", "N0"}
	isTrue := false
	// isFalse := false
	for _, option := range optionsTrue {
		if newSession == option {
			isTrue = true
			break
		}

	}

	if isTrue {
		url := "https://margincalculator.angelbroking.com/OpenAPI_File/files/OpenAPIScripMaster.json"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err, "error from fetching instruments")

		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			return
		}
		// fmt.Println(body, "bodyyyyyy")
		var result []model.TokenData
		var insertData []model.TokenData
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return
		}

		for _, instrument := range result {
			if instrument.Name == "NIFTY" && instrument.InstrumentType == "OPTIDX" {
				instrument.InstType = instrument.Symbol[len(instrument.Symbol)-2:]
				floatValue, err := strconv.ParseFloat(instrument.Strike, 64)
				if err != nil {
					fmt.Println(err)
				}

				instrument.Strike_Int = int64(floatValue / 100)
				dateStr := instrument.Expiry
				// dateStr = strings.ToUpper(dateStr)
				// dateStr = strings.TrimSpace(dateStr)
				fmt.Println(dateStr, "dateStr", len(dateStr))
				instrument.Expiry_Zrd_Fmt = date.ConvertDate(dateStr)
				// parsedDate, err := time.Parse("02JAN2006", dateStr)
				// if err != nil {
				// 	fmt.Println(err, "errrr")
				// }
				// // fmt.Println(parsedDate, "parsedDate")
				// instrument.Expiry_Zrd_Fmt = parsedDate.Format("2006-01-02")
				fmt.Println(instrument.Expiry_Zrd_Fmt, "instrument.Expiry_Zrd_Fmt")
				insertData = append(insertData, instrument)
			}
		}
		var database = db.GetCollection("algoTrading", "instrumentAngleOne")

		var interfaceSlice []interface{}
		for _, data := range insertData {
			interfaceSlice = append(interfaceSlice, data)
		}
		database.InsertMany(context.Background(), interfaceSlice)
		fmt.Printf("Data inserted succesfully")
	}

}
