package providers

import (
	"database/sql"
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
)

func (db Database) AddYear(year models.Year) (int, error) {
	var id int
	query := `INSERT INTO Years (name,description) values ($1, $2) RETURNING id`
	err := db.Conn.QueryRow(query, year.Name, year.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db Database) GetYearByName(name string) (models.Year, error) {
	year := models.Year{}
	query := `SELECT * FROM years WHERE name = $1;`
	row := db.Conn.QueryRow(query, name)

	switch err := row.Scan(&year.ID, &year.Name, &year.Description); err {
	case sql.ErrNoRows:
		return year, ErrNoMatch
	default:
		return year, err
	}
}

func (db Database) GetAllYears() ([]models.Year, error) {
	query := `SELECT * FROM years`
	rows, err := db.Conn.Query(query)
	var yrs []models.Year
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var yr models.Year
		err := rows.Scan(&yr.ID, &yr.Name, &yr.Description)
		if err != nil {
			return yrs, err
		}
		yrs = append(yrs, yr)
	}
	return yrs, err
}

func (db Database) DeleteYear(year models.Year) error {
	query := `DELETE FROM years WHERE id = $1;`
	_, err := db.Conn.Exec(query, year.ID)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
