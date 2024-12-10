package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "./orders.db")
	if err != nil {
		log.Fatal("Could not open database: ", err)
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS customers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		price FLOAT NOT NULL,
		quantity INTEGER NOT NULL
	);
		
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_id INTEGER NOT NULL,
		product_id INTEGER NOT NULL,
		order_date TEXT NOT NULL,
		FOREIGN KEY (customer_id) REFERENCES customers(id),
		FOREIGN KEY (product_id) REFERENCES items(id)
	);`

	_, err = DB.Exec(createTables)
	if err != nil {
		log.Fatal("Error to create tables: ", err)
	}
}
