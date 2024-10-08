package tests

import (
	"api-go/cmd/app"
	"api-go/internal/postgredb"
	"api-go/pkg/models"
	"fmt"
	"log"
	"testing"
)

func TestDB(t *testing.T) {

	dbconf, _ := app.ParseDBConfig("ConConf.json")
	db, _ := dbconf.NewDB()
	db.RecreateTableInternalTransactions()

	for _, message := range TestValue {
		db.SaveInternalTransaction(message.Msg)
	}

	count, err := db.CountInternalTransactions()
	if err != nil {
		log.Fatal(`error "CountInternalTransactions"`)
	}

	if countRes != count {
		t.Error(fmt.Errorf(`result field %v != %v`, countRes, count))
	}

	filterInternalTransactions := postgredb.FilterInternalTransaction("amount", "DESC", count, 0)

	res, err := db.GetInternalTrasaction(1)
	if err != nil {
		log.Fatal(`error "GetInternalTrasaction"`)
	}

	if TestValue[0].Msg.AccountSender != res.AccountSender {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].Msg.AccountSender, res.AccountSender))
	}
	if TestValue[0].Msg.AccountRecipient != res.AccountRecipient {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].Msg.AccountRecipient, res.AccountRecipient))
	}
	if TestValue[0].Msg.Amount != res.Amount {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].Msg.Amount, res.Amount))
	}
	if TestValue[0].Msg.CreatedAt != res.CreatedAt {
		t.Error(fmt.Errorf(`result field %v != %v`, TestValue[0].Msg.CreatedAt, res.CreatedAt))
	}

	resData, err := db.ListInternalTransaction(filterInternalTransactions)
	if err != nil {
		log.Fatal(`error "ListInternalTransaction"`)
	}

	for key, value := range resData {
		if TestValue[key].Msg.AccountSender != value.AccountSender {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValue[key].Msg.AccountSender, value.AccountSender))
		}
		if TestValue[key].Msg.AccountRecipient != value.AccountRecipient {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValue[key].Msg.AccountRecipient, value.AccountRecipient))
		}
		if TestValue[key].Msg.Amount != value.Amount {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValue[key].Msg.Amount, value.Amount))
		}
		if TestValue[key].Msg.CreatedAt != value.CreatedAt {
			t.Error(fmt.Errorf(`result field %v != %v`, TestValue[key].Msg.CreatedAt, value.CreatedAt))
		}
	}
}

var (
	TestValue = []struct {
		Msg *models.InternalTransaction
	}{
		{
			Msg: &models.InternalTransaction{
				AccountSender:    `1`,
				AccountRecipient: `2`,
				Amount:           `500000`,
				CreatedAt:        `2024-10-06T15:34:43+05:00`,
			},
		},
		{
			Msg: &models.InternalTransaction{
				AccountSender:    `3`,
				AccountRecipient: `4`,
				Amount:           `50000`,
				CreatedAt:        `2024-10-06T15:35:43+05:00`,
			},
		},
		{
			Msg: &models.InternalTransaction{
				AccountSender:    `5`,
				AccountRecipient: `6`,
				Amount:           `5000`,
				CreatedAt:        `2024-10-06T15:36:43+05:00`,
			},
		},
		{
			Msg: &models.InternalTransaction{
				AccountSender:    `7`,
				AccountRecipient: `8`,
				Amount:           `500`,
				CreatedAt:        `2024-10-06T15:37:43+05:00`,
			},
		},
		{
			Msg: &models.InternalTransaction{
				AccountSender:    `9`,
				AccountRecipient: `10`,
				Amount:           `50`,
				CreatedAt:        `2024-10-06T15:38:43+05:00`,
			},
		},
	}
	countRes = 5
)
