package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/aswearingen91/account-service/internal/services"
	"github.com/golang-jwt/jwt/v5"
)

// UserHandler handles HTTP endpoints for user-related actions.
type UserHandler struct {
	svc services.UserService
}

// NewUserHandler constructs a new UserHandler.
func NewUserHandler(svc services.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// ------------------------------------------------------------
// CreateUser handles POST /user
// Expects JSON body: {"username":"...", "password":"..."}
// ------------------------------------------------------------
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.svc.CreateUser(body.Username, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

// ------------------------------------------------------------
// GetUser handles GET /user?id=123
// ------------------------------------------------------------
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	id := uint(id64)
	user, err := h.svc.GetUser(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}

// ------------------------------------------------------------
// GetUserByUsername handles GET /user?username=alice
// ------------------------------------------------------------
func (h *UserHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "missing username parameter", http.StatusBadRequest)
		return
	}

	user, err := h.svc.GetUserByUsername(username)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}

// ------------------------------------------------------------
// Login handles POST /user/login
// Returns: { "token": "...", "message": "Logged in" }
// ------------------------------------------------------------
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var jwtSecret = []byte("super-secret-key-change-me")

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, `{"message":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Validate user
	if err := h.svc.Login(body.Username, body.Password); err != nil {
		http.Error(w, `{"message":"login failed"}`, http.StatusUnauthorized)
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": body.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, `{"message":"could not generate token"}`, http.StatusInternalServerError)
		return
	}

	// Return final JSON response
	resp := map[string]string{
		"token":   tokenString,
		"message": "Logged in",
	}

	_ = json.NewEncoder(w).Encode(resp)
}
