package service

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CallLogic(fn func(context.Context, []byte) ([]byte, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		buf, _ := io.ReadAll(r.Body)

		// fn for call
		res, err := fn(r.Context(), buf)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print(err)
		}

		data, err := json.Marshal(res)
		w.Write(data)
	}
}
