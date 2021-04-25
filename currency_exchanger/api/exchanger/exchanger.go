package exchanger

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Envelope struct {
	Cube []struct {
		Date  string `xml:"time,attr"`
		Rates []struct {
			Currency string `xml:"currency,attr"`
			Rate     string `xml:"rate,attr"`
		} `xml:"Cube"`
	} `xml:"Cube>Cube"`
}

func GetRate(val float64, input string, output string, env Envelope) float64 {
	if output == "EUR" {
		var inputRate = GetRateUERtoInput(input, env)
		var f, _ = strconv.ParseFloat(inputRate, 8)
		var eur = val / f
		return eur
	}

	var inputRate = GetRateUERtoInput(input, env)
	var f, _ = strconv.ParseFloat(inputRate, 8)
	var eur = val / f
	var outRate = GetRateUERtoInput(output, env)
	f, _ = strconv.ParseFloat(outRate, 8)
	return f * eur
}

func GetRateUERtoInput(input string, env Envelope) string {
	for _, v := range env.Cube[0].Rates {
		if v.Currency == input {
			return v.Rate
		}
	}
	return ""
}

func GetEnvelope() Envelope {
	// get the latest exchange rate
	resp, err := http.Get("http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	xmlCurrenciesData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var env Envelope
	err = xml.Unmarshal(xmlCurrenciesData, &env)

	if err != nil {
		log.Fatal(err)
	}
	return env
}
