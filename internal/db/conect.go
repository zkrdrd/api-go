package conect

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/zkrdrd/ConfigParser"
)

type DataConnection struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     int    `json:"user"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
	SSLmode  string `json:"sslmode,omitempty"`
}

func Parse() error {
	var cfg = &DataConnection{}
	if err := ConfigParser.Read("ConConf.json", cfg); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (dbcon *DataConnection) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf(`host=%v 
	port=%v 
	user=%v 
	password=%v 
	dbname=%v 
	sslmode=%v`,
			dbcon.Host,
			dbcon.Port,
			dbcon.User,
			dbcon.Password,
			dbcon.DBname,
			dbcon.SSLmode))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// todo:
	// 1. проверка на наличие активных сессий
	// 2. закрыть сессию если она неактивна
	return db, nil
}
