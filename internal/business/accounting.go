package business

import (
	"api-go/internal/postgre"
	"api-go/pkg/models"
	"context"
)

// Я знаю как делать операции с счетами пользователя
type Accouting struct {
	db *postgre.DB
	// Во мне лежит все  необходимое для работы
	// к примеру подключение к БД, а возможно и подключения
	// к другим сервисам
	//db *db.Conn
}

// Тут я пополняю счет наличными
func (a *Accouting) CacheOut(ctx context.Context, checkout *models.CacheOut) error {
	// TODO:
	// 1. Блокирую баланс
	// 2. Разблокирую баланс

	// Обналичиваю средства
	// Изменяю сумму баланса
	return nil
}

// Тут я снимаю со счета начличные
func (a *Accouting) CacheIn(ctx context.Context, checkin *models.CacheIn) error {
	//...
	return nil
}

// Тут я перевожу деньги между внетренними счетами
func (a *Accouting) InternalTransfer(ctx context.Context, transfer *models.InternalTransaction) error {
	a.db.SaveInternalTransaction(transfer)
	return nil
}
