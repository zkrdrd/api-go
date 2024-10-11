package tests

import (
	"api-go/cmd/app"
	"api-go/internal/business"
	"api-go/pkg/models"
	"context"
	"fmt"
	"testing"
)

func TestBusines(t *testing.T) {

	dbconf, _ := app.ParseDBConfig("D:\\Programming\\api-go\\ConConf.json")
	db, _ := dbconf.NewDB()

	db.RecreateTableCustomers()
	db.RecreateTableAccountBalance()
	db.RecreateTableInternalTransactions()

	for _, user := range ValueCustomers {
		if err := db.SaveUser(user.MsgCustomers); err != nil {
			t.Error(fmt.Errorf(`error "SaveUser" %v`, err))
		}
	}

	for _, balance := range ValueAccountBalances {
		db.SaveAccountBalance(balance.MsgAccountBalance)
	}

	a := &business.Accouting{
		DB: db,
	}

	for _, cashIn := range ValueCasheIn {
		a.CashIn(context.Background(), cashIn.MsgValueCashIn)
	}

	for _, transaction := range ValueInternalTransactions {
		a.InternalTransfer(context.Background(), transaction.MsgInternalTransaction)
	}
}

var (
	ValueCustomers = []struct {
		MsgCustomers *models.Users
	}{
		{
			MsgCustomers: &models.Users{
				FirstName:  "A",
				LastName:   "A",
				MiddleName: "A",
			},
		},
		{
			MsgCustomers: &models.Users{
				FirstName:  "B",
				LastName:   "B",
				MiddleName: "B",
			},
		},
		{
			MsgCustomers: &models.Users{
				FirstName:  "C",
				LastName:   "C",
				MiddleName: "C",
			},
		},
		{
			MsgCustomers: &models.Users{
				FirstName:  "D",
				LastName:   "D",
				MiddleName: "D",
			},
		},
		{
			MsgCustomers: &models.Users{
				FirstName:  "E",
				LastName:   "E",
				MiddleName: "E",
			},
		},
	}

	ValueAccountBalances = []struct {
		MsgAccountBalance *models.Balance
	}{
		{
			MsgAccountBalance: &models.Balance{
				Account:   "1",
				Amount:    "0",
				CreatedAt: "2024-10-04T15:34:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "2",
				Amount:    "0",
				CreatedAt: "2024-10-04T15:35:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "3",
				Amount:    "0",
				CreatedAt: "2024-10-04T15:36:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "4",
				Amount:    "0",
				CreatedAt: "2024-10-04T15:37:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "5",
				Amount:    "0",
				CreatedAt: "2024-10-04T15:38:43+05:00",
				UpdatedAt: "",
			},
		},
	}

	ValueCasheIn = []struct {
		MsgValueCashIn *models.CashIn
	}{
		{
			MsgValueCashIn: &models.CashIn{
				Account: `1`,
				Amount:  `10`,
			},
		},
		{
			MsgValueCashIn: &models.CashIn{
				Account: `2`,
				Amount:  `1000000`,
			},
		},
		{
			MsgValueCashIn: &models.CashIn{
				Account: `3`,
				Amount:  `100000`,
			},
		},
		{
			MsgValueCashIn: &models.CashIn{
				Account: `4`,
				Amount:  `10000`,
			},
		},
		{
			MsgValueCashIn: &models.CashIn{
				Account: `5`,
				Amount:  `1000`,
			},
		},
	}

	ValueInternalTransactions = []struct {
		MsgInternalTransaction *models.InternalTranser
	}{
		{
			MsgInternalTransaction: &models.InternalTranser{
				AccountSender:    `1`,
				AccountRecipient: `2`,
				Amount:           `500000`,
				CreatedAt:        `2024-10-06T15:34:43+05:00`,
			},
		},
		{
			MsgInternalTransaction: &models.InternalTranser{
				AccountSender:    `2`,
				AccountRecipient: `3`,
				Amount:           `50000`,
				CreatedAt:        `2024-10-06T15:35:43+05:00`,
			},
		},
		{
			MsgInternalTransaction: &models.InternalTranser{
				AccountSender:    `3`,
				AccountRecipient: `4`,
				Amount:           `5000`,
				CreatedAt:        `2024-10-06T15:36:43+05:00`,
			},
		},
		{
			MsgInternalTransaction: &models.InternalTranser{
				AccountSender:    `4`,
				AccountRecipient: `5`,
				Amount:           `500`,
				CreatedAt:        `2024-10-06T15:37:43+05:00`,
			},
		},
		{
			MsgInternalTransaction: &models.InternalTranser{
				AccountSender:    `5`,
				AccountRecipient: `6`,
				Amount:           `50`,
				CreatedAt:        `2024-10-06T15:38:43+05:00`,
			},
		},
	}
)
