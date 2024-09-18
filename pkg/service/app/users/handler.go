package users

import (
	"api-go/pkg/service"
	"net/http"
)

// Mux Creator
func (b *Users) Handlers() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(`/users`, service.CallLogic(b.Create))

	return mux
}
