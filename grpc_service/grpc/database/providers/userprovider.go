package providers

import (
	"database/sql"
	"fmt"
	"github.com/Andrew161644/full_app/grpc_service/grpc/database/models"
)

var ErrNoMatch = fmt.Errorf("no matching record")

func (db Database) AddUser(user models.UserModel) (int, error) {
	var id int
	query := `INSERT INTO Users (name, roleId) values ($1, $2) RETURNING id`
	err := db.Conn.QueryRow(query, user.Name, user.RoleId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db Database) GetUserById(userId int) (models.UserModel, error) {
	user := models.UserModel{}
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

func (db Database) GetUserByRole(roleId int) (models.UserList, error) {
	users := models.UserList{}

	rows, err := db.Conn.Query("SELECT * FROM users WHERE roleId = $1;", roleId)
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
