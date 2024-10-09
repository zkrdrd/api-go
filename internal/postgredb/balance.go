package postgredb

import (
	"api-go/pkg/models"
	"fmt"
	"log"
	"strconv"
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
	INSERT INTO account_balance (account, amount) 
	VALUES (
		$1, --account
		$2) --amount
	ON CONFLICT (account) DO UPDATE SET 
		amount = EXCLUDED.amount;`,
		cacheIn.Account,
		cacheIn.Amount); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// Увеличение значения баланса
func (db *DB) AddAccountBalance(cacheInQuery *models.CacheIn) error {
	getFromDB, err := db.GetAccountBalance(cacheInQuery.Account)
	if err != nil {
		log.Print(err)
		return err
	}
	query, _ := strconv.ParseFloat(cacheInQuery.Amount, 32)
	dbquery, _ := strconv.ParseFloat(getFromDB.Amount, 32)

	if query > dbquery {
		return fmt.Errorf(`error money not enough`)
	} else {

		cahceInNew := &models.CacheIn{
			Account: cacheInQuery.Account,
			Amount:  fmt.Sprintf("%.2f", dbquery+query),
		}

		if err := db.SaveAccountBalance(cahceInNew); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

// Уменьшение значения баланса
func (db *DB) SubtractAccountBalance(cacheOutQuery *models.CacheOut) error {
	getFromDB, err := db.GetAccountBalance(cacheOutQuery.Account)
	if err != nil {
		log.Print(err)
		return err
	}

	query, _ := strconv.ParseFloat(cacheOutQuery.Amount, 32)
	dbquery, _ := strconv.ParseFloat(getFromDB.Amount, 32)

	if query > dbquery {
		return fmt.Errorf(`error money not enough`)
	} else {
		cahceInNew := &models.CacheIn{
			Account: cacheOutQuery.Account,
			Amount:  fmt.Sprintf("%.2f", dbquery-query),
		}

		if err := db.SaveAccountBalance(cahceInNew); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

// разделение модели InternalTransation на модели CacheIn CacheOut
func SeparationInternalTransactionToCahceInOut(internalTransactions *models.InternalTransaction) (*models.CacheIn, *models.CacheOut) {
	cacheOut := &models.CacheOut{
		Account: internalTransactions.AccountSender,
		Amount:  internalTransactions.Amount,
	}
	cacheIn := &models.CacheIn{
		Account: internalTransactions.AccountRecipient,
		Amount:  internalTransactions.Amount,
	}
	return cacheIn, cacheOut
}
