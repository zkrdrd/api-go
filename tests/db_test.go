package tests

import (
	"api-go/cmd/app"
	postgredb "api-go/internal/postgre"
	"api-go/pkg/models"
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	dbconf, _ := app.ParseDBConfig("ConConf.json")
	db, _ := dbconf.NewDB()

	db.RecreateTableCustomers()
	db.RecreateTableAccountBalance()
	db.RecreateTableInternalTransactions()

	customer(t, db)

	accountBalances(t, db)

	internalTransactions(t, db)

}

// test db customers
func customer(t *testing.T, db *postgredb.DB) {

	for _, message := range TestCustomer {
		if err := db.SaveUser(message.MsgCustomers); err != nil {
			t.Error(fmt.Errorf(`error "SaveUser" %v`, err))
		}
	}

	res, err := db.GetUser("1")
	if err != nil {
		t.Error(fmt.Errorf(`error "GetUser" %v`, err))
	}
	if TestCustomer[0].MsgCustomers.FirstName != res.FirstName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestCustomer[0].MsgCustomers.FirstName, res.FirstName))
	}
	if TestCustomer[0].MsgCustomers.LastName != res.LastName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestCustomer[0].MsgCustomers.LastName, res.LastName))
	}
	if TestCustomer[0].MsgCustomers.MiddleName != res.MiddleName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestCustomer[0].MsgCustomers.MiddleName, res.MiddleName))
	}
	if TestCustomer[0].MsgCustomers.MiddleName != res.MiddleName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestCustomer[0].MsgCustomers.MiddleName, res.MiddleName))
	}
}

// test db account_balance
func accountBalances(t *testing.T, db *postgredb.DB) {

	for _, message := range TestAccountBalance {
		if err := db.SaveAccountBalance(message.MsgAccountBalance); err != nil {
			t.Error(fmt.Errorf(`error "SaveAccountBalance" %v`, err))
		}
	}

	resData, err := db.GetAccountBalance("1")
	if err != nil {
		t.Error(fmt.Errorf(`error "GetAccountBalance" %v`, err))
	}

	if TestAccountBalance[0].MsgAccountBalance.Account != resData.Account {
		t.Error(fmt.Errorf(`result field %v != %v`, TestAccountBalance[0].MsgAccountBalance.Account, resData.Account))
	}
	if TestAccountBalance[0].MsgAccountBalance.Amount != resData.Amount {
		t.Error(fmt.Errorf(`result field %v != %v`, TestAccountBalance[0].MsgAccountBalance.Amount, resData.Amount))
	}
	if TestAccountBalance[0].MsgAccountBalance.CreatedAt != resData.CreatedAt {
		t.Error(fmt.Errorf(`result field %v != %v`, TestAccountBalance[0].MsgAccountBalance.CreatedAt, resData.CreatedAt))
	}

	err = db.UpdateAccountBalance(TestUpdateAccountBalance)
	if err != nil {
		t.Error(fmt.Errorf(`error "UpdateAccountBalance" %v`, err))
	}

	resDataUpdate, err := db.GetAccountBalance("1")
	if err != nil {
		t.Error(fmt.Errorf(`error "GetAccountBalance" %v`, err))
	}

	if TestUpdateAccountBalance.Account != resDataUpdate.Account {
		t.Error(fmt.Errorf(`result field %v != %v`, TestUpdateAccountBalance.Account, resDataUpdate.Account))
	}
	if TestUpdateAccountBalance.Amount != resDataUpdate.Amount {
		t.Error(fmt.Errorf(`result field %v != %v`, TestUpdateAccountBalance.Amount, resDataUpdate.Amount))
	}
	if TestUpdateAccountBalance.CreatedAt != resDataUpdate.CreatedAt {
		t.Error(fmt.Errorf(`result field %v != %v`, TestUpdateAccountBalance.CreatedAt, resDataUpdate.CreatedAt))
	}

}

