//go:build integration

package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/zkrdrd/api-go/cmd/app"
	"github.com/zkrdrd/api-go/internal/locker"
	"github.com/zkrdrd/api-go/internal/postgredb"
	"github.com/zkrdrd/api-go/internal/services"
	"github.com/zkrdrd/api-go/pkg/models"
)

func TestBusines(t *testing.T) {

	dbconf, _ := app.ParseDBConfig("D:\\Programming\\api-go\\ConConf.json")
	db, _ := dbconf.NewDB()

	if err := db.RecreateTableCustomers(); err != nil {
		t.Error(err)
		return
	}
	if err := db.RecreateTableAccountBalance(); err != nil {
		t.Error(err)
		return
	}
	if err := db.RecreateTableInternalTransactions(); err != nil {
		t.Error(err)
		return
	}

	for _, user := range ValueCustomers {
		if err := db.SaveUser(user.MsgCustomers); err != nil {
			t.Error(fmt.Errorf(`"save user %s: %w`, user.MsgCustomers.FirstName, err))
		}
	}

	for _, balance := range ValueAccountBalances {
		db.SaveAccountBalance(balance.MsgAccountBalance)
	}

	a := services.NewAccouting(db, locker.NewLocker())

	for _, cashIn := range ValueCasheIn {
		a.CashIn(context.Background(), cashIn.MsgValueCashIn)
	}

	for _, balance := range ValueCasheIn {

		accountBalance, _ := db.GetAccountBalance(balance.MsgValueCashIn.Account)

		if balance.MsgValueCashIn.Account != accountBalance.Account {
			t.Error(fmt.Errorf(`result field %s != %s`, balance.MsgValueCashIn.Account, accountBalance.Account))
		}
		if balance.MsgValueCashIn.Amount != accountBalance.Amount {
			t.Error(fmt.Errorf(`result field %s != %s`, balance.MsgValueCashIn.Amount, accountBalance.Amount))
		}
	}

	for _, transaction := range ValueInternalTransactions {
		a.InternalTransfer(context.Background(), transaction.MsgInternalTransaction)
	}

	count, err := db.CountInternalTransactions()
	if err != nil {
		t.Error(fmt.Errorf(`error "CountInternalTransactions" %w`, err))
	}

	if countResult != count {
		t.Error(fmt.Errorf(`result field %v != %v`, countResult, count))
	}

	filterInternalTransactions := postgredb.FilterInternalTransaction("", "", count, 0)

	resData, err := db.ListInternalTransaction(filterInternalTransactions)
	if err != nil {
		t.Error(fmt.Errorf(`error "ListInternalTransaction" %w`, err))
	}

	for index, value := range resData {
		if ExpectedValueTransactions[index].MsgTransactions.AccountSender != value.AccountSender {
			t.Error(fmt.Errorf(`result field %v != %v`, ExpectedValueTransactions[index].MsgTransactions.AccountSender, value.AccountSender))
		}
		if ExpectedValueTransactions[index].MsgTransactions.AccountRecipient != value.AccountRecipient {
			t.Error(fmt.Errorf(`result field %v != %v`, ExpectedValueTransactions[index].MsgTransactions.AccountRecipient, value.AccountRecipient))
		}
		if ExpectedValueTransactions[index].MsgTransactions.Amount != value.Amount {
			t.Error(fmt.Errorf(`result field %v != %v`, ExpectedValueTransactions[index].MsgTransactions.Amount, value.Amount))
		}
		if ExpectedValueTransactions[index].MsgTransactions.TransactionType != value.TransactionType {
			t.Error(fmt.Errorf(`result field %v != %v`, ExpectedValueTransactions[index].MsgTransactions.TransactionType, value.TransactionType))
		}
	}

	for _, balance := range ExpectedValueAccountBalance {

		accountBalance, _ := db.GetAccountBalance(balance.MsgAccountBalance.Account)

		if balance.MsgAccountBalance.Account != accountBalance.Account {
			t.Error(fmt.Errorf(`result field %v != %v`, balance.MsgAccountBalance.Account, accountBalance.Account))
		}
		if balance.MsgAccountBalance.Amount != accountBalance.Amount {
			t.Error(fmt.Errorf(`result field %v != %v`, balance.MsgAccountBalance.Amount, accountBalance.Amount))
		}
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
				Amount:  `10000000`,
			},
		},
		{
			MsgValueCashIn: &models.CashIn{
				Account: `2`,
				Amount:  `10000000`,
			},
		},
		{
			MsgValueCashIn: &models.CashIn{
				Account: `3`,
				Amount:  `10000000`,
			},
		},
		{
			MsgValueCashIn: &models.CashIn{
				Account: `4`,
				Amount:  `10000000`,
			},
		},
		{
			MsgValueCashIn: &models.CashIn{
				Account: `5`,
				Amount:  `10000000`,
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
				Amount:           `5000000`,
				CreatedAt:        `2024-10-06T15:34:43+05:00`,
			},
		},
		{
			MsgInternalTransaction: &models.InternalTranser{
				AccountSender:    `2`,
				AccountRecipient: `3`,
				Amount:           `5000000`,
				CreatedAt:        `2024-10-06T15:35:43+05:00`,
			},
		},
		{
			MsgInternalTransaction: &models.InternalTranser{
				AccountSender:    `3`,
				AccountRecipient: `4`,
				Amount:           `5000000`,
				CreatedAt:        `2024-10-06T15:36:43+05:00`,
			},
		},
		{
			MsgInternalTransaction: &models.InternalTranser{
				AccountSender:    `4`,
				AccountRecipient: `5`,
				Amount:           `5000000`,
				CreatedAt:        `2024-10-06T15:37:43+05:00`,
			},
		},
		{
			MsgInternalTransaction: &models.InternalTranser{
				AccountSender:    `5`,
				AccountRecipient: `6`,
				Amount:           `5000000`,
				CreatedAt:        `2024-10-06T15:38:43+05:00`,
			},
		},
	}

	ExpectedValueTransactions = []struct {
		MsgTransactions *models.Transactions
	}{
		{
			MsgTransactions: &models.Transactions{
				AccountSender:    "1",
				AccountRecipient: "1",
				Amount:           "10000000",
				CreatedAt:        "",
				TransactionType:  "Cash in",
			},
		},
		{
			MsgTransactions: &models.Transactions{
				AccountSender:    "2",
				AccountRecipient: "2",
				Amount:           "10000000",
				CreatedAt:        "",
				TransactionType:  "Cash in",
			},
		},
		{
			MsgTransactions: &models.Transactions{
				AccountSender:    "3",
				AccountRecipient: "3",
				Amount:           "10000000",
				CreatedAt:        "",
				TransactionType:  "Cash in",
			},
		},
		{
			MsgTransactions: &models.Transactions{
				AccountSender:    "4",
				AccountRecipient: "4",
				Amount:           "10000000",
				CreatedAt:        "",
				TransactionType:  "Cash in",
			},
		},
		{
			MsgTransactions: &models.Transactions{
				AccountSender:    "5",
				AccountRecipient: "5",
				Amount:           "10000000",
				CreatedAt:        "",
				TransactionType:  "Cash in",
			},
		},
		{
			MsgTransactions: &models.Transactions{
				AccountSender:    "1",
				AccountRecipient: "2",
				Amount:           "5000000",
				CreatedAt:        "",
				TransactionType:  "Transfer",
			},
		},
		{
			MsgTransactions: &models.Transactions{
				AccountSender:    "2",
				AccountRecipient: "3",
				Amount:           "5000000",
				CreatedAt:        "",
				TransactionType:  "Transfer",
			},
		},
		{
			MsgTransactions: &models.Transactions{
				AccountSender:    "3",
				AccountRecipient: "4",
				Amount:           "5000000",
				CreatedAt:        "",
				TransactionType:  "Transfer",
			},
		},
		{
			MsgTransactions: &models.Transactions{
				AccountSender:    "4",
				AccountRecipient: "5",
				Amount:           "5000000",
				CreatedAt:        "",
				TransactionType:  "Transfer",
			},
		},
	}

	ExpectedValueAccountBalance = []struct {
		MsgAccountBalance *models.Balance
	}{
		{
			MsgAccountBalance: &models.Balance{
				Account:   "1",
				Amount:    "5000000",
				CreatedAt: "2024-10-04T15:34:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "2",
				Amount:    "10000000",
				CreatedAt: "2024-10-04T15:35:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "3",
				Amount:    "10000000",
				CreatedAt: "2024-10-04T15:36:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "4",
				Amount:    "10000000",
				CreatedAt: "2024-10-04T15:37:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "5",
				Amount:    "15000000",
				CreatedAt: "2024-10-04T15:38:43+05:00",
				UpdatedAt: "",
			},
		},
	}

	countResult = 9
)
