package postgredb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/zkrdrd/ConfigParser"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     int    `json:"user"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
	SSLmode  string `json:"sslmode,omitempty"`
}

type DB struct {
	Conn *sql.DB
}

func Parse() error {
	var cfg = &DBConfig{}
	if err := ConfigParser.Read("ConConf.json", cfg); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

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
	// todo:
	// 1. подключение передается в запросы
	return &DB{Conn: db}, nil
}
