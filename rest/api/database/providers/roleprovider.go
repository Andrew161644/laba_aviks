package providers

import (
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
)

func (db Database) GetAllRoles() ([]models.Role, error) {
	query := `SELECT * FROM Roles`
	rows, err := db.Conn.Query(query)
	var roles []models.Role
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.RoleName)
		if err != nil {
			return roles, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
