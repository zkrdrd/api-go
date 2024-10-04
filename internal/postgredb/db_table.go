package postgredb

import "fmt"

func CreateTableDB() {
	createDB := `CREATE DATABASE api;`

	createTableCustomers := `
	CREATE TABLE customers(
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		middle_name VARCHAR(255));`

	createTableTransactions := `
	CREATE TABLE transactions(
		id SERIAL PRIMARY KEY,
		account_sender VARCHAR(255) NOT NULL,
		account_recipient VARCHAR(255) NOT NULL,
		amount VARCHAR(255) NOT NULL);`

	fmt.Print(createDB, createTableCustomers, createTableTransactions)
}