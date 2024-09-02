package users

import (
	service "api-go/internal"
	"net/http"
)

// Mux Creator
func (b *Users) Handlers() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(`/users`, service.CallLogic(b.Create))
	//mux.HandleFunc(`/v1/echo`, internal.CallLogic(b.Boo))
	//mux.HandleFunc(`/v1/echo`, internal.CallLogic(b.Foo))

	return mux
}
