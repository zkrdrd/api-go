package models

import (
	"errors"
	"math/big"
)

// CashIn - описываю пополнение наличными.
type CashIn struct {
	Account string `json:"account"`
	Amount  string `json:"amount"`
}

// CashOut - описываю снятие наличными.
type CashOut struct {
	Account string `json:"account"`
	Amount  string `json:"amount"`
}

// 6 символов точности.
type Balance struct {
	Account   string `json:"account"`
	Amount    string `json:"amount"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updateAt"`
}

func (b *Balance) GetBalance() *big.Int {
	amount := big.NewInt(0)
	_, _ = amount.SetString(b.Amount, 10)
	return amount
}

func (b *Balance) SetBalance(amount *big.Int) error {
	if amount == nil {
		return errors.New(`amount is nil`)
	}
	b.Amount = amount.String()
	return nil
}

// Я тип - я описываю перевод с одного счета на другой.
type InternalTranser struct {
	AccountSender    string `json:"accountSender"`
	AccountRecipient string `json:"accountRecipient"`
	Amount           string `json:"amount"`
	CreatedAt        string `json:"createdAt"`
}

// Я тип - я описываю перевод с одного счета на другой.
type Transactions struct {
	AccountSender    string `json:"accountSender"`
	AccountRecipient string `json:"accountRecipient"`
	Amount           string `json:"amount"`
	CreatedAt        string `json:"createdAt"`
	TransactionType  string `json:"transactionType"`
}
