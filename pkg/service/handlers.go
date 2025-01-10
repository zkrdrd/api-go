package service

import (
	"encoding/json"
	"net/http"

	"github.com/zkrdrd/api-go/pkg/models"
)

const (
	MsgInvalidBody = `invalid body`
)

func (s *Service) handlersCashIn(w http.ResponseWriter, r *http.Request) {
	var cashIn *models.CashIn

	if err := json.NewDecoder(r.Body).Decode(cashIn); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(MsgInvalidBody))
		return
	}

	if err := s.accounting.CashIn(r.Context(), cashIn); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Service) handlersCashOut(w http.ResponseWriter, r *http.Request) {
	var cashOut *models.CashOut

	if err := json.NewDecoder(r.Body).Decode(cashOut); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(MsgInvalidBody))
		return
	}

	if err := s.accounting.CashOut(r.Context(), cashOut); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Service) handlersGetTransaction(w http.ResponseWriter, r *http.Request) {
	// /source/type/method/arg1/arg2...
	// /source/type/method?arg1=...&arg2=...

	idTrans := r.PathValue(`id`)
	if idTrans == `` {
		_, _ = w.Write([]byte(`id is empty`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := s.accounting.GetTransaction(r.Context(), idTrans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
	w.WriteHeader(http.StatusOK)
}
