package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"myjunkpal/models"
	"myjunkpal/storage"

	"github.com/google/uuid"
)

var CurrentUser *models.User

type AuthHandler struct {
	store *storage.JSONStore
}

func NewAuthHandler(store *storage.JSONStore) *AuthHandler {
	return &AuthHandler{store: store}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Load existing users
	var users []models.User
	h.store.LoadFromFile("users.json", &users)

	// Check if user already exists
	for _, u := range users {
		if u.Email == req.Email {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
	}

	// Create new user
	user := models.User{
		ID:               uuid.New().String(),
		Email:            req.Email,
		Name:             req.Name,
		Password:         req.Password,
		DailyCalorieGoal: 2000,
		DailyProteinGoal: 150,
		DailyCarbsGoal:   250,
		DailyFatsGoal:    65,
		CreatedAt:        time.Now(),
	}

	users = append(users, user)

	// Save to file
	if err := h.store.SaveToFile("users.json", users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Auto-login after registration
	CurrentUser = &user

	// Don't send password back
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Load users
	var users []models.User
	h.store.LoadFromFile("users.json", &users)

	// Find user
	for _, u := range users {
		if u.Email == req.Email && u.Password == req.Password {
			CurrentUser = &u
			u.Password = ""
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(u)
			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func (h *AuthHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	if CurrentUser == nil {
		http.Error(w, "Not logged in", http.StatusUnauthorized)
		return
	}

	user := *CurrentUser
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
