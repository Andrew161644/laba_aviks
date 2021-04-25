package database_test

import (
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
	"testing"
)

func TestCanGetCurrencyByName(t *testing.T) {
	var result, err = db.GetCurrencyByName(models.Currency{Name: "EUR"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}

func TestCanGetAllCurrencies(t *testing.T) {
	var result, err = db.GetAllCurrencies()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}
