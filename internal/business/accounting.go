package business

import (
	"api-go/internal/locker"
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
	db *postgredb.DB
	// Во мне лежит все  необходимое для работы
	// к примеру подключение к БД, а возможно и подключения
	// к другим сервисам
	//db *db.Conn
	lock *locker.Locker
}

var (
	ErrMoneyNotEnough   = errors.New(`error money not enough`)
	ErrAccoutingIsEmpty = errors.New(`account is empty`)
)

func NewAccouting(dbConn *postgredb.DB, lock *locker.Locker) *Accouting {
	return &Accouting{
		db:   dbConn,
		lock: lock,
	}
}

// return time as string format RFC3339 "2006-01-02T15:04:05Z07:00"
func dateTime() string {
	return time.Now().Format(time.RFC3339)
}

// Тут я пополняю счет наличными
func (a *Accouting) CashOut(ctx context.Context, cacheOut *models.CashOut) error {
	if cacheOut.Account == `` {
		return ErrAccoutingIsEmpty
	}

	a.lock.Lock(cacheOut.Account)
	defer a.lock.Unlock(cacheOut.Account)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

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

	err = a.db.AsTx(ctx,
		func(tx postgredb.Storage) error {
			if err := tx.UpdateAccountBalance(accountSender); err != nil {
				return err
			}

			if err := tx.SaveInternalTransaction(transaction); err != nil {
				return err
			}
			return nil
		},
	)
	return fmt.Errorf(`cashout: %w`, err)
}

// Тут я снимаю со счета начличные
func (a *Accouting) CashIn(ctx context.Context, cacheIn *models.CashIn) error {
	if cacheIn.Account == `` {
		return ErrAccoutingIsEmpty
	}

	a.lock.Lock(cacheIn.Account)
	defer a.lock.Unlock(cacheIn.Account)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	accountRecipient, err := a.db.GetAccountBalance(cacheIn.Account)
	if err != nil {
		return err
	}

	amount, _ := strconv.ParseFloat(cacheIn.Amount, 32)
	recipientBalance, _ := strconv.ParseFloat(accountRecipient.Amount, 32)

	accountRecipient.Amount = fmt.Sprintf("%.2f", recipientBalance+amount)
	accountRecipient.UpdatedAt = dateTime()

	transaction := &models.Transactions{
		AccountSender:    accountRecipient.Account,
		AccountRecipient: accountRecipient.Account,
		Amount:           accountRecipient.Amount,
		CreatedAt:        dateTime(),
		TransactionType:  "Cash in",
	}

	err = a.db.AsTx(ctx,
		func(tx postgredb.Storage) error {
			if err := tx.UpdateAccountBalance(accountRecipient); err != nil {
				return err
			}

			if err := tx.SaveInternalTransaction(transaction); err != nil {
				return err
			}
			return nil
		},
	)

	return fmt.Errorf(`cashin: %w`, err)
}

// Тут я перевожу деньги между внетренними счетами
func (a *Accouting) InternalTransfer(ctx context.Context, transfer *models.InternalTranser) error {
	if transfer.AccountRecipient == `` || transfer.AccountSender == `` {
		return ErrAccoutingIsEmpty
	}

	a.lock.Lock(transfer.AccountSender)
	defer a.lock.Unlock(transfer.AccountSender)

	a.lock.Lock(transfer.AccountRecipient)
	defer a.lock.Unlock(transfer.AccountRecipient)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	accountSender, err := a.db.GetAccountBalance(transfer.AccountSender)
	if err != nil {
		log.Print(err)
		return err
	}

	//big.NewInt(0)
	amount, _ := strconv.ParseFloat(transfer.Amount, 32)
	senderBalance, _ := strconv.ParseFloat(accountSender.Amount, 32)

	if amount > senderBalance {
		return ErrMoneyNotEnough
	}

	accountSender.Amount = fmt.Sprintf("%.2f", senderBalance-amount)
	accountSender.UpdatedAt = dateTime()

	accountRecipient, err := a.db.GetAccountBalance(transfer.AccountRecipient)
	if err != nil {
		return err
	}

	recipientBalance, _ := strconv.ParseFloat(accountRecipient.Amount, 32)

	accountRecipient.Amount = fmt.Sprintf("%.2f", recipientBalance+amount)
	accountRecipient.UpdatedAt = dateTime()

	transaction := &models.Transactions{
		AccountSender:    accountSender.Account,
		AccountRecipient: accountRecipient.Account,
		Amount:           transfer.Amount,
		CreatedAt:        dateTime(),
		TransactionType:  "Transfer",
	}

	err = a.db.AsTx(ctx,
		func(s postgredb.Storage) error {
			if err := a.db.UpdateAccountBalance(accountSender); err != nil {
				return err
			}

			if err := a.db.UpdateAccountBalance(accountRecipient); err != nil {
				return err
			}

			if err := a.db.SaveInternalTransaction(transaction); err != nil {
				return err
			}
			return nil
		},
	)

	return fmt.Errorf(`internaltransaction: %w`, err)
}
