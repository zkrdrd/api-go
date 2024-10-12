package business

import (
	"api-go/internal/postgre"
	"api-go/pkg/models"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Accouting struct {
	db *postgre.DB
}

var (
	ErrMoneyNotEnough = errors.New(`error money not enough`)
)

// return time as string format RFC3339 "2006-01-02T15:04:05Z07:00"
func dateTime() string {
	return time.Now().Format(time.RFC3339)
}

// TODO:
// 1. Блокирую баланс
// 2. Разблокирую баланс
// Пополняю счет наличными
func (a *Accouting) CashOut(ctx context.Context, cacheOut *models.CashOut) error {
	accountSender, err := a.db.GetAccountBalance(cacheOut.Account)
	if err != nil {
		return err
	}

	amount, _ := strconv.ParseFloat(cacheOut.Amount, 32)
	senderBalance, _ := strconv.ParseFloat(accountSender.Amount, 32)

	if amount > senderBalance {
		return ErrMoneyNotEnough
	}

	accountSender.Amount = fmt.Sprintf("%.2f", senderBalance-amount)
	accountSender.UpdatedAt = dateTime()

	transaction := &models.Transactions{
		AccountSender:    accountSender.Account,
		AccountRecipient: accountSender.Account,
		Amount:           accountSender.Amount,
		CreatedAt:        dateTime(),
		TransactionType:  "Cash out",
	}

	err = a.db.AsTx(ctx, func(s postgre.Storage) error {
		if err := a.db.UpdateAccountBalance(accountSender); err != nil {
			return err
		}
		if err := a.db.SaveInternalTransaction(transaction); err != nil {
			return err
		}

		return nil
	})

	return err
}

// Тут я снимаю со счета начличные
func (a *Accouting) CashIn(ctx context.Context, cacheIn *models.CashIn) error {

	a.DB.StartTransaction()

	accountRecipient, err := a.DB.GetAccountBalance(cacheIn.Account)
	if err != nil {
		a.DB.RollBackTransaction()
		return err
	}

	amount, _ := strconv.ParseFloat(cacheIn.Amount, 32)
	recipientBalance, _ := strconv.ParseFloat(accountRecipient.Amount, 32)

	accountRecipient.Amount = fmt.Sprintf("%.2f", recipientBalance+amount)
	accountRecipient.UpdatedAt = dateTime()

	if err := a.DB.UpdateAccountBalance(accountRecipient); err != nil {
		a.DB.RollBackTransaction()
		return err
	}

	transaction := &models.Transactions{
		AccountSender:    accountRecipient.Account,
		AccountRecipient: accountRecipient.Account,
		Amount:           accountRecipient.Amount,
		CreatedAt:        dateTime(),
		TransactionType:  "Cash in",
	}

	if err := a.DB.SaveInternalTransaction(transaction); err != nil {
		a.DB.RollBackTransaction()
		return err
	}

	a.DB.CommitTransaction()
	return nil
}

// Тут я перевожу деньги между внетренними счетами
func (a *Accouting) InternalTransfer(ctx context.Context, transfer *models.InternalTranser) error {

	a.DB.StartTransaction()

	accountSender, err := a.DB.GetAccountBalance(transfer.AccountSender)
	if err != nil {
		log.Print(err)
		return err
	}

	amount, _ := strconv.ParseFloat(transfer.Amount, 32)
	senderBalance, _ := strconv.ParseFloat(accountSender.Amount, 32)

	if amount > senderBalance {
		a.DB.RollBackTransaction()
		return ErrMoneyNotEnough
	}

	accountSender.Amount = fmt.Sprintf("%.2f", senderBalance-amount)
	accountSender.UpdatedAt = dateTime()

	if err := a.DB.UpdateAccountBalance(accountSender); err != nil {
		a.DB.RollBackTransaction()
		return err
	}

	accountRecipient, err := a.DB.GetAccountBalance(transfer.AccountRecipient)
	if err != nil {
		a.DB.RollBackTransaction()
		return err
	}

	recipientBalance, _ := strconv.ParseFloat(accountRecipient.Amount, 32)

	accountRecipient.Amount = fmt.Sprintf("%.2f", recipientBalance+amount)
	accountRecipient.UpdatedAt = dateTime()

	if err := a.DB.UpdateAccountBalance(accountRecipient); err != nil {
		a.DB.RollBackTransaction()
		return err
	}

	transaction := &models.Transactions{
		AccountSender:    accountSender.Account,
		AccountRecipient: accountRecipient.Account,
		Amount:           transfer.Amount,
		CreatedAt:        dateTime(),
		TransactionType:  "Transfer",
	}

	if err := a.DB.SaveInternalTransaction(transaction); err != nil {
		a.DB.RollBackTransaction()
		return err
	}

	a.DB.CommitTransaction()

	return nil
}
