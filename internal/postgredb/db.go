package postgredb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
	SSLmode  string `json:"sslmode,omitempty"`
}

type DB struct {
	conn *sql.DB
}

const (
	createDB = `CREATE DATABASE api;`

	createTableCustomers = `
	CREATE TABLE customers(
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		middle_name VARCHAR(255));`

	createTableInternalTransactions = `
	CREATE TABLE internal_transactions(
		id SERIAL PRIMARY KEY,
		account_sender VARCHAR(255) NOT NULL,
		account_recipient VARCHAR(255) NOT NULL,
		amount VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL);`

	createTableAccountBalance = `
	CREATE TABLE account_balance(
		id SERIAL PRIMARY KEY,
		account int NOT NULL UNIQUE,
		amount VARCHAR(255) NOT NULL);`

	dropTableCustomers            = `DROP TABLE customers;`
	dropTableInternalTransactions = `DROP TABLE internal_transactions;`
	dropTableAccountBalance       = `DROP TABLE account_balance;`
)

// Инициализация соединения с БД
func (dbconf *DBConfig) NewDB() (*DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf(`host=%v 
	port=%v 
	user=%v 
	password=%v 
	dbname=%v 
	sslmode=%v`,
			dbconf.Host,
			dbconf.Port,
			dbconf.User,
			dbconf.Password,
			dbconf.DBname,
			dbconf.SSLmode))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &DB{conn: db}, nil
}

func CreateDB() {
	fmt.Print(createDB, createTableCustomers, createTableInternalTransactions)
}

// Пересоздание табилцы iternal_transaction
func (db *DB) RecreateTableInternalTransactions() {
	db.conn.Exec(dropTableInternalTransactions)
	db.conn.Exec(createTableInternalTransactions)
}

// Пересоздание табилцы account_balacnce
func (db *DB) RecreateTableAccountBalance() {
	db.conn.Exec(dropTableAccountBalance)
	db.conn.Exec(createTableAccountBalance)
}

// Пересоздание табилцы customers
func (db *DB) RecreateTableCustomers() {
	db.conn.Exec(dropTableCustomers)
	db.conn.Exec(createTableCustomers)
}

// Удаление базы данных
func (db *DB) DeleteDatabase() {
	db.conn.Query(`DROP DATABASE api;`)
}
