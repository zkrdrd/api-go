package main

import (
	"api-go/internal/app/customers"
	"api-go/internal/app/transaction/refill"
	"api-go/internal/app/transaction/remittance"
	"net/http"
)

func main() {
	//var customer *customers.Customer]
	http.HandleFunc("/customers", cust.CreateCustomer)
	http.HandleFunc("/transaction/refill", ref.CreateRefill)
	http.HandleFunc("/transaction/remittance", rem.CreateRemittance)
	http.ListenAndServe(":8001", nil)
}

var cust = &customers.Customer{
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
