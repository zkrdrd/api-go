package postgredb

import (
	"api-go/pkg/models"
	"log"
)

type filter struct {
	order_by_collumn  string
	order_by_asc_desc string
	order_by_limit    string
	order_by_offset   string
}

const (
	order_by_asc   = "ASC"
	order_by_desc  = "DESC"
	order_by_limit = "(SELECT COUNT(id) FROM transactions)"
)

// Определение фильтра для запроса
// column - Колонка по которой будет производиться фильтрация default "created_at";
// ask_desc - Фильтрация "ASC" от меньшега к большему, "DESC" от большега к меньшему, default "ASC";
// limit - "число" сколько элементов брать default "ALL";
// offset - "число" сколько элементов пропустить default "0";
func Filter(column, ask_desc, limit, offset string) *filter {
	if column == "" {
		column = "created_at"
	}
	if ask_desc == "" || (ask_desc != order_by_asc && ask_desc != order_by_desc) {
		ask_desc = order_by_asc
	}
	if limit == "" || limit == "ALL" {
		limit = order_by_limit
	}
	if offset == "" {
		offset = "0"
	}
	return &filter{
		order_by_collumn:  column,
		order_by_asc_desc: ask_desc,
		order_by_limit:    limit,
		order_by_offset:   offset,
	}
}

// Получение транзакции по id
func (db *DB) GetInternalTrasaction(id string) (*models.InternalTransaction, error) {
	transf := &models.InternalTransaction{}
	if err := db.conn.QueryRow(`
	SELECT account_sender, account_recipient, amount, created_at 
	FROM transactions WHERE id = $1;`, id).Scan(
		&transf.AccountRecipient,
		&transf.AccountSender,
		&transf.Amount,
		&transf.CreatedAt); err != nil {
		log.Print(err)
		return nil, err
	}
	return transf, nil
}

// Получение всех транзакций из БД в slice
func (db *DB) ListInternalTransaction(filt *filter) ([]*models.InternalTransaction, error) {
	// TODO:
	// 1. не принимает параметр ALL для limit
	transfSlice := []*models.InternalTransaction{}
	rows, err := db.conn.Query(`
	SELECT account_sender, account_recipient, amount, created_at 
	FROM transactions ORDER BY $1, $2 LIMIT $3 OFFSET $4;`,
		filt.order_by_collumn,
		filt.order_by_asc_desc,
		filt.order_by_limit,
		filt.order_by_offset,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		transf := &models.InternalTransaction{}
		if err := rows.Scan(&transf.AccountSender, &transf.AccountRecipient, &transf.Amount, &transf.CreatedAt); err != nil {
			log.Fatal(err)
		}
		transfSlice = append(transfSlice, transf)
	}
	return transfSlice, nil
}

// Запись транзакций в БД
func (db *DB) SaveInternalTransaction(transf *models.InternalTransaction) error {
	// todo
	// 1. изменение баланса
	if _, err := db.conn.Exec(`
	INSERT INTO transactions (account_sender, account_recipient, amount, created_at) 
	VALUES (
	$1, --AccountSender
    $2, --AccountRecipient
    $3, --Amount
	$4); --CreatedAt`,
		transf.AccountSender,
		transf.AccountRecipient,
		transf.Amount,
		transf.CreatedAt); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
