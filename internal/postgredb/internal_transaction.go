package postgredb

import (
	"api-go/pkg/models"
	"log"
)

type filter struct {
	order_by_collumn  string
	order_by_asc_desc string
	order_by_limit    int
	order_by_offset   int
}

const (
	order_by_asc  = "ASC"
	order_by_desc = "DESC"
)

// Определение фильтра для запроса
// column - Колонка по которой будет производиться фильтрация default "created_at";
// ask_desc - Фильтрация "ASC" от меньшега к большему, "DESC" от большега к меньшему, default "ASC";
// limit - "число" сколько элементов брать;
// offset - "число" сколько элементов пропустить default 0;
func FilterInternalTransaction(column, ask_desc string, limit, offset int) *filter {
	if column == "" {
		column = "created_at"
	}
	if ask_desc == "" || (ask_desc != order_by_asc && ask_desc != order_by_desc) {
		ask_desc = order_by_asc
	}
	if offset < 0 {
		offset = 0
	}
	return &filter{
		order_by_collumn:  column,
		order_by_asc_desc: ask_desc,
		order_by_limit:    limit,
		order_by_offset:   offset,
	}
}

// Получение количества строк в таблице
func (db *DB) CountInternalTransactions() (int, error) {
	var count int
	if err := db.conn.QueryRow(`SELECT COUNT(*) FROM internal_transactions;`).Scan(&count); err != nil {
		log.Print(err)
		return 0, err
	}
	return count, nil
}

// Получение транзакции по id
func (db *DB) GetInternalTrasaction(id int) (*models.InternalTranser, error) {
	transf := &models.InternalTranser{}
	if err := db.conn.QueryRow(`
	SELECT account_sender, account_recipient, amount, created_at 
	FROM internal_transactions WHERE id = $1;`, id).Scan(
		&transf.AccountSender,
		&transf.AccountRecipient,
		&transf.Amount,
		&transf.CreatedAt); err != nil {
		return nil, err
	}
	return transf, nil
}

// Получение всех транзакций из БД в slice
func (db *DB) ListInternalTransaction(filt *filter) ([]*models.InternalTranser, error) {
	transfSlice := []*models.InternalTranser{}
	rows, err := db.conn.Query(`
	SELECT account_sender, account_recipient, amount, created_at 
	FROM internal_transactions ORDER BY $1, $2 LIMIT $3 OFFSET $4;`,
		filt.order_by_collumn,
		filt.order_by_asc_desc,
		filt.order_by_limit,
		filt.order_by_offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		transf := &models.InternalTranser{}
		if err := rows.Scan(&transf.AccountSender, &transf.AccountRecipient, &transf.Amount, &transf.CreatedAt); err != nil {
			return nil, err
		}
		transfSlice = append(transfSlice, transf)
	}
	return transfSlice, nil
}

// Запись транзакций в БД
func (db *DB) SaveInternalTransaction(transf *models.Transactions) error {
	if _, err := db.conn.Exec(`
	INSERT INTO internal_transactions (account_sender, account_recipient, amount, created_at, transaction_type) 
	VALUES (
	$1, --AccountSender
    $2, --AccountRecipient
    $3, --Amount
	$4, --CreatedAt
	$5); --TransactionType`,
		transf.AccountSender,
		transf.AccountRecipient,
		transf.Amount,
		transf.CreatedAt,
		transf.TransactionType); err != nil {
		return err
	}
	return nil
}
