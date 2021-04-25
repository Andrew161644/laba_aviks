package providers

import (
	"database/sql"
	"github.com/Andrew161644/avicks_laba/api/database/models"
)

func (db Database) AddIndex(index models.Index) (int, error) {
	var id int
	query := `INSERT INTO indices (pastYear, currentYear, value, currencyId) values ($1, $2, $3, $4) RETURNING id`
	err := db.Conn.QueryRow(query, index.PastYear, index.CurrentYear, index.Value, index.CurrencyId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db Database) DeleteIndexById(index models.Index) error {
	query := `DELETE FROM indices WHERE id = $1;`
	_, err := db.Conn.Exec(query, index.Id)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) GetAllIndices() ([]models.Index, error) {
	var indices []models.Index
	query := `SELECT * FROM indices`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return indices, err
	}
	for rows.Next() {
		var index = models.Index{}
		err := rows.Scan(&index.Id, &index.PastYear, &index.CurrentYear, &index.Value, &index.CurrencyId)
		if err != nil {
			return indices, err
		}
		indices = append(indices, index)
	}
	return indices, err
}

func (db Database) GetAllIndicesByCurrentYear(index models.Index) ([]models.Index, error) {
	var indices []models.Index
	query := `SELECT * FROM indices WHERE currentyear = $1`
	rows, err := db.Conn.Query(query, index.CurrentYear)
	if err != nil {
		return indices, err
	}
	for rows.Next() {
		var index = models.Index{}
		err := rows.Scan(&index.Id, &index.PastYear, &index.CurrentYear, &index.Value, &index.CurrencyId)
		if err != nil {
			return indices, err
		}
		indices = append(indices, index)
	}
	return indices, err
}

func (db Database) GetAllIndicesByCurrentYearAndCurrency(index models.Index) ([]models.Index, error) {
	var indices []models.Index
	query := `SELECT * FROM indices WHERE currentyear = $1 AND currencyid = $2`
	rows, err := db.Conn.Query(query, index.CurrentYear, index.CurrencyId)
	if err != nil {
		return indices, err
	}
	for rows.Next() {
		var index = models.Index{}
		err := rows.Scan(&index.Id, &index.PastYear, &index.CurrentYear, &index.Value, &index.CurrencyId)
		if err != nil {
			return indices, err
		}
		indices = append(indices, index)
	}
	return indices, err
}
