package postgredb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

// Инициализация соединения с БД
func (dbconf *DBConfig) NewDB() (*DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf(`host=%v 
	port=%v 
	user=%v 
	password=%v 
	dbname=%v 
	sslmode=%v`,
			dbconf.Host,
			dbconf.Port,
			dbconf.User,
			dbconf.Password,
			dbconf.DBname,
			dbconf.SSLmode))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &DB{conn: db}, nil
}
