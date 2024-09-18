package business

import "context"

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

// Я тип - я описываю перевод с одного счета на другой
type InternalTransfer struct {
	AccountSender    string
	AccountRecipient string
	Amount           string
}

// Я знаю как делать операции с счетами пользователя
type Accouting struct {
	// Во мне лежит все  необходимое для работы
	// к примеру подключение к БД, а возможно и подключения
	// к другим сервисам
	//db *db.Conn
}

// Тут я пополняю счет наличными
func (a *Accouting) CacheOut(ctx context.Context, checkout *CacheOut) error {
	// Блокирую баланс
	// Обналичиваю средства
	// Изменяю сумму баланса
	// Разблокирую баланс
	return nil
}

// Тут я снимаю со счета начличные
func (a *Accouting) CacheIn(ctx context.Context, checkin *CacheIn) error {
	//...
	return nil
}

// Тут я перевожу деньги между внетренними счетами
func (a *Accouting) InternalTransfer(ctx context.Context, transfer *InternalTransfer) error {
	//...
	return nil
}
