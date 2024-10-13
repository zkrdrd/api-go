package transfer

import (
	"net/http"
)

// Mux Creator
func (a *Transfer) Handlers() http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc(`/transaction/CacheIn`, service.CallLogic(a.CashIn))
	// mux.HandleFunc(`/transaction/CacheIn`, service.CallLogic(a.CashOut))
	// mux.HandleFunc(`/transaction/Transfer`, service.CallLogic(a.InternalTransfer))

	return mux
}
