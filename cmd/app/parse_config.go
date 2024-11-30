package app

import (
	"errors"
	"log"

	"github.com/zkrdrd/api-go/internal/postgredb"

	"github.com/zkrdrd/ConfigParser"
)

func ParseDBConfig(confPath string) (*postgredb.DBConfig, error) {
	var cfg = &postgredb.DBConfig{}
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
