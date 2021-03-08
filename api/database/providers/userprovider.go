package providers

import (
	"database/sql"
	"fmt"
	. "github.com/Andrew161644/avicks_laba/api/database/models"
)

var ErrNoMatch = fmt.Errorf("no matching record")

func (db Database) AddUser(user User) (int, error) {
	var id int
	query := `INSERT INTO Users (name, roleId) values ($1, $2) RETURNING id`
	err := db.Conn.QueryRow(query, user.Name, user.RoleId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db Database) GetUserById(userId int) (User, error) {
	user := User{}
	query := `SELECT * FROM users WHERE id = $1;`
	row := db.Conn.QueryRow(query, userId)

	switch err := row.Scan(&user.ID, &user.Name, &user.RoleId); err {
	case sql.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
}

func (db Database) DeleteUsers(userId int) error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := db.Conn.Exec(query, userId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateUser(userId int, user User) (User, error) {
	query := `UPDATE users SET name=$1, roleId=$2 WHERE id=$3 RETURNING id, name, roleId;`
	err := db.Conn.QueryRow(query, user.Name, user.RoleId, userId).Scan(&user.ID, &user.Name, &user.RoleId)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrNoMatch
		}
		return user, err
	}

	return user, nil
}
