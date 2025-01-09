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
	http.HandleFunc(`/v1/accounting/balances`, handlersTransactionsOrGetBalance(http.MethodGet, storage.GetAccountBalance))

	// http.HandleFunc(`/v1/accounting/transaction`, handlersGerAllTransactions(http.MethodGet, accounting.ListInternalTransaction)) - how add Filter
	//
	// http.HandleFunc(`/v1/accounting/transactions`, handlersTransactions(http.MethodGet, storage.GetInternalTrasaction))
	http.HandleFunc(`/v1/accounting/transactions/{id}`, handlersTransactionsOrGetBalance(http.MethodGet, storage.GetInternalTrasaction))
	http.HandleFunc(`/v1/accounting/transaction/cash-out`, handlersCashInOrCashOutOrInternalTransfer(http.MethodPost, accounting.CashOut))
	http.HandleFunc(`/v1/accounting/transaction/cash-in`, handlersCashInOrCashOutOrInternalTransfer(http.MethodPost, accounting.CashIn))
	http.HandleFunc(`/v1/accounting/transaction/internal-transfer`, handlersCashInOrCashOutOrInternalTransfer(http.MethodPost, accounting.InternalTransfer))

	return mux
}
