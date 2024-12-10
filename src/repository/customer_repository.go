package repository

import "github.com/wbluan/orders-api/src/models"

func GetCustomers() ([]models.Customer, error) {
	collection, err := DB.Query("SELECT id, name, email FROM customers")
	if err != nil {
		return nil, err
	}
	defer collection.Close()

	var customers []models.Customer
	for collection.Next() {
		var customer models.Customer
		if err := collection.Scan(&customer.ID, &customer.Name, &customer.Email); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func CreateCustomer(customer models.Customer) (int, error) {
	insertOne, err := DB.Prepare("INSERT INTO customers(name, email) VALUES(?, ?)")
	if err != nil {
		return 0, nil
	}

	res, err := insertOne.Exec(customer.Name, customer.Email)
	if err != nil {
		return 0, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}
