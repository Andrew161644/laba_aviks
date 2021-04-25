package models

type Year struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type YearList struct {
	Years []Year `json:"years"`
}
