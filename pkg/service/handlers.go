package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zkrdrd/api-go/pkg/models"
)

const (
	MsgInvalidBody = `invalid body`
)

func handlersCashInOrCashOutOrInternalTransfer[modelType *models.CashIn | *models.CashOut | *models.InternalTranser](method string, fn func(context.Context, modelType) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var model modelType

		if err := json.NewDecoder(r.Body).Decode(model); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(MsgInvalidBody))
			return
		}

		if err := fn(r.Context(), model); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// func handlersCashOut(method string, fn func(context.Context, *models.CashOut) error) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if method != r.Method {
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}

// 		cashOut := &models.CashOut{}

// 		if err := json.NewDecoder(r.Body).Decode(cashOut); err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			_, _ = w.Write([]byte(MsgInvalidBody))
// 			return
// 		}

// 		if err := fn(r.Context(), cashOut); err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			_, _ = w.Write([]byte(err.Error()))
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 	}
// }

// func handlersCashIn(method string, fn func(context.Context, *models.CashIn) error) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		if method != r.Method {
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}

// 		cashIn := &models.CashIn{}

// 		if err := json.NewDecoder(r.Body).Decode(cashIn); err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			_, _ = w.Write([]byte(MsgInvalidBody))
// 			return
// 		}

// 		if err := fn(r.Context(), cashIn); err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			_, _ = w.Write([]byte(err.Error()))
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 	}
// }

// func handlersInternalTransfer(method string, fn func(context.Context, *models.InternalTranser) error) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		if method != r.Method {
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}

// 		internalTransfer := &models.InternalTranser{}

// 		if err := json.NewDecoder(r.Body).Decode(internalTransfer); err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			_, _ = w.Write([]byte(MsgInvalidBody))
// 			return
// 		}

// 		if err := fn(r.Context(), internalTransfer); err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			_, _ = w.Write([]byte(err.Error()))
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 	}
// }

// todo
func handlersTransactionsOrGetBalance[modelType *models.Transactions | *models.Balance](method string, fn func(id string) (modelType, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var model modelType

		if err := json.NewDecoder(r.Body).Decode(model); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(MsgInvalidBody))
			return
		}

		// if model == balance; types = model.Account
		// if model == internalTransfer; types = model.AccountSender
		// reflect
		_, err := fn( /*types*/ )
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// // handlersTransactions - метод для чтения транзакций.
// func handlersTransactions(method string, fn func(id string) (*models.Transactions, error)) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if method != r.Method {
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}

// 		internalTransfer := &models.InternalTranser{}

// 		if err := json.NewDecoder(r.Body).Decode(internalTransfer); err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			_, _ = w.Write([]byte(MsgInvalidBody))
// 			return
// 		}

// 		_, err := fn(internalTransfer.AccountSender)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			_, _ = w.Write([]byte(err.Error()))
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 	}
// }

// func handlersGetBalance(method string, fn func(id string) (*models.Balance, error)) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		if method != r.Method {
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}

// 		balanceID := &models.Balance{}

// 		if err := json.NewDecoder(r.Body).Decode(balanceID); err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			_, _ = w.Write([]byte(MsgInvalidBody))
// 			return
// 		}

// 		_, err := fn(balanceID.Account)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			_, _ = w.Write([]byte(err.Error()))
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 	}
// }
