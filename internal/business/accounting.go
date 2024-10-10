package business

import (
	"api-go/internal/postgredb"
	"api-go/pkg/models"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

// Я знаю как делать операции с счетами пользователя
type Accouting struct {
	DB *postgredb.DB
	// Во мне лежит все  необходимое для работы
	// к примеру подключение к БД, а возможно и подключения
	// к другим сервисам
	//db *db.Conn
}

// return time as string format RFC3339 "2006-01-02T15:04:05Z07:00"
func dateTime() string {
	return time.Now().Format(time.RFC3339)
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
			Account:   cacheOut.Account,
			Amount:    fmt.Sprintf("%.2f", dbquery-query),
			UpdatedAt: dateTime(),
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
			Account:   cacheIn.Account,
			Amount:    fmt.Sprintf("%.2f", dbquery+query),
			UpdatedAt: dateTime(),
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
	accountSender, err := a.DB.GetAccountBalance(transfer.AccountSender)
	if err != nil {
		log.Print(err)
		return err
	}

	senderAmount, _ := strconv.ParseFloat(transfer.Amount, 32)
	senderBalance, _ := strconv.ParseFloat(accountSender.Amount, 32)

	if senderAmount > senderBalance {
		return fmt.Errorf(`error money not enough`)
	} else {

		newBalance := &models.Balance{
			Account:   accountSender.Account,
			Amount:    fmt.Sprintf("%.2f", senderBalance-senderAmount),
			UpdatedAt: dateTime(),
		}

		if err := a.DB.UpdateAccountBalance(newBalance); err != nil {
			log.Fatal(err)
			return err
		}
	}

	accountRecipient, err := a.DB.GetAccountBalance(transfer.AccountRecipient)
	if err != nil {
		log.Print(err)
		return err
	}

	recipientAmount, _ := strconv.ParseFloat(transfer.Amount, 32)
	recipientBalance, _ := strconv.ParseFloat(accountRecipient.Amount, 32)

	newBalance := &models.Balance{
		Account:   accountRecipient.Account,
		Amount:    fmt.Sprintf("%.2f", recipientBalance+recipientAmount),
		UpdatedAt: dateTime(),
	}

	if err := a.DB.UpdateAccountBalance(newBalance); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
