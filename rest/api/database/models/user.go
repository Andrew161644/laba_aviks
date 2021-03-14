package models

type UserModel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleId   int    `json:"roleId"`
}

type List struct {
	Users []UserModel `json:"users"`
}
