package postgredb

import (
	"github.com/zkrdrd/api-go/pkg/models"
)

// Получение пользователя из БД по id
func (db *DB) GetAccountBalance(id string) (*models.Balance, error) {
	balance := &models.Balance{}
	if err := db.useConn().QueryRow(`
	SELECT account, amount, created_at	FROM account_balance WHERE id = $1;`, id).Scan(
		&balance.Account,
		&balance.Amount,
		&balance.CreatedAt); err != nil {
		return nil, err
	}
	return balance, nil
}

// Запись пользователя в БД
func (db *DB) SaveAccountBalance(balance *models.Balance) error {
	if _, err := db.useConn().Exec(`
	INSERT INTO account_balance (account, amount, created_at) 
	VALUES (
		$1, --account
		$2,--amount
		$3); --Created_at `,
		balance.Account,
		balance.Amount,
		balance.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdateAccountBalance(balance *models.Balance) error {
	if _, err := db.useConn().Exec(`
	UPDATE account_balance 
	SET amount = $1, updated_at = $2
	WHERE account = $3;`,
		balance.Amount,
		balance.UpdatedAt,
		balance.Account); err != nil {
		return err
	}
	return nil
}
