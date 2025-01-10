package service

import (
	"context"
	"net/http"

	"github.com/zkrdrd/api-go/pkg/models"
)

type AccoutingManager interface {
	CashIn(ctx context.Context, cacheIn *models.CashIn) error
	CashOut(ctx context.Context, cacheOut *models.CashOut) error
	InternalTransfer(ctx context.Context, transfer *models.InternalTranser) error
	GetTransaction(ctx context.Context, id string) (*models.Transaction, error)
}

type Service struct {
	accounting AccoutingManager
}

func NewService(accounting AccoutingManager) *Service {
	return &Service{
		accounting: accounting,
	}
}

func callMiddlewareByMethod(method string, fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fn(w, r)
	}
}

func (s *Service) Handlers() *http.ServeMux {
	mux := &http.ServeMux{}

	// http.HandleFunc(`/v1/accounting/balance`, handlersGetAllBalances(http.MethodGet, accounting.CashOut)) - no method ListAccountinBalance
	// http.HandleFunc(`/v1/accounting/balances`, handlersTransactionsOrGetBalance(http.MethodGet, storage.GetAccountBalance))

	// http.HandleFunc(`/v1/accounting/transaction`, handlersGerAllTransactions(http.MethodGet, accounting.ListInternalTransaction)) - how add Filter
	//
	// http.HandleFunc(`/v1/accounting/transactions`, handlersTransactions(http.MethodGet, storage.GetInternalTrasaction))
	mux.HandleFunc(`/v1/accounting/transactions/{id}`, callMiddlewareByMethod(http.MethodGet, s.handlersGetTransaction))
	mux.HandleFunc(`/v1/accounting/transaction/cash-out`, callMiddlewareByMethod(http.MethodPost, s.handlersCashOut))
	mux.HandleFunc(`/v1/accounting/transaction/cash-in`, callMiddlewareByMethod(http.MethodPost, s.handlersCashIn))
	// http.HandleFunc(`/v1/accounting/transaction/internal-transfer`, callMiddlewareByMethod(http.MethodPost, accounting.InternalTransfer))

	return mux
}
