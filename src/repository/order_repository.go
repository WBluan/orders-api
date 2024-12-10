package repository

import "github.com/wbluan/orders-api/src/models"

func GetOrders() ([]models.Order, error) {
	collection, err := DB.Query("SELECT id, customer_id, product_id, order_date FROM orders")
	if err != nil {
		return nil, err
	}
	defer collection.Close()

	var orders []models.Order
	for collection.Next() {
		var order models.Order
		if err := collection.Scan(&order.ID, &order.CustomerID, &order.ProductID, &order.OrderDate); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func CreateOrder(order models.Order) (int, error) {
	insertOne, err := DB.Prepare("INSERT INTO orders(customer_id, product_id, order_date) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}

	res, err := insertOne.Exec(order.CustomerID, order.ProductID, order.OrderDate)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
