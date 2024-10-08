package tests

import (
	"api-go/cmd/app"
	"api-go/internal/postgredb"
	"api-go/pkg/models"
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {

	dbconf, _ := app.ParseDBConfig("D:\\Programming\\api-go\\ConConf.json")
	db, _ := dbconf.NewDB()

	customers(t, db)

	accountBalance(t, db)

	internalTransactions(t, db)

}

// test db customers
func customers(t *testing.T, db *postgredb.DB) {
	db.RecreateTableCustomers()

	for _, message := range TestValue {
		if err := db.SaveUser(message.MsgCustomers); err != nil {
			t.Error(fmt.Errorf(`error "SaveUser" %v`, err))
		}
	}

	res, err := db.GetUser("1")
	if err != nil {
		t.Error(fmt.Errorf(`error "GetUser" %v`, err))
	}
	if TestValue[0].MsgCustomers.FirstName != res.FirstName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].MsgCustomers.FirstName, res.FirstName))
	}
	if TestValue[0].MsgCustomers.LastName != res.LastName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].MsgCustomers.LastName, res.LastName))
	}
	if TestValue[0].MsgCustomers.MiddleName != res.MiddleName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].MsgCustomers.MiddleName, res.MiddleName))
	}
	if TestValue[0].MsgCustomers.MiddleName != res.MiddleName {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].MsgCustomers.MiddleName, res.MiddleName))
	}
}

// test db account_balance
func accountBalance(t *testing.T, db *postgredb.DB) {
	db.RecreateTableAccountBalance()

	for _, message := range TestValue {
		if err := db.SaveAccountBalance(message.MsgAccountBalance); err != nil {
			t.Error(fmt.Errorf(`error "SaveAccountBalance" %v`, err))
		}
	}

	for _, message := range TestValue {
		casheIn := &models.CacheIn{
			Account: message.MsgInternalTransaction.AccountSender,
			Amount:  message.MsgInternalTransaction.AccountSender,
		}
		casheOut := &models.CacheOut{
			Account: TestValue[0].MsgInternalTransaction.AccountRecipient,
			Amount:  TestValue[0].MsgInternalTransaction.AccountRecipient,
		}

		res, err := db.GetAccountBalance(casheIn.Account)
		if err != nil {
			t.Error(fmt.Errorf(`error "GetAccountBalance" %v`, err))
		}

		// TODO:
		// 1. Добавить обработчик балансов прибавить убавить баланс
		// 2. Проверка на существование пользователя
	}

}

// test db InternalTransactions
func internalTransactions(t *testing.T, db *postgredb.DB) {
	db.RecreateTableInternalTransactions()

	for _, message := range TestValue {
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

	if TestValue[0].MsgInternalTransaction.AccountSender != res.AccountSender {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].MsgInternalTransaction.AccountSender, res.AccountSender))
	}
	if TestValue[0].MsgInternalTransaction.AccountRecipient != res.AccountRecipient {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].MsgInternalTransaction.AccountRecipient, res.AccountRecipient))
	}
	if TestValue[0].MsgInternalTransaction.Amount != res.Amount {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].MsgInternalTransaction.Amount, res.Amount))
	}
	if TestValue[0].MsgInternalTransaction.CreatedAt != res.CreatedAt {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].MsgInternalTransaction.CreatedAt, res.CreatedAt))
	}

	resData, err := db.ListInternalTransaction(filterInternalTransactions)
	if err != nil {
		t.Error(fmt.Errorf(`error "ListInternalTransaction" %v`, err))
	}

	for key, value := range resData {
		if TestValue[key].MsgInternalTransaction.AccountSender != value.AccountSender {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValue[key].MsgInternalTransaction.AccountSender, value.AccountSender))
		}
		if TestValue[key].MsgInternalTransaction.AccountRecipient != value.AccountRecipient {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValue[key].MsgInternalTransaction.AccountRecipient, value.AccountRecipient))
		}
		if TestValue[key].MsgInternalTransaction.Amount != value.Amount {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValue[key].MsgInternalTransaction.Amount, value.Amount))
		}
		if TestValue[key].MsgInternalTransaction.CreatedAt != value.CreatedAt {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValue[key].MsgInternalTransaction.CreatedAt, value.CreatedAt))
		}
	}
}

var (
	TestValue = []struct {
		MsgCustomers           *models.Users
		MsgAccountBalance      *models.CacheIn
		MsgInternalTransaction *models.InternalTransaction
	}{
		{
			MsgCustomers: &models.Users{
				FirstName:  "A",
				LastName:   "A",
				MiddleName: "A",
			},
			MsgAccountBalance: &models.CacheIn{
				Account: "1",
				Amount:  "90000000",
			},
			MsgInternalTransaction: &models.InternalTransaction{
				AccountSender:    `1`,
				AccountRecipient: `2`,
				Amount:           `500000`,
				CreatedAt:        `2024-10-06T15:34:43+05:00`,
			},
		},
		{
			MsgCustomers: &models.Users{
				FirstName:  "B",
				LastName:   "B",
				MiddleName: "B",
			},
			MsgAccountBalance: &models.CacheIn{
				Account: "2",
				Amount:  "9000000",
			},
			MsgInternalTransaction: &models.InternalTransaction{
				AccountSender:    `2`,
				AccountRecipient: `3`,
				Amount:           `50000`,
				CreatedAt:        `2024-10-06T15:35:43+05:00`,
			},
		},
		{
			MsgCustomers: &models.Users{
				FirstName:  "C",
				LastName:   "C",
				MiddleName: "C",
			},
			MsgAccountBalance: &models.CacheIn{
				Account: "3",
				Amount:  "900000",
			},
			MsgInternalTransaction: &models.InternalTransaction{
				AccountSender:    `3`,
				AccountRecipient: `4`,
				Amount:           `5000`,
				CreatedAt:        `2024-10-06T15:36:43+05:00`,
			},
		},
		{
			MsgCustomers: &models.Users{
				FirstName:  "D",
				LastName:   "D",
				MiddleName: "D",
			},
			MsgAccountBalance: &models.CacheIn{
				Account: "4",
				Amount:  "90000",
			},
			MsgInternalTransaction: &models.InternalTransaction{
				AccountSender:    `4`,
				AccountRecipient: `5`,
				Amount:           `500`,
				CreatedAt:        `2024-10-06T15:37:43+05:00`,
			},
		},
		{
			MsgCustomers: &models.Users{
				FirstName:  "E",
				LastName:   "E",
				MiddleName: "E",
			},
			MsgAccountBalance: &models.CacheIn{
				Account: "5",
				Amount:  "9000",
			},
			MsgInternalTransaction: &models.InternalTransaction{
				AccountSender:    `5`,
				AccountRecipient: `1`,
				Amount:           `50`,
				CreatedAt:        `2024-10-06T15:38:43+05:00`,
			},
		},
	}
	countRes = 5
)
