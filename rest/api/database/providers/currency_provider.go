package providers

import (
	"database/sql"
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
)

func (db Database) AddCurrency(currency models.Currency) (int, error) {
	var id int
	query := `INSERT INTO currency (name) values ($1) RETURNING id`
	err := db.Conn.QueryRow(query, currency.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db Database) GetCurrencyByName(currency models.Currency) (models.Currency, error) {
	currencyRes := models.Currency{}
	query := `SELECT * FROM currency WHERE name = $1;`
	row := db.Conn.QueryRow(query, currency.Name)

	switch err := row.Scan(&currencyRes.ID, &currencyRes.Name); err {
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
		err := rows.Scan(&cur.ID, cur.Name)
		if err != nil {
			return curs, err
		}
		curs = append(curs, cur)
	}
	return curs, err
}

func (db Database) DeleteCurrencyByName(currency models.Currency) error {
	query := `DELETE FROM currency WHERE name = $1;`
	_, err := db.Conn.Exec(query, currency.Name)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
