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

	dbconf, _ := parseDBConfig("D:\\Programming\\api-go\\ConConf.json")
	db, _ := dbconf.NewDB()
	//db.SaveUser(usr)

	// res, _ := db.GetUser("1")
	// fmt.Println(*res)

	// db.SaveTransfer(transf)

	// res, _ := db.GetTransfer("1")
	// fmt.Println(*res)

	resData, err := db.ListInternalTransaction()
	if err != nil {
		log.Fatal("error")
	}

	// Prepare message
	dataForCheck := models.InternalTransaction{
		AccountSender:    `1`,
		AccountRecipient: `2`,
		Amount:           `500`,
	}

	//buf := &bytes.Buffer{}

	for _, value := range resData {
		log.Print("print")
		if dataForCheck.AccountSender != value.AccountSender {
			t.Error(fmt.Errorf(`result field %v != %v`, dataForCheck.AccountSender, value.AccountSender))
		}
		if dataForCheck.AccountRecipient != value.AccountRecipient {
			t.Error(fmt.Errorf(`result field %v != %v`, dataForCheck.AccountRecipient, value.AccountRecipient))
		}
		if dataForCheck.Amount != value.Amount {
			t.Error(fmt.Errorf(`result field %v != %v`, dataForCheck.Amount, value.Amount))
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
