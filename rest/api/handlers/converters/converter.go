package converters

import (
	"log"
	"strconv"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func StringToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 8)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
