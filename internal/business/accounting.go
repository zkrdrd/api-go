package business

import (
	"api-go/internal/postgredb"
	"api-go/pkg/models"
	"context"
	"errors"
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

var (
	ErrMoneyNotEnough = errors.New(`error money not enough`)
)

// return time as string format RFC3339 "2006-01-02T15:04:05Z07:00"
func dateTime() string {
	return time.Now().Format(time.RFC3339)
}

// Тут я пополняю счет наличными
func (a *Accouting) CashOut(ctx context.Context, cacheOut *models.CashOut) error {

	if err := a.DB.StartTransaction(); err != nil {
		return err
	}
	// TODO:
	// 1. Блокирую баланс
	// 2. Разблокирую баланс

	// Обналичиваю средства
	// Изменяю сумму баланса
	accountSender, err := a.DB.GetAccountBalance(cacheOut.Account)
	if err != nil {
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
		return err
	}

	amount, _ := strconv.ParseFloat(cacheOut.Amount, 32)
	senderBalance, _ := strconv.ParseFloat(accountSender.Amount, 32)

	if amount > senderBalance {
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
		return ErrMoneyNotEnough
	}

	accountSender.Amount = fmt.Sprintf("%.2f", senderBalance-amount)
	accountSender.UpdatedAt = dateTime()

	if err := a.DB.UpdateAccountBalance(accountSender); err != nil {
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
		return err
	}

	transaction := &models.Transactions{
		AccountSender:    accountSender.Account,
		AccountRecipient: accountSender.Account,
		Amount:           accountSender.Amount,
		CreatedAt:        dateTime(),
		TransactionType:  "Cash out",
	}

	if err := a.DB.SaveInternalTransaction(transaction); err != nil {
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
		return err
	}

	if err := a.DB.CommitTransaction(); err != nil {
		return err
	}
	return nil
}

// Тут я снимаю со счета начличные
func (a *Accouting) CashIn(ctx context.Context, cacheIn *models.CashIn) error {

	if err := a.DB.StartTransaction(); err != nil {
		return err
	}

	accountRecipient, err := a.DB.GetAccountBalance(cacheIn.Account)
	if err != nil {
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
		return err
	}

	amount, _ := strconv.ParseFloat(cacheIn.Amount, 32)
	recipientBalance, _ := strconv.ParseFloat(accountRecipient.Amount, 32)

	accountRecipient.Amount = fmt.Sprintf("%.2f", recipientBalance+amount)
	accountRecipient.UpdatedAt = dateTime()

	if err := a.DB.UpdateAccountBalance(accountRecipient); err != nil {
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
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
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
		return err
	}

	if err := a.DB.CommitTransaction(); err != nil {
		return err
	}
	return nil
}

// Тут я перевожу деньги между внетренними счетами
func (a *Accouting) InternalTransfer(ctx context.Context, transfer *models.InternalTranser) error {

	if err := a.DB.StartTransaction(); err != nil {
		return err
	}

	accountSender, err := a.DB.GetAccountBalance(transfer.AccountSender)
	if err != nil {
		log.Print(err)
		return err
	}

	amount, _ := strconv.ParseFloat(transfer.Amount, 32)
	senderBalance, _ := strconv.ParseFloat(accountSender.Amount, 32)

	if amount > senderBalance {
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
		return ErrMoneyNotEnough
	}

	accountSender.Amount = fmt.Sprintf("%.2f", senderBalance-amount)
	accountSender.UpdatedAt = dateTime()

	if err := a.DB.UpdateAccountBalance(accountSender); err != nil {
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
		return err
	}

	accountRecipient, err := a.DB.GetAccountBalance(transfer.AccountRecipient)
	if err != nil {
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
		return err
	}

	recipientBalance, _ := strconv.ParseFloat(accountRecipient.Amount, 32)

	accountRecipient.Amount = fmt.Sprintf("%.2f", recipientBalance+amount)
	accountRecipient.UpdatedAt = dateTime()

	if err := a.DB.UpdateAccountBalance(accountRecipient); err != nil {
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
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
		if err := a.DB.RollBackTransaction(); err != nil {
			return err
		}
		return err
	}

	if err := a.DB.CommitTransaction(); err != nil {
		return err
	}

	return nil
}
