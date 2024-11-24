package date

import (
	"fmt"
	"time"
)

func CurrentDate() string {
	// Get the current date and time
	currentTime := time.Now()

	// Format the date as "YYYY-MM-DD"
	formattedDate := currentTime.Format("2006-01-02")

	// Print the formatted date
	fmt.Println("Current Date:", formattedDate)
	return formattedDate
}

func ConvertDate(inp string) string {

	months := map[string]string{
		"JAN": "01", "FEB": "02", "MAR": "03", "APR": "04",
		"MAY": "05", "JUN": "06", "JUL": "07", "AUG": "08",
		"SEP": "09", "OCT": "10", "NOV": "11", "DEC": "12",
	}

	// Extract parts of the date
	day := inp[0:2]
	month := inp[2:5]
	year := inp[5:]

	// Convert to desired format
	return fmt.Sprintf("%s-%s-%s", year, months[month], day)

}
