package postgredb

import (
	"api-go/pkg/models"
	"log"
)

// Получение пользователя из БД по id
func (db *DB) GetAccountBalance(id string) (*models.Balance, error) {
	balance := &models.Balance{}
	if err := db.conn.QueryRow(`
	SELECT account, amount	FROM account_balance WHERE id = $1;`, id).Scan(
		&balance.Account,
		&balance.Amount); err != nil {
		log.Print(err)
		return nil, err
	}
	return balance, nil
}

// Запись пользователя в БД
func (db *DB) SaveAccountBalance(balance *models.Balance) error {
	if _, err := db.conn.Exec(`
	INSERT INTO account_balance (account, amount, created_at) 
	VALUES (
		$1, --account
		$2,--amount
		$3); --Created_at `,
		balance.Account,
		balance.Amount,
		balance.CreatedAt); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (db *DB) UpdateAccountBalance(balance *models.Balance) error {
	if _, err := db.conn.Exec(`
	UPDATE account_balance 
	SET amount = $1, updated_at = $2
	WHERE account = $3;`,
		balance.Amount,
		balance.UpdatedAt,
		balance.Account); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
