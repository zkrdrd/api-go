package models

// я тип - я описываю пополнение наличными
type CacheIn struct {
	Account string
	Amount  string
}

// я тип - я описываю снятие наличными
type CacheOut struct {
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
type InternalTransaction struct {
	AccountSender    string
	AccountRecipient string
	Amount           string
	CreatedAt        string
}
