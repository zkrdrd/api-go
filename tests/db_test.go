package tests

import (
	"api-go/internal/postgredb"
	"api-go/pkg/models"
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/zkrdrd/ConfigParser"
)

func TestDB(t *testing.T) {

	dbconf, _ := parseDBConfig("ConConf.json")
	db, _ := dbconf.NewDB()
	db.DeleteAllRowsInTableTransactions()

	for _, message := range TestValue {
		db.SaveInternalTransaction(message.Msg)
	}

	resData, err := db.ListInternalTransaction()
	if err != nil {
		log.Fatal("error")
	}

	for key, value := range resData {
		if TestValue[key].Msg.AccountSender != value.AccountSender {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValue[key].Msg.AccountSender, value.AccountSender))
		}
		if TestValue[key].Msg.AccountRecipient != value.AccountRecipient {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValue[key].Msg.AccountRecipient, value.AccountRecipient))
		}
		if TestValue[key].Msg.Amount != value.Amount {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValue[key].Msg.Amount, value.Amount))
		}
	}
}

func parseDBConfig(confPath string) (*postgredb.DBConfig, error) {
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

var TestValue = []struct {
	Msg *models.InternalTransaction
}{
	{
		Msg: &models.InternalTransaction{
			AccountSender:    `1`,
			AccountRecipient: `2`,
			Amount:           `50`,
		},
	},
	{
		Msg: &models.InternalTransaction{
			AccountSender:    `3`,
			AccountRecipient: `4`,
			Amount:           `500`,
		},
	},
	{
		Msg: &models.InternalTransaction{
			AccountSender:    `5`,
			AccountRecipient: `6`,
			Amount:           `5000`,
		},
	},
	{
		Msg: &models.InternalTransaction{
			AccountSender:    `7`,
			AccountRecipient: `8`,
			Amount:           `50000`,
		},
	},
	{
		Msg: &models.InternalTransaction{
			AccountSender:    `9`,
			AccountRecipient: `10`,
			Amount:           `500000`,
		},
	},
}
