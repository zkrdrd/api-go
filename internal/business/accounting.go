package business

import (
	"api-go/internal/postgredb"
	"api-go/pkg/models"
	"context"
	"fmt"
	"log"
	"strconv"
)

// Я знаю как делать операции с счетами пользователя
type Accouting struct {
	DB *postgredb.DB
	// Во мне лежит все  необходимое для работы
	// к примеру подключение к БД, а возможно и подключения
	// к другим сервисам
	//db *db.Conn
}

// Тут я пополняю счет наличными
func (a *Accouting) CacheOut(ctx context.Context, cacheOut *models.CacheOut) error {
	// TODO:
	// 1. Блокирую баланс
	// 2. Разблокирую баланс

	// Обналичиваю средства
	// Изменяю сумму баланса
	getFromDB, err := a.DB.GetAccountBalance(cacheOut.Account)
	if err != nil {
		log.Print(err)
		return err
	}

	query, _ := strconv.ParseFloat(cacheOut.Amount, 32)
	dbquery, _ := strconv.ParseFloat(getFromDB.Amount, 32)

	if query > dbquery {
		return fmt.Errorf(`error money not enough`)
	} else {
		cahceInNew := &models.Balance{
			Account: cacheOut.Account,
			Amount:  fmt.Sprintf("%.2f", dbquery-query),
		}

		if err := a.DB.UpdateAccountBalance(cahceInNew); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

// Тут я снимаю со счета начличные
func (a *Accouting) CacheIn(ctx context.Context, cacheIn *models.CacheIn) error {
	getFromDB, err := a.DB.GetAccountBalance(cacheIn.Account)
	if err != nil {
		log.Print(err)
		return err
	}
	query, _ := strconv.ParseFloat(cacheIn.Amount, 32)
	dbquery, _ := strconv.ParseFloat(getFromDB.Amount, 32)

	if query > dbquery {
		return fmt.Errorf(`error money not enough`)
	} else {

		cahceInNew := &models.Balance{
			Account: cacheIn.Account,
			Amount:  fmt.Sprintf("%.2f", dbquery+query),
		}

		if err := a.DB.UpdateAccountBalance(cahceInNew); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

// Тут я перевожу деньги между внетренними счетами
func (a *Accouting) InternalTransfer(ctx context.Context, transfer *models.InternalTransaction) error {
	cacheIn, cacheOut := a.SeparationInternalTransactionToCahceInOut(transfer)
	a.ProcessingAccountBalance(cacheIn, cacheOut)
	return nil
}

// Processing AccountBalance
func (a *Accouting) ProcessingAccountBalance(cacheIn *models.CacheIn, cacheOut *models.CacheOut) error {

	if cacheIn.Account != "" && cacheIn.Amount != "" {
		a.CacheIn(context.Background(), cacheIn)
	}
	if cacheIn.Account != "" && cacheIn.Amount != "" {
		a.CacheOut(context.Background(), cacheOut)
	}

	return nil
}

// разделение модели InternalTransation на модели CacheIn CacheOut
func (a *Accouting) SeparationInternalTransactionToCahceInOut(internalTransactions *models.InternalTransaction) (*models.CacheIn, *models.CacheOut) {

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
