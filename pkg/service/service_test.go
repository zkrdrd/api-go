package service_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zkrdrd/api-go/pkg/models"
	"github.com/zkrdrd/api-go/pkg/service"
)

type PrototypeAccounting struct {
}

func (s *PrototypeAccounting) CashIn(ctx context.Context, cacheIn *models.CashIn) error {
	return nil
}
func (s *PrototypeAccounting) CashOut(ctx context.Context, cacheOut *models.CashOut) error {
	return nil
}
func (s *PrototypeAccounting) InternalTransfer(ctx context.Context, transfer *models.InternalTranser) error {
	return nil
}
func (s *PrototypeAccounting) GetTransaction(ctx context.Context, id string) (*models.Transaction, error) {
	return &models.Transaction{
		ID: id,
	}, nil
}

func TestServiceCallHandlers(t *testing.T) {
	service := service.NewService(&PrototypeAccounting{})
	server := httptest.NewServer(service.Handlers())
	defer server.Close()

	client := server.Client()

	t.Run(`test get transaction`, func(t *testing.T) {
		url := server.URL + `/v1/accounting/transactions/1`
		log.Printf(`check url: %s`, url)

		res, err := client.Get(url)
		if err != nil {
			t.Error(err)
			return
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Error(`invalid code`)
			return
		}
	})

	log.Print(`ok`)
}
