package database_test

import (
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
	"testing"
)

func TestCanAddIndex(t *testing.T) {
	var id, error = db.AddIndex(models.Index{
		PastYear:    2,
		CurrentYear: 3,
		Value:       1.057,
		CurrencyId:  2,
	})
	if error != nil {
		log.Fatal(error)
	}
	log.Println(id)
}

func TestGetAllIndicesByCurrentYear(t *testing.T) {
	var indexes, err = db.GetAllIndicesByCurrentYear(models.Index{
		CurrentYear: 12,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(indexes)
}

func TestCanIndicesByCurrentYearAndCurrency(t *testing.T) {
	var indexes, err = db.GetAllIndicesByCurrentYearAndCurrency(models.Index{
		CurrentYear: 2,
		CurrencyId:  1,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(indexes)
}

func TestCanGetAllIndices(t *testing.T) {
	var indexes, err = db.GetAllIndices()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(indexes)
}
