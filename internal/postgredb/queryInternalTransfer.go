package postgredb

import (
	"api-go/internal/business"
	"log"
)

func (db *DB) GetTransfer(id string, transf *business.InternalTransfer) (*business.InternalTransfer, error) {
	if err := db.conn.QueryRow(`
	SELECT account_sender, account_recipient, amount 
	FROM transaction WHERE id = $1`, id).Scan(
		&transf.AccountRecipient,
		&transf.AccountSender,
		&transf.Amount); err != nil {
		log.Print(err)
		return nil, err
	}
	return transf, nil
}

func (db *DB) ListTransfer(transf *business.InternalTransfer) error {
	rows, err := db.conn.Query(`
	SELECT account_sender, account_recipient, amount 
	FROM transaction`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&transf.AccountSender, &transf.AccountRecipient, &transf.Amount); err != nil {
			log.Fatal(err)
		}
		log.Printf("Account sender: %v; Account recipient: %v; Amount: %v\n", &transf.AccountSender, &transf.AccountRecipient, &transf.Amount)
	}
	if !rows.NextResultSet() {
		log.Fatalf("expected more result sets: %v", rows.Err())
	}
	return nil
}

func (db *DB) SaveTransfer(transf *business.InternalTransfer) error {
	if _, err := db.conn.Exec(`
	INSERT INTO transaction (account_sender, account_recipient, amount) 
	VALUES (
	$1, -- AccountSender
    $2, -- AccountRecipient
    $3  -- Amount)`,
		transf.AccountRecipient,
		transf.AccountSender,
		transf.Amount); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
