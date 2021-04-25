package providers

import (
	"database/sql"
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
)

func (db Database) GetCurrencyByName(currency models.Currency) (models.Currency, error) {
	currencyRes := models.Currency{}
	query := `SELECT * FROM currency WHERE name = $1;`
	row := db.Conn.QueryRow(query, currency.Name)

	switch err := row.Scan(&currencyRes.ID, &currencyRes.Name, &currencyRes.Description); err {
	case sql.ErrNoRows:
		return currencyRes, ErrNoMatch
	default:
		return currencyRes, err
	}
}

func (db Database) GetAllCurrencies() ([]models.Currency, error) {
	query := `SELECT * FROM currency`
	rows, err := db.Conn.Query(query)
	var curs []models.Currency
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var cur models.Currency
		err := rows.Scan(&cur.ID, &cur.Name, &cur.Description)
		if err != nil {
			return curs, err
		}
		curs = append(curs, cur)
	}
	return curs, err
}
