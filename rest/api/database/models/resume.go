package models

type Resume struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Speciality string `json:"speciality"`
	About      string `json:"about"`
}
