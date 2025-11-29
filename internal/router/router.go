package router

import (
	"net/http"

	"/internal/handlers"
)

func NewRouter(
	publicKeyHandler *handlers.PublicKeyHandler,
	userHandler *handlers.UserHandler,
) *http.ServeMux {

	mux := http.NewServeMux()

	// ------------------------
	// Public Key routes
	// ------------------------
	mux.HandleFunc("POST /keys", publicKeyHandler.CreatePublicKey)
	mux.HandleFunc("GET /keys", publicKeyHandler.GetPublicKeys)
	mux.HandleFunc("GET /keys/user/{id}", publicKeyHandler.GetPublicKeysByUser)
	mux.HandleFunc("DELETE /keys/{id}", publicKeyHandler.DeletePublicKey)

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

	// ------------------------
	// Verification routes
	// ------------------------
	mux.HandleFunc("POST /verify", handlers.VerifySignature)

	return mux
}
