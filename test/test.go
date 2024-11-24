package test

import (
	"fmt"
	"math"
	"strconv"
	// "time"
)

func Test() {

	input := "19DEC2024"
	months := map[string]string{
		"JAN": "01", "FEB": "02", "MAR": "03", "APR": "04",
		"MAY": "05", "JUN": "06", "JUL": "07", "AUG": "08",
		"SEP": "09", "OCT": "10", "NOV": "11", "DEC": "12",
	}

	// Extract parts of the date
	day := input[0:2]
	month := input[2:5]
	year := input[5:]

	// Convert to desired format
	formattedDate := fmt.Sprintf("%s-%s-%s", year, months[month], day)
	fmt.Printf("Converted date: %s\n", formattedDate)

	// rms := ""

	cash := "3340.68"
	ltp := 123.45
	cashFlot, err := strconv.ParseFloat(cash, 64)
	if err != nil {
		fmt.Println(err, "err from parse float")
	}
	quantity := int(math.Floor(cashFlot / ltp))
	numOfLots := quantity / 25

	fmt.Println("numOfLots==>", numOfLots, "quantity==>", quantity)

}
