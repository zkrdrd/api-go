package postgredb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/zkrdrd/api-go/pkg/models"

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

type Connecter interface {
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Storage interface {
	AsTx(ctx context.Context, fn func(Storage) error) error
	CountInternalTransactions() (int, error)
	CreateDB() error
	DeleteDatabase() error
	GetAccountBalance(id string) (*models.Balance, error)
	GetInternalTrasaction(id string) (*models.Transactions, error)
	GetUser(id string) (*models.Users, error)
	ListInternalTransaction(filt *filter) ([]*models.Transactions, error)
	RecreateTableAccountBalance() error
	RecreateTableCustomers() error
	RecreateTableInternalTransactions() error
	Rollback() error
	SaveAccountBalance(balance *models.Balance) error
	SaveInternalTransaction(transf *models.Transactions) error
	SaveUser(user *models.Users) error
	UpdateAccountBalance(balance *models.Balance) error
}

type DB struct {
	conn   *sql.DB
	connTx *sql.Tx

	isTx bool
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
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		transaction_type VARCHAR(255) NOT NULL);`

	createTableAccountBalance = `
	CREATE TABLE account_balance(
		id SERIAL PRIMARY KEY,
		account int NOT NULL UNIQUE,
		amount VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE,
		updated_at TIMESTAMP WITH TIME ZONE);`

	dropTableCustomers            = `DROP TABLE customers;`
	dropTableInternalTransactions = `DROP TABLE internal_transactions;`
	dropTableAccountBalance       = `DROP TABLE account_balance;`
	dropDB                        = `DROP DATABASE api;`
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

	return &DB{conn: db, isTx: false}, nil
}

// Новая транзакция
func (db *DB) NewDBTx(conn *sql.Tx) *DB {
	if db.isTx {
		return db
	}
	return &DB{connTx: conn, isTx: true}
}

func (db *DB) useConn() Connecter {
	if db.isTx {
		return db.connTx
	}
	return db.conn
}

func (db *DB) Commit() error {
	if !db.isTx {
		return nil
	}
	return db.connTx.Commit()
}

func (db *DB) Rollback() error {
	if !db.isTx {
		return nil
	}

	if err := db.connTx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	return nil
}

func (db *DB) AsTx(ctx context.Context, fn func(Storage) error) error {
	if db.isTx {
		return fn(db)
	}

	conn, err := db.conn.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return fmt.Errorf(`as tx begin: %w`, err)
	}

	txDB := db.NewDBTx(conn)

	if err := fn(txDB); err != nil {
		err = fmt.Errorf(`as tx fn: %w`, err)

		if rollbackErr := txDB.Rollback(); rollbackErr != nil {
			err = fmt.Errorf(`rollback: %w`, err)
		}
		return err
	}

	return txDB.Commit()
}

func (db *DB) CreateDB() error {
	if _, err := db.conn.Exec(createDB); err != nil {
		return err
	}
	if _, err := db.conn.Exec(createTableCustomers); err != nil {
		return err
	}
	if _, err := db.conn.Exec(createTableInternalTransactions); err != nil {
		return err
	}
	if _, err := db.conn.Exec(createTableAccountBalance); err != nil {
		return err
	}
	return nil
}

// Пересоздание табилцы iternal_transaction
func (db *DB) RecreateTableInternalTransactions() error {
	if _, err := db.conn.Exec(dropTableInternalTransactions); err != nil {
		return err
	}
	if _, err := db.conn.Exec(createTableInternalTransactions); err != nil {
		return err
	}
	return nil
}

// Пересоздание табилцы account_balacnce
func (db *DB) RecreateTableAccountBalance() error {
	if _, err := db.conn.Exec(dropTableAccountBalance); err != nil {
		return err
	}
	if _, err := db.conn.Exec(createTableAccountBalance); err != nil {
		return err
	}
	return nil
}

// Пересоздание табилцы customers
func (db *DB) RecreateTableCustomers() error {
	if _, err := db.conn.Exec(dropTableCustomers); err != nil {
		return err
	}
	if _, err := db.conn.Exec(createTableCustomers); err != nil {
		return err
	}
	return nil
}

// Удаление базы данных
func (db *DB) DeleteDatabase() error {
	if _, err := db.conn.Exec(dropDB); err != nil {
		return err
	}
	return nil
}
