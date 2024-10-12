package main

import (
	"api-go/cmd/app"
	"api-go/internal/postgre"
)

func main() {

	dbconf, _ := app.ParseDBConfig("ConConf.json")
	db, _ := dbconf.NewDB()
	//db.SaveUser(usr)

	// res, _ := db.GetUser("1")
	// fmt.Println(*res)

	//db.SaveInternalTransaction(transf)

	// res, _ := db.GetTransfer("1")
	// fmt.Println(*res)
	filter := postgre.FilterInternalTransaction("", "", 0, 0)

	_, _ = db.ListInternalTransaction(filter)
}

// var usr = &models.Users{
// 	FirstName:  "FirstName",
// 	LastName:   "LastName",
// 	MiddleName: "MiddleName",
// }

// var transf = &models.InternalTransaction{
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
