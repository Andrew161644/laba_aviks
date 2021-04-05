package providers

import (
	"database/sql"
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
)

// Провайдер для сущности - роль
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

func (db Database) AddRole(role models.Role) (int, error) {
	var id int
	query := `INSERT INTO roles (roleName) values ($1) RETURNING id`
	err := db.Conn.QueryRow(query, role.RoleName).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db Database) DeleteRole(roleId int) error {
	query := `DELETE FROM roles WHERE id = $1;`
	_, err := db.Conn.Exec(query, roleId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) GetRoleByName(role models.Role) (models.Role, error) {
	resRole := models.Role{}
	query := `SELECT * FROM roles WHERE rolename = $1;`
	row := db.Conn.QueryRow(query, role.RoleName)

	switch err := row.Scan(&resRole.ID, &resRole.RoleName); err {
	case sql.ErrNoRows:
		return resRole, ErrNoMatch
	default:
		return resRole, err
	}
}
