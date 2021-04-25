package providers

import (
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
)

func (db Database) AddResume(resume models.Resume) (int, error) {
	var id int
	query := `INSERT INTO resume (name, email, speciality, about)  values ($1, $2, $3, $4) RETURNING id`
	err := db.Conn.QueryRow(query, resume.Name, resume.Email, resume.Speciality, resume.About).Scan(&id)
	if err != nil {
		return 0, err
	}
	log.Println("Resume Added")
	return id, nil
}
