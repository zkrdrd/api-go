package service

import (
	"net/http"

	"github.com/zkrdrd/api-go/internal/postgredb"
	"github.com/zkrdrd/api-go/internal/services"
)

func Handlers(accounting *services.Accouting,
	storage postgredb.Storage) *http.ServeMux {
	mux := &http.ServeMux{}

	// http.HandleFunc(`/v1/accounting/balance`, handlersGetAllBalances(http.MethodGet, accounting.CashOut)) - no method ListAccountinBalance
	http.HandleFunc(`/v1/accounting/balances`, handlersGetBalance(http.MethodGet, storage.GetAccountBalance))

	// http.HandleFunc(`/v1/accounting/transaction`, handlersGerAllTransactions(http.MethodGet, accounting.ListInternalTransaction)) - how add Filter
	//
	http.HandleFunc(`/v1/accounting/transactions`, handlersTransactions(http.MethodGet, storage.GetInternalTrasaction))
	http.HandleFunc(`/v1/accounting/transactions/{id}`, handlersTransactions(http.MethodGet, storage.GetInternalTrasaction))
	http.HandleFunc(`/v1/accounting/transaction/cash-out`, handlersCashOut(http.MethodPost, accounting.CashOut))
	http.HandleFunc(`/v1/accounting/transaction/cash-in`, handlersCashIn(http.MethodPost, accounting.CashIn))
	http.HandleFunc(`/v1/accounting/transaction/internal-transfer`, handlersInternalTransfer(http.MethodPost, accounting.InternalTransfer))

	return mux
}
