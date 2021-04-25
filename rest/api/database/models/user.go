package models

// Сущность - пользователь в базе
type UserModel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleId   int    `json:"roleId"`
}

// Вспомогательная структура - список пользователей
type UserList struct {
	Users []UserModel `json:"users"`
}
