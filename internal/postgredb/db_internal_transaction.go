package postgredb

import (
	"api-go/internal/business"
	"errors"
	"log"
)

var (
	ErrNoMoreResults = errors.New("no more results")
)

func (db *DB) GetInternalTrasaction(id string) (*business.InternalTransaction, error) {
	transf := &business.InternalTransaction{}
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

func (db *DB) ListInternalTransaction() error {
	transf := &business.InternalTransaction{}
	rows, err := db.conn.Query(`
	SELECT account_sender, account_recipient, amount 
	FROM transactions;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&transf.AccountSender, &transf.AccountRecipient, &transf.Amount); err != nil {
			log.Fatal(err)
		}
		log.Printf("Account sender: %v; Account recipient: %v; Amount: %v\n", transf.AccountSender, transf.AccountRecipient, transf.Amount)
	}

	return nil
}

func (db *DB) SaveInternalTransaction(transf *business.InternalTransaction) error {
	if _, err := db.conn.Exec(`
	INSERT INTO transactions (account_sender, account_recipient, amount) 
	VALUES (
	$1, --AccountSender
    $2, --AccountRecipient
    $3); --Amount`,
		transf.AccountRecipient,
		transf.AccountSender,
		transf.Amount); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
