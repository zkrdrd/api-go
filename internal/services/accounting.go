package services

import (
	"api-go/internal/locker"
	postgre "api-go/internal/postgredb"
	"api-go/pkg/models"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

// Я знаю как делать операции с счетами пользователя
type Accouting struct {
	db   *postgre.DB
	lock *locker.Locker
}

var (
	ErrMoneyNotEnough   = errors.New(`error money not enough`)
	ErrAccoutingIsEmpty = errors.New(`account is empty`)
)

func NewAccouting(dbConn *postgre.DB, lock *locker.Locker) *Accouting {
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
func (a *Accouting) CashOut(ctx context.Context, cashOut *models.CashOut) error {
	if cashOut.Account == `` {
		return ErrAccoutingIsEmpty
	}

	a.lock.Lock(cashOut.Account)
	defer a.lock.Unlock(cashOut.Account)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	cashOutToBalance := &models.Balance{
		Account: cashOut.Account,
		Amount:  cashOut.Amount,
	}

	accountSender, err := a.db.GetAccountBalance(cashOut.Account)
	if err != nil {
		return err
	}

	amount := cashOutToBalance.GetBalance()
	senderBalance := accountSender.GetBalance()

	if res := amount.Cmp(senderBalance); res == +1 {
		return ErrMoneyNotEnough
	}

	// accountSender.Amount
	_ = accountSender.SetBalance(senderBalance.Sub(senderBalance, amount))
	accountSender.UpdatedAt = dateTime()

	transaction := &models.Transactions{
		AccountSender:    accountSender.Account,
		AccountRecipient: accountSender.Account,
		Amount:           accountSender.Amount,
		CreatedAt:        dateTime(),
		TransactionType:  "Cash out",
	}

	err = a.db.AsTx(ctx,
		func(tx postgre.Storage) error {
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
func (a *Accouting) CashIn(ctx context.Context, cashIn *models.CashIn) error {
	if cashIn.Account == `` {
		return ErrAccoutingIsEmpty
	}

	a.lock.Lock(cashIn.Account)
	defer a.lock.Unlock(cashIn.Account)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	cashInToBalance := &models.Balance{
		Account: cashIn.Account,
		Amount:  cashIn.Amount,
	}

	accountRecipient, err := a.db.GetAccountBalance(cashIn.Account)
	if err != nil {
		return err
	}

	amount := cashInToBalance.GetBalance()
	recipientBalance := accountRecipient.GetBalance()

	// accountSender.Amount
	_ = accountRecipient.SetBalance(recipientBalance.Add(recipientBalance, amount))
	accountRecipient.UpdatedAt = dateTime()

	transaction := &models.Transactions{
		AccountSender:    accountRecipient.Account,
		AccountRecipient: accountRecipient.Account,
		Amount:           accountRecipient.Amount,
		CreatedAt:        dateTime(),
		TransactionType:  "Cash in",
	}

	err = a.db.AsTx(ctx,
		func(tx postgre.Storage) error {
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

	accountRecipient, err := a.db.GetAccountBalance(transfer.AccountRecipient)
	if err != nil {
		return err
	}

	accountSenderToBalance := &models.Balance{
		Account: accountSender.Account,
		Amount:  accountSender.Amount,
	}

	accountRecipientToBalance := &models.Balance{
		Account: accountRecipient.Account,
		Amount:  accountRecipient.Amount,
	}

	transferToBalance := &models.Balance{
		Amount: transfer.Amount,
	}

	amount := transferToBalance.GetBalance()
	senderBalance := accountSenderToBalance.GetBalance()
	recipientBalance := accountRecipientToBalance.GetBalance()

	if res := amount.Cmp(senderBalance); res == +1 {
		return ErrMoneyNotEnough
	}

	// accountSender.Amount
	_ = accountSender.SetBalance(senderBalance.Sub(senderBalance, amount))
	accountSender.UpdatedAt = dateTime()

	_ = accountRecipient.SetBalance(recipientBalance.Add(recipientBalance, amount))
	accountRecipient.UpdatedAt = dateTime()

	transaction := &models.Transactions{
		AccountSender:    accountSender.Account,
		AccountRecipient: accountRecipient.Account,
		Amount:           transfer.Amount,
		CreatedAt:        dateTime(),
		TransactionType:  "Transfer",
	}

	err = a.db.AsTx(ctx,
		func(s postgre.Storage) error {
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
