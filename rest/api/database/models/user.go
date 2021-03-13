package models

type UserModel struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	RoleId int    `json:"roleId"`
}

type UserList struct {
	Users []UserModel `json:"users"`
}
