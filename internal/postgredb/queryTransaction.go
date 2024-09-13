package postgredb

import (
	"api-go/internal/buisness"
	"log"
)

func (db *DB) GetTransfer(id, transf *buisness.Transfer) (*buisness.Transfer, error) {
	if err := db.Conn.QueryRow(`
	SELECT id account_sender, account_recipient, amount 
	FROM transaction WHERE id = ?`, id).Scan(
		&transf.AccountRecipient,
		&transf.AccountSender,
		&transf.Amount); err != nil {
		log.Print(err)
		return nil, err
	}
	return transf, nil
}

func (db *DB) ListTransfer(transf *buisness.Transfer) error {
	rows, err := db.Conn.Query(`
	SELECT id account_sender, account_recipient, amount 
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

func (db *DB) SaveTransfer(transf *buisness.Transfer) error {
	if _, err := db.Conn.Exec(`
	INSERT INTO transaction (account_sender, account_recipient, amount) 
	VALUES (
	?, -- AccountSender
    ?, -- AccountRecipient
    ?  -- Amount)`,
		transf.AccountRecipient,
		transf.AccountSender,
		transf.Amount); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
