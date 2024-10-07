package postgredb

import "fmt"

func CreateTableDB() {
	createDB := `CREATE DATABASE api;`

	// todo
	// 1. Добавить дату создания

	createTableCustomers := `
	CREATE TABLE customers(
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		middle_name VARCHAR(255));`

	createTableTransactions := `
	CREATE TABLE internal_transactions(
		id SERIAL PRIMARY KEY,
		account_sender VARCHAR(255) NOT NULL,
		account_recipient VARCHAR(255) NOT NULL,
		amount VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL);`

	fmt.Print(createDB, createTableCustomers, createTableTransactions)
}
