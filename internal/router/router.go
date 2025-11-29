package router

import (
	"net/http"

	"github.com/aswearingen91/account-service/internal/handlers"
)

func NewRouter(
	userHandler *handlers.UserHandler,
) *http.ServeMux {

	mux := http.NewServeMux()

	// ------------------------
	// User routes (FIXED)
	// ------------------------
	mux.HandleFunc("POST /user", userHandler.CreateUser)
	mux.HandleFunc("GET /user", func(w http.ResponseWriter, r *http.Request) {
		// The handler decides whether to process id or username
		if r.URL.Query().Has("id") {
			userHandler.GetUser(w, r)
			return
		}
		if r.URL.Query().Has("username") {
			userHandler.GetUserByUsername(w, r)
			return
		}

		http.Error(w, "missing id or username parameter", http.StatusBadRequest)
	})

	return mux
}
