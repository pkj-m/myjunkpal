package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"myjunkpal/models"
	"myjunkpal/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type FoodHandler struct {
	store *storage.JSONStore
}

func NewFoodHandler(store *storage.JSONStore) *FoodHandler {
	return &FoodHandler{store: store}
}

func (h *FoodHandler) GetFoods(w http.ResponseWriter, r *http.Request) {
	var foods []models.Food
	h.store.LoadFromFile("foods.json", &foods)

	// Optional filters
	name := r.URL.Query().Get("name")
	category := r.URL.Query().Get("category")

	// Filter foods
	var filtered []models.Food
	for _, f := range foods {
		// Only show system foods and user's custom foods
		if f.UserID != "" && f.UserID != CurrentUser.ID {
			continue
		}

		if name != "" && !strings.Contains(strings.ToLower(f.Name), strings.ToLower(name)) {
			continue
		}

		if category != "" && !strings.EqualFold(f.Category, category) {
			continue
		}

		filtered = append(filtered, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func (h *FoodHandler) GetFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var foods []models.Food
	h.store.LoadFromFile("foods.json", &foods)

	for _, f := range foods {
		if f.ID == id {
			// Check if user has access
			if f.UserID != "" && f.UserID != CurrentUser.ID {
				http.Error(w, "Food not found", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(f)
			return
		}
	}

	http.Error(w, "Food not found", http.StatusNotFound)
}

func (h *FoodHandler) CreateFood(w http.ResponseWriter, r *http.Request) {
	var req models.CreateFoodRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var foods []models.Food
	h.store.LoadFromFile("foods.json", &foods)

	food := models.Food{
		ID:          uuid.New().String(),
		UserID:      CurrentUser.ID,
		Name:        req.Name,
		Calories:    req.Calories,
		Protein:     req.Protein,
		Carbs:       req.Carbs,
		Fats:        req.Fats,
		ServingSize: req.ServingSize,
		ServingUnit: req.ServingUnit,
		Category:    req.Category,
		CreatedAt:   time.Now(),
	}

	foods = append(foods, food)

	if err := h.store.SaveToFile("foods.json", foods); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(food)
}

func (h *FoodHandler) UpdateFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req models.UpdateFoodRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var foods []models.Food
	h.store.LoadFromFile("foods.json", &foods)

	for i, f := range foods {
		if f.ID == id {
			// Check if user owns this food
			if f.UserID != CurrentUser.ID {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			// Update fields
			foods[i].Name = req.Name
			foods[i].Calories = req.Calories
			foods[i].Protein = req.Protein
			foods[i].Carbs = req.Carbs
			foods[i].Fats = req.Fats
			foods[i].ServingSize = req.ServingSize
			foods[i].ServingUnit = req.ServingUnit
			foods[i].Category = req.Category

			if err := h.store.SaveToFile("foods.json", foods); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(foods[i])
			return
		}
	}

	http.Error(w, "Food not found", http.StatusNotFound)
}

func (h *FoodHandler) DeleteFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var foods []models.Food
	h.store.LoadFromFile("foods.json", &foods)

	for i, f := range foods {
		if f.ID == id {
			// Check if user owns this food
			if f.UserID != CurrentUser.ID {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			// Remove from slice
			foods = append(foods[:i], foods[i+1:]...)

			if err := h.store.SaveToFile("foods.json", foods); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Food not found", http.StatusNotFound)
}
