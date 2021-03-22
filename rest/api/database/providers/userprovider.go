package providers

import (
	"database/sql"
	"fmt"
	"github.com/Andrew161644/avicks_laba/api/database/models"
)

// Провайдер для сущности - пользователь

var ErrNoMatch = fmt.Errorf("no matching record")

func (db Database) AddUser(user models.UserModel) (int, error) {
	var id int
	query := `INSERT INTO Users (name, password, roleId) values ($1, $2, $3) RETURNING id`
	err := db.Conn.QueryRow(query, user.Name, user.Password, user.RoleId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db Database) GetUserById(userId int) (models.UserModel, error) {
	user := models.UserModel{}
	query := `SELECT * FROM users WHERE id = $1;`
	row := db.Conn.QueryRow(query, userId)

	switch err := row.Scan(&user.ID, &user.Name, &user.Password, &user.RoleId); err {
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

func (db Database) UpdateUser(userId int, user models.UserModel) (models.UserModel, error) {
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

func (db Database) GetUserByRole(roleId int) (models.List, error) {
	users := models.List{}

	rows, err := db.Conn.Query("SELECT id, name, roleId FROM users WHERE roleId = $1;", roleId)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user models.UserModel
		err := rows.Scan(&user.ID, &user.Name, &user.RoleId)
		if err != nil {
			return users, err
		}
		users.Users = append(users.Users, user)
	}
	return users, nil
}

func (db Database) GetUserByNameAndPassword(name string, password string) (models.UserModel, error) {
	user := models.UserModel{}
	query := `SELECT * FROM users WHERE name = $1 and password = $2;`
	row := db.Conn.QueryRow(query, name, password)

	switch err := row.Scan(&user.ID, &user.Name, &user.Password, &user.RoleId); err {
	case sql.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
}

func (db Database) GetUserByName(name string) (models.UserModel, error) {
	user := models.UserModel{}
	query := `SELECT * FROM users WHERE name = $1`
	row := db.Conn.QueryRow(query, name)

	switch err := row.Scan(&user.ID, &user.Name, &user.Password, &user.RoleId); err {
	case sql.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
}
