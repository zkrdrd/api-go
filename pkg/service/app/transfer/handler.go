package transfer

import (
	"net/http"
)

// Mux Creator
func (b *Transaction) Handlers() http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc(`/transaction/CacheIn`, service.CallLogic(b.CacheIn))
	// mux.HandleFunc(`/transaction/Transfer`, service.CallLogic(b.Transfer))

	return mux
}
