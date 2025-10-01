package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"myjunkpal/models"
	"myjunkpal/storage"

	"github.com/gorilla/mux"
)

type NutritionHandler struct {
	store *storage.JSONStore
}

func NewNutritionHandler(store *storage.JSONStore) *NutritionHandler {
	return &NutritionHandler{store: store}
}

func (h *NutritionHandler) GetDailySummary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dateStr := vars["date"]

	// Parse date
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format, use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	// Load entries
	var entries []models.Entry
	h.store.LoadFromFile("entries.json", &entries)

	// Filter entries for this date and user
	var dayEntries []models.Entry
	var totalCalories, totalProtein, totalCarbs, totalFats float64

	for _, e := range entries {
		if e.UserID != CurrentUser.ID {
			continue
		}

		// Check if entry is on the same day
		if e.EatenAt.Year() == date.Year() &&
			e.EatenAt.Month() == date.Month() &&
			e.EatenAt.Day() == date.Day() {
			dayEntries = append(dayEntries, e)
			totalCalories += e.Calories
			totalProtein += e.Protein
			totalCarbs += e.Carbs
			totalFats += e.Fats
		}
	}

	summary := models.NutritionSummary{
		Date:     dateStr,
		Calories: totalCalories,
		Protein:  totalProtein,
		Carbs:    totalCarbs,
		Fats:     totalFats,
		Entries:  dayEntries,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

func (h *NutritionHandler) GetWeeklySummary(w http.ResponseWriter, r *http.Request) {
	// Get start date from query param, default to today
	startDateStr := r.URL.Query().Get("start_date")
	var startDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			http.Error(w, "Invalid start_date format, use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
	} else {
		startDate = time.Now()
	}

	// Calculate 7 days back from start date
	endDate := startDate.Add(-7 * 24 * time.Hour)

	// Load entries
	var entries []models.Entry
	h.store.LoadFromFile("entries.json", &entries)

	// Create a map to store daily summaries
	dailySummaries := make(map[string]*models.NutritionSummary)

	// Initialize 7 days
	for i := 0; i < 7; i++ {
		date := startDate.Add(-time.Duration(i) * 24 * time.Hour)
		dateStr := date.Format("2006-01-02")
		dailySummaries[dateStr] = &models.NutritionSummary{
			Date:    dateStr,
			Entries: []models.Entry{},
		}
	}

	// Aggregate entries
	for _, e := range entries {
		if e.UserID != CurrentUser.ID {
			continue
		}

		if e.EatenAt.After(endDate) && e.EatenAt.Before(startDate.Add(24*time.Hour)) {
			dateStr := e.EatenAt.Format("2006-01-02")
			if summary, exists := dailySummaries[dateStr]; exists {
				summary.Calories += e.Calories
				summary.Protein += e.Protein
				summary.Carbs += e.Carbs
				summary.Fats += e.Fats
				summary.Entries = append(summary.Entries, e)
			}
		}
	}

	// Convert map to slice
	var summaries []models.NutritionSummary
	for _, summary := range dailySummaries {
		summaries = append(summaries, *summary)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summaries)
}

func (h *NutritionHandler) GetGoals(w http.ResponseWriter, r *http.Request) {
	goals := models.NutritionGoals{
		DailyCalorieGoal: CurrentUser.DailyCalorieGoal,
		DailyProteinGoal: CurrentUser.DailyProteinGoal,
		DailyCarbsGoal:   CurrentUser.DailyCarbsGoal,
		DailyFatsGoal:    CurrentUser.DailyFatsGoal,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goals)
}

func (h *NutritionHandler) UpdateGoals(w http.ResponseWriter, r *http.Request) {
	var goals models.NutritionGoals
	if err := json.NewDecoder(r.Body).Decode(&goals); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Load users
	var users []models.User
	h.store.LoadFromFile("users.json", &users)

	// Update current user's goals
	for i, u := range users {
		if u.ID == CurrentUser.ID {
			users[i].DailyCalorieGoal = goals.DailyCalorieGoal
			users[i].DailyProteinGoal = goals.DailyProteinGoal
			users[i].DailyCarbsGoal = goals.DailyCarbsGoal
			users[i].DailyFatsGoal = goals.DailyFatsGoal

			if err := h.store.SaveToFile("users.json", users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Update in-memory user
			CurrentUser.DailyCalorieGoal = goals.DailyCalorieGoal
			CurrentUser.DailyProteinGoal = goals.DailyProteinGoal
			CurrentUser.DailyCarbsGoal = goals.DailyCarbsGoal
			CurrentUser.DailyFatsGoal = goals.DailyFatsGoal

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(goals)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
