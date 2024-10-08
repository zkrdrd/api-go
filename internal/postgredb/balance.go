package postgredb

import (
	"api-go/pkg/models"
	"log"
)

// Получение пользователя из БД по id
func (db *DB) GetAccountBalance(id string) (*models.CacheOut, error) {
	cacheOut := &models.CacheOut{}
	if err := db.conn.QueryRow(`
	SELECT account, amount	FROM account_balance WHERE id = $1;`, id).Scan(
		&cacheOut.Account,
		&cacheOut.Amount); err != nil {
		log.Print(err)
		return nil, err
	}
	return cacheOut, nil
}

// Запись пользователя в БД
func (db *DB) SaveAccountBalance(cacheIn *models.CacheIn) error {
	if _, err := db.conn.Exec(`
	INSERT INTO account_balance (account, amount, middle_name) 
	VALUES (
	$1, --FirstName
	$2) --LastName
	ON CONFLICT DO UPDATE SET 
	amount = EXCLUDED.amount;  `,
		cacheIn.Account,
		cacheIn.Amount); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
