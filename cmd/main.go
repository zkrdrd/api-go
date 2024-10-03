package main

import (
	"api-go/internal/postgredb"
	"api-go/pkg/models"
	"errors"
	"log"

	"github.com/zkrdrd/ConfigParser"
)

func main() {

	dbconf, _ := parseDBConfig("ConConf.json")
	db, _ := dbconf.NewDB()
	db.SaveUser(usr)

	// res, _ := db.GetUser("1")
	// fmt.Println(*res)

	// db.SaveTransfer(transf)

	// res, _ := db.GetTransfer("1")
	// fmt.Println(*res)

	//_ = db.ListTransfer()
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

var usr = &models.Users{
	FirstName:  "FirstName",
	LastName:   "LastName",
	MiddleName: "MiddleName",
}

// var transf = &business.InternalTransfer{
// 	AccountSender:    "2",
// 	AccountRecipient: "1",
// 	Amount:           "500",
// }

/*
var cust = &handlers.Customer{
	Customer_ID: "",
	User_Name:   "Tertyfun",
	Password:    "wertghjm,",
}

var ref = &refill.Refill{
	Refill_ID:   "",
	Customer_ID: "14363456234",
	ATM:         "5112536",
	Amount:      150,
}

var rem = &remittance.Remittance{
	Remittance_ID: "",
	Customer_From: "3456789045656",
	Customer_To:   "345678903421234",
	Amount:        200,
}

/*var cust = []struct {
	Msg *customers.Customer
}{
	{
		Msg: &customers.Customer{
			Customer_ID: "",
			User_Name:   "Tertyfun",
			Password:    "wertghjm,",
		},
	},
}*/
