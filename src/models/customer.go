package models

type Customer struct {
	ID    int    `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
