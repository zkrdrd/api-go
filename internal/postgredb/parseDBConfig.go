package postgredb

import (
	"errors"
	"log"

	"github.com/zkrdrd/ConfigParser"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
	SSLmode  string `json:"sslmode,omitempty"`
}

func Parse(confPath string) (*DBConfig, error) {
	var cfg = &DBConfig{}
	if err := ConfigParser.Read(confPath, cfg); err != nil {
		log.Fatal(err)
		return nil, err
	}
	if cfg.Host == "" || (cfg.Port <= 0 && cfg.Port >= 65536) || cfg.User == "" || cfg.Password == "" || cfg.DBname == "" {
		err := errors.New("config error: config is not filled")
		log.Fatal(err)
		return nil, err
	}
	return cfg, nil
}
