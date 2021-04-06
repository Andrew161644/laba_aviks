package database_test

import (
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
	"testing"
)

func TestCanAddYear(t *testing.T) {
	var id, err = db.AddYear(models.Year{
		Name:        "2022",
		Description: "Прекрасный год для бизнеса",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id)
}

func TestCanYearByName(t *testing.T) {
	var year, err = db.GetYearByName("2014")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(year)
}

func TestCanGetAllYears(t *testing.T) {
	var years, err = db.GetAllYears()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(years)
}

func TestDeleteYear(t *testing.T) {
	var err = db.DeleteYear(models.Year{ID: 13})
	if err != nil {
		log.Fatal(err)
	}
}
