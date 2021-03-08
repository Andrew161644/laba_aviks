package models

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	RoleId int    `json:"roleId"`
}

type UserList struct {
	Users []User `json:"users"`
}
