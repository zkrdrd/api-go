package postgredb

import (
	"api-go/pkg/models"
	"errors"
	"log"
)

var (
	ErrNoMoreResults = errors.New("no more results")
)

// Получение транзакции по id
func (db *DB) GetInternalTrasaction(id string) (*models.InternalTransaction, error) {
	transf := &models.InternalTransaction{}
	if err := db.conn.QueryRow(`
	SELECT account_sender, account_recipient, amount 
	FROM transactions WHERE id = $1;`, id).Scan(
		&transf.AccountRecipient,
		&transf.AccountSender,
		&transf.Amount); err != nil {
		log.Print(err)
		return nil, err
	}
	return transf, nil
}

// Получение всех транзакций из БД в slice
func (db *DB) ListInternalTransaction() ([]*models.InternalTransaction, error) {
	// TODO:
	// 1. order by по дате создания
	transfSlice := []*models.InternalTransaction{}
	rows, err := db.conn.Query(`
	SELECT account_sender, account_recipient, amount 
	FROM transactions;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		transf := &models.InternalTransaction{}
		if err := rows.Scan(&transf.AccountSender, &transf.AccountRecipient, &transf.Amount); err != nil {
			log.Fatal(err)
		}
		transfSlice = append(transfSlice, transf)
	}
	return transfSlice, nil
}

// Запись транзакций в БД
func (db *DB) SaveInternalTransaction(transf *models.InternalTransaction) error {
	// todo
	// 1. добавить дату создания
	// 2. изменение баланса
	if _, err := db.conn.Exec(`
	INSERT INTO transactions (account_sender, account_recipient, amount) 
	VALUES (
	$1, --AccountSender
    $2, --AccountRecipient
    $3); --Amount`,
		transf.AccountSender,
		transf.AccountRecipient,
		transf.Amount); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
