package models

type Item struct {
	ID          int     `json:"_id"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}