// test db InternalTransactions
func internalTransactions(t *testing.T, db *postgredb.DB) {

	for _, message := range TestTransactions {

		if err := db.SaveInternalTransaction(message.MsgTransaction); err != nil {
			t.Error(fmt.Errorf(`error %v`, err))
		}
	}

	count, err := db.CountInternalTransactions()
	if err != nil {
		t.Error(fmt.Errorf(`error "CountInternalTransactions" %v`, err))
	}

	if countRes != count {
		t.Error(fmt.Errorf(`result field %v != %v`, countRes, count))
	}

	filterInternalTransactions := postgre.FilterInternalTransaction("amount", "DESC", count, 0)

	res, err := db.GetInternalTrasaction("1")
	if err != nil {
		t.Error(fmt.Errorf(`error "GetInternalTrasaction" %v`, err))
	}

	if TestTransactions[0].MsgTransaction.AccountSender != res.AccountSender {
		t.Error(fmt.Errorf(`result field %v != %v`, TestTransactions[0].MsgTransaction.AccountSender, res.AccountSender))
	}
	if TestTransactions[0].MsgTransaction.AccountRecipient != res.AccountRecipient {
		t.Error(fmt.Errorf(`result field %v != %v`, TestTransactions[0].MsgTransaction.AccountRecipient, res.AccountRecipient))
	}
	if TestTransactions[0].MsgTransaction.Amount != res.Amount {
		t.Error(fmt.Errorf(`result field %v != %v`, TestTransactions[0].MsgTransaction.Amount, res.Amount))
	}
	if TestTransactions[0].MsgTransaction.CreatedAt != res.CreatedAt {
		t.Error(fmt.Errorf(`result field %v != %v`, TestTransactions[0].MsgTransaction.CreatedAt, res.CreatedAt))
	}
	if TestTransactions[0].MsgTransaction.TransactionType != res.TransactionType {
		t.Error(fmt.Errorf(`result field %v != %v`, TestTransactions[0].MsgTransaction.TransactionType, res.TransactionType))
	}

	resData, err := db.ListInternalTransaction(filterInternalTransactions)
	if err != nil {
		t.Error(fmt.Errorf(`error "ListInternalTransaction" %v`, err))
	}

	for key, value := range resData {
		if TestTransactions[key].MsgTransaction.AccountSender != value.AccountSender {
			t.Error(fmt.Errorf(`result field %v != %v`, TestTransactions[key].MsgTransaction.AccountSender, value.AccountSender))
		}
		if TestTransactions[key].MsgTransaction.AccountRecipient != value.AccountRecipient {
			t.Error(fmt.Errorf(`result field %v != %v`, TestTransactions[key].MsgTransaction.AccountRecipient, value.AccountRecipient))
		}
		if TestTransactions[key].MsgTransaction.Amount != value.Amount {
			t.Error(fmt.Errorf(`result field %v != %v`, TestTransactions[key].MsgTransaction.Amount, value.Amount))
		}
		if TestTransactions[key].MsgTransaction.CreatedAt != value.CreatedAt {
			t.Error(fmt.Errorf(`result field %v != %v`, TestTransactions[key].MsgTransaction.CreatedAt, value.CreatedAt))
		}
		if TestTransactions[key].MsgTransaction.TransactionType != value.TransactionType {
			t.Error(fmt.Errorf(`result field %v != %v`, TestTransactions[key].MsgTransaction.TransactionType, value.TransactionType))
		}
	}
}

var (
	TestCustomer = []struct {
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

	TestAccountBalance = []struct {
		MsgAccountBalance *models.Balance
	}{
		{
			MsgAccountBalance: &models.Balance{
				Account:   "1",
				Amount:    "90000000",
				CreatedAt: "2024-10-04T15:34:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "2",
				Amount:    "9000000",
				CreatedAt: "2024-10-04T15:35:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "3",
				Amount:    "900000",
				CreatedAt: "2024-10-04T15:36:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "4",
				Amount:    "90000",
				CreatedAt: "2024-10-04T15:37:43+05:00",
				UpdatedAt: "",
			},
		},
		{
			MsgAccountBalance: &models.Balance{
				Account:   "5",
				Amount:    "9000",
				CreatedAt: "2024-10-04T15:38:43+05:00",
				UpdatedAt: "",
			},
		},
	}

	TestValueCasheIn = []struct {
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

	TestTransactions = []struct {
		MsgTransaction *models.Transactions
	}{
		{
			MsgTransaction: &models.Transactions{
				AccountSender:    `1`,
				AccountRecipient: `2`,
				Amount:           `500000`,
				CreatedAt:        `2024-10-06T15:34:43+05:00`,
				TransactionType:  "Transfer",
			},
		},
		{
			MsgTransaction: &models.Transactions{
				AccountSender:    `2`,
				AccountRecipient: `3`,
				Amount:           `50000`,
				CreatedAt:        `2024-10-06T15:35:43+05:00`,
				TransactionType:  "Transfer",
			},
		},
		{
			MsgTransaction: &models.Transactions{
				AccountSender:    `3`,
				AccountRecipient: `4`,
				Amount:           `5000`,
				CreatedAt:        `2024-10-06T15:36:43+05:00`,
				TransactionType:  "Transfer",
			},
		},
		{
			MsgTransaction: &models.Transactions{
				AccountSender:    `4`,
				AccountRecipient: `5`,
				Amount:           `500`,
				CreatedAt:        `2024-10-06T15:37:43+05:00`,
				TransactionType:  "Transfer",
			},
		},
		{
			MsgTransaction: &models.Transactions{
				AccountSender:    `5`,
				AccountRecipient: `1`,
				Amount:           `50`,
				CreatedAt:        `2024-10-06T15:38:43+05:00`,
				TransactionType:  "Transfer",
			},
		},
	}
	countRes                 = 5
	TestUpdateAccountBalance = &models.Balance{
		Account:   "1",
		Amount:    "1",
		CreatedAt: "2024-10-04T15:34:43+05:00",
		UpdatedAt: "2024-10-05T15:34:43+05:00",
	}
)
