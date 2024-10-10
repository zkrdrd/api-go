package tests

import (
	"api-go/cmd/app"
	"api-go/internal/business"
	"api-go/internal/postgredb"
	"api-go/pkg/models"
	"context"
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {

	dbconf, _ := app.ParseDBConfig("D:\\Programming\\api-go\\ConConf.json")
	db, _ := dbconf.NewDB()

	db.RecreateTableCustomers()
	db.RecreateTableAccountBalance()
	db.RecreateTableInternalTransactions()

	customers(t, db)

	accountBalance(t, db)

	//internalTransactions(t, db)

}

// test db customers
func customers(t *testing.T, db *postgredb.DB) {

	for _, message := range TestValueCustomers {
		if err := db.SaveUser(message.MsgCustomers); err != nil {
			t.Error(fmt.Errorf(`error "SaveUser" %v`, err))
		}
	}

	res, err := db.GetUser("1")
	if err != nil {
		t.Error(fmt.Errorf(`error "GetUser" %v`, err))
	}
	if TestValueCustomers[0].MsgCustomers.FirstName != res.FirstName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValueCustomers[0].MsgCustomers.FirstName, res.FirstName))
	}
	if TestValueCustomers[0].MsgCustomers.LastName != res.LastName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValueCustomers[0].MsgCustomers.LastName, res.LastName))
	}
	if TestValueCustomers[0].MsgCustomers.MiddleName != res.MiddleName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValueCustomers[0].MsgCustomers.MiddleName, res.MiddleName))
	}
	if TestValueCustomers[0].MsgCustomers.MiddleName != res.MiddleName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValueCustomers[0].MsgCustomers.MiddleName, res.MiddleName))
	}
}

// test db account_balance
func accountBalance(t *testing.T, db *postgredb.DB) {

	for _, message := range TestValueAccountBalance {
		if err := db.SaveAccountBalance(message.MsgAccountBalance); err != nil {
			t.Error(fmt.Errorf(`error "SaveAccountBalance" %v`, err))
		}
	}

	a := &business.Accouting{
		DB: db,
	}

	for _, message := range TestMsgValueCashIn {
		if err := a.CashIn(context.Background(), message.MsgValueCashIn); err != nil {
			t.Error(fmt.Errorf(`error "InternalTransfer" %v`, err))
		}
	}

	for _, message := range TestValueInternalTransaction {
		if err := a.InternalTransfer(context.Background(), message.MsgInternalTransaction); err != nil {
			t.Error(fmt.Errorf(`error "InternalTransfer" %v`, err))
		}
	}
}

/*
// test db InternalTransactions

	func internalTransactions(t *testing.T, db *postgredb.DB) {


		for _, message := range TestValueInternalTransaction {
			if err := db.SaveInternalTransaction(message.MsgInternalTransaction); err != nil {
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

		filterInternalTransactions := postgredb.FilterInternalTransaction("amount", "DESC", count, 0)

		res, err := db.GetInternalTrasaction(1)
		if err != nil {
			t.Error(fmt.Errorf(`error "GetInternalTrasaction" %v`, err))
		}

		if TestValueInternalTransaction[0].MsgInternalTransaction.AccountSender != res.AccountSender {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValueInternalTransaction[0].MsgInternalTransaction.AccountSender, res.AccountSender))
		}
		if TestValueInternalTransaction[0].MsgInternalTransaction.AccountRecipient != res.AccountRecipient {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValueInternalTransaction[0].MsgInternalTransaction.AccountRecipient, res.AccountRecipient))
		}
		if TestValueInternalTransaction[0].MsgInternalTransaction.Amount != res.Amount {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValueInternalTransaction[0].MsgInternalTransaction.Amount, res.Amount))
		}
		if TestValueInternalTransaction[0].MsgInternalTransaction.CreatedAt != res.CreatedAt {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValueInternalTransaction[0].MsgInternalTransaction.CreatedAt, res.CreatedAt))
		}

		resData, err := db.ListInternalTransaction(filterInternalTransactions)
		if err != nil {
			t.Error(fmt.Errorf(`error "ListInternalTransaction" %v`, err))
		}

		for key, value := range resData {
			if TestValueInternalTransaction[key].MsgInternalTransaction.AccountSender != value.AccountSender {
				t.Error(fmt.Errorf(`result field %v != %v`, TestValueInternalTransaction[key].MsgInternalTransaction.AccountSender, value.AccountSender))
			}
			if TestValueInternalTransaction[key].MsgInternalTransaction.AccountRecipient != value.AccountRecipient {
				t.Error(fmt.Errorf(`result field %v != %v`, TestValueInternalTransaction[key].MsgInternalTransaction.AccountRecipient, value.AccountRecipient))
			}
			if TestValueInternalTransaction[key].MsgInternalTransaction.Amount != value.Amount {
				t.Error(fmt.Errorf(`result field %v != %v`, TestValueInternalTransaction[key].MsgInternalTransaction.Amount, value.Amount))
			}
			if TestValueInternalTransaction[key].MsgInternalTransaction.CreatedAt != value.CreatedAt {
				t.Error(fmt.Errorf(`result field %v != %v`, TestValueInternalTransaction[key].MsgInternalTransaction.CreatedAt, value.CreatedAt))
			}
		}
	}
*/
var (
	TestValueCustomers = []struct {
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

	TestValueAccountBalance = []struct {
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

	TestMsgValueCashIn = []struct {
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

	TestValueInternalTransaction = []struct {
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
				AccountRecipient: `1`,
				Amount:           `50`,
				CreatedAt:        `2024-10-06T15:38:43+05:00`,
			},
		},
	}
	//countRes = 5
)
