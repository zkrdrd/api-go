package main

import (
	"api-go/internal/business"
	"api-go/internal/postgredb"
)

var usr = &business.Users{
	FirstName:  "FirstName",
	LastName:   "LastName",
	MiddleName: "MiddleName",
}

func main() {
	//var use = users.NewUser()
	dbconf, _ := postgredb.Parse("ConConf.json")
	db, _ := dbconf.NewDB()
	db.SaveUser(usr)

	// ----------------------------------------------------------
	// addressServer := "127.0.0.1:8080"

	// ctx := context.Background()
	// accoutingService := users.NewUsers()

	// server := server.NewServer(addressServer)
	// server.AddHandler(accoutingService.Handlers())
	// server.Run(ctx)

	// time.Sleep(time.Second * 1)

	// // Prepare message
	// dataForCheck := `{"id": "1",
	// "UserName": "asdf",
	// "Password": "asdfasdf"}`
	// buf := &bytes.Buffer{}
	// buf.WriteString(dataForCheck)
	// -----------------------------------------------------------
	// request builder
	//	_, _ := http.NewRequest(http.MethodPost, `http://`+addressServer+`/users`, buf)

	// ctx := context.Background()
	// srv := server.NewServer("127.0.0.1:8080")
	// user := users.NewUsers()
	// mux := user.Handlers()
	// srv.AddHandler(mux)
	// srv.Run(ctx)
	//var customer *customers.Customer]
	// mux := http.NewServeMux()
	// mux.HandleFunc("/customers", cust.CreateCustomer)
	// mux.HandleFunc("/transaction/refill", ref.CreateRefill)
	// mux.HandleFunc("/transaction/remittance", rem.CreateRemittance)
	// http.ListenAndServe(":8001", mux)
}

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
