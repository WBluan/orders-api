package models

type Order struct {
	ID         int    `json:"_id"`
	CustomerID int    `json:"customer_id"`
	ProductID  int    `json:"product_id"`
	OrderDate  string `json:"order_date"`
}
