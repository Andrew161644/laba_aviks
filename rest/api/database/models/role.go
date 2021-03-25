package models

// Сущность - роль в базе
type Role struct {
	ID       int    `json:"id"`
	RoleName string `json:"roleName"`
}
