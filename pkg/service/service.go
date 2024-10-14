package service

import (
	"api-go/internal/services"
	"api-go/pkg/models"
	"context"
	"encoding/json"
	"net/http"
)

//		 /v1
//		    /accounting
//						 /balances
//GET		               		  		- все балансы
//GET	                          /{id} - баланс по ID

//     					 /transactions
//GET												- все операции
//GET	                          	  /{id} 		- операция по ID

//						 /transaction
//POST                   		  	  /cash-out  	- снятие средств
//POST                   		  	  /cash-in  	- снятие средств
//POST								  /transfer     - перевод

//		    /users

// func CallLogic(fn func(context.Context, []byte) ([]byte, error)) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
//

// 		// fn for call
// 		res, err := fn(r.Context(), buf)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			log.Print(err)
// 		}

// 		w.Write(res)
// 	}
// }

const (
	MsgInvalidBody = `invalid body`
)

func handlerAccouting[
	Type models.CashOut |
		models.CashIn |
		models.InternalTranser](
	method string,
	fn func(context.Context, *Type) error,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerData := new(Type)

		if method != r.Method {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(handlerData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(MsgInvalidBody))
			return
		}

		if err := fn(r.Context(), handlerData); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func Handlers(
	accountingService *services.Accouting,
	// userService *services.User
) *http.ServeMux {
	mux := &http.ServeMux{}

	http.HandleFunc(`/v1/accounting/transaction/cash-out`, handlerAccouting(http.MethodPost, accountingService.CashOut))
	http.HandleFunc(`/v1/accounting/transaction/cash-in`, handlerAccouting(http.MethodPost, accountingService.CashIn))
	http.HandleFunc(`/v1/accounting/transaction/internal-transfer`, handlerAccouting(http.MethodPost, accountingService.InternalTransfer))

	return mux
}
