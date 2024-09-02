package main

import (
	"api-go/internal/app/users"
	"api-go/pkg/server"
)

func main() {
	srv := server.NewServer("127.0.0.1:8080")
	user := users.NewUsers()
	mux := user.Handlers()
	srv.AddHandler(mux)
	srv.Run()
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
