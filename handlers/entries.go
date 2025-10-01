package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"myjunkpal/models"
	"myjunkpal/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type EntryHandler struct {
	store *storage.JSONStore
}

func NewEntryHandler(store *storage.JSONStore) *EntryHandler {
	return &EntryHandler{store: store}
}

func (h *EntryHandler) GetEntries(w http.ResponseWriter, r *http.Request) {
	var entries []models.Entry
	h.store.LoadFromFile("entries.json", &entries)

	// Filter by user
	var userEntries []models.Entry
	for _, e := range entries {
		if e.UserID == CurrentUser.ID {
			userEntries = append(userEntries, e)
		}
	}

	// Optional filters
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	mealType := r.URL.Query().Get("meal_type")

	// Apply filters
	var filtered []models.Entry
	for _, e := range userEntries {
		if startDate != "" {
			start, _ := time.Parse("2006-01-02", startDate)
			if e.EatenAt.Before(start) {
				continue
			}
		}

		if endDate != "" {
			end, _ := time.Parse("2006-01-02", endDate)
			if e.EatenAt.After(end.Add(24 * time.Hour)) {
				continue
			}
		}

		if mealType != "" && e.MealType != mealType {
			continue
		}

		filtered = append(filtered, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func (h *EntryHandler) GetEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var entries []models.Entry
	h.store.LoadFromFile("entries.json", &entries)

	for _, e := range entries {
		if e.ID == id && e.UserID == CurrentUser.ID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(e)
			return
		}
	}

	http.Error(w, "Entry not found", http.StatusNotFound)
}

func (h *EntryHandler) CreateEntry(w http.ResponseWriter, r *http.Request) {
	var req models.CreateEntryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Load food to get nutritional info
	var foods []models.Food
	h.store.LoadFromFile("foods.json", &foods)

	var food *models.Food
	for _, f := range foods {
		if f.ID == req.FoodID {
			food = &f
			break
		}
	}

	if food == nil {
		http.Error(w, "Food not found", http.StatusNotFound)
		return
	}

	// Parse eaten_at time
	eatenAt, err := time.Parse(time.RFC3339, req.EatenAt)
	if err != nil {
		http.Error(w, "Invalid eaten_at format, use ISO8601", http.StatusBadRequest)
		return
	}

	// Load existing entries
	var entries []models.Entry
	h.store.LoadFromFile("entries.json", &entries)

	// Create entry with calculated nutrition
	entry := models.Entry{
		ID:        uuid.New().String(),
		UserID:    CurrentUser.ID,
		FoodID:    req.FoodID,
		FoodName:  food.Name,
		Quantity:  req.Quantity,
		MealType:  req.MealType,
		EatenAt:   eatenAt,
		Calories:  food.Calories * req.Quantity,
		Protein:   food.Protein * req.Quantity,
		Carbs:     food.Carbs * req.Quantity,
		Fats:      food.Fats * req.Quantity,
		CreatedAt: time.Now(),
	}

	entries = append(entries, entry)

	if err := h.store.SaveToFile("entries.json", entries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entry)
}

func (h *EntryHandler) UpdateEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req models.UpdateEntryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var entries []models.Entry
	h.store.LoadFromFile("entries.json", &entries)

	for i, e := range entries {
		if e.ID == id && e.UserID == CurrentUser.ID {
			// Load food to recalculate nutrition
			var foods []models.Food
			h.store.LoadFromFile("foods.json", &foods)

			var food *models.Food
			for _, f := range foods {
				if f.ID == e.FoodID {
					food = &f
					break
				}
			}

			if food == nil {
				http.Error(w, "Food not found", http.StatusNotFound)
				return
			}

			// Parse eaten_at time
			eatenAt, err := time.Parse(time.RFC3339, req.EatenAt)
			if err != nil {
				http.Error(w, "Invalid eaten_at format, use ISO8601", http.StatusBadRequest)
				return
			}

			// Update fields
			entries[i].Quantity = req.Quantity
			entries[i].MealType = req.MealType
			entries[i].EatenAt = eatenAt
			entries[i].Calories = food.Calories * req.Quantity
			entries[i].Protein = food.Protein * req.Quantity
			entries[i].Carbs = food.Carbs * req.Quantity
			entries[i].Fats = food.Fats * req.Quantity

			if err := h.store.SaveToFile("entries.json", entries); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(entries[i])
			return
		}
	}

	http.Error(w, "Entry not found", http.StatusNotFound)
}

func (h *EntryHandler) DeleteEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var entries []models.Entry
	h.store.LoadFromFile("entries.json", &entries)

	for i, e := range entries {
		if e.ID == id && e.UserID == CurrentUser.ID {
			// Remove from slice
			entries = append(entries[:i], entries[i+1:]...)

			if err := h.store.SaveToFile("entries.json", entries); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Entry not found", http.StatusNotFound)
}
