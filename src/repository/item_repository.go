package repository

import "github.com/wbluan/orders-api/src/models"

func GetItems() ([]models.Item, error) {
	collection, err := DB.Query("SELECT id, description, price, quantity FROM items")
	if err != nil {
		return nil, err
	}

	defer collection.Close()

	var items []models.Item
	for collection.Next() {
		var item models.Item
		if err := collection.Scan(&item.ID, &item.Description, &item.Price, &item.Quantity); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func CreateItem(item models.Item) (int, error) {
	insertOne, err := DB.Prepare("INSERT INTO items(description, price, quantity) VALUES (?, ?, ?)")
	if err != nil {
		return 0, nil
	}

	res, err := insertOne.Exec(item.Description, item.Price, item.Quantity)
	if err != nil {
		return 0, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}
