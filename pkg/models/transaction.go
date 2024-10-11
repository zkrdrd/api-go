package models

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
