package postgredb

// Удаление Всех данных из табилцы iternal_transaction
func (db *DB) DeleteAllRowsInTableTransactions() {
	db.conn.Query(`TRUNCATE TABLE transactions;`)
}

// Удаление базы данных
func (db *DB) DeleteDatabase() {
	db.conn.Query(`DROP DATABASE api;`)
}
