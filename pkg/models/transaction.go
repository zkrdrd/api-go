package models

import (
	"fmt"
	"math/big"
)

// я тип - я описываю пополнение наличными
type CashIn struct {
	Account string
	Amount  string
}

// я тип - я описываю снятие наличными
type CashOut struct {
	Account string
	Amount  string
}

type Balance struct {
	Account   string
	Amount    string
	CreatedAt string
	UpdatedAt string
}

func (b *Balance) SetAmount(amount *big.Int) error {
	if amount == nil {
		return fmt.Errorf(`amount is nil`)
	}
	b.Amount = amount.String()
	return nil
}

func (b *Balance) GetAmount() *big.Int {
	amount := big.NewInt(0)
	_, _ = amount.SetString(b.Amount, 10)
	return amount
}

// Я тип - я описываю перевод с одного счета на другой
type InternalTranser struct {
	AccountSender    string
	AccountRecipient string
	Amount           string
	CreatedAt        string
}

// Я тип - я описываю перевод с одного счета на другой
type Transactions struct {
	AccountSender    string
	AccountRecipient string
	Amount           string
	CreatedAt        string
	TransactionType  string
}
