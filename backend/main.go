package main

import (
	"log"
	"net/http"

	"myjunkpal/handlers"
	"myjunkpal/middleware"
	"myjunkpal/storage"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize storage
	store := storage.NewJSONStore("./data")
	if err := store.EnsureDataDir(); err != nil {
		log.Fatal("Failed to create data directory:", err)
	}

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(store)
	foodHandler := handlers.NewFoodHandler(store)
	entryHandler := handlers.NewEntryHandler(store)
	nutritionHandler := handlers.NewNutritionHandler(store)

	// Setup router
	r := mux.NewRouter()

	// Auth routes (no auth required)
	r.HandleFunc("/api/auth/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/api/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/users/me", middleware.RequireAuth(authHandler.GetCurrentUser)).Methods("GET")

	// Food routes (auth required)
	r.HandleFunc("/api/foods", middleware.RequireAuth(foodHandler.GetFoods)).Methods("GET")
	r.HandleFunc("/api/foods/{id}", middleware.RequireAuth(foodHandler.GetFood)).Methods("GET")
	r.HandleFunc("/api/foods", middleware.RequireAuth(foodHandler.CreateFood)).Methods("POST")
	r.HandleFunc("/api/foods/{id}", middleware.RequireAuth(foodHandler.UpdateFood)).Methods("PUT")
	r.HandleFunc("/api/foods/{id}", middleware.RequireAuth(foodHandler.DeleteFood)).Methods("DELETE")

	// Entry routes (auth required)
	r.HandleFunc("/api/entries", middleware.RequireAuth(entryHandler.GetEntries)).Methods("GET")
	r.HandleFunc("/api/entries/{id}", middleware.RequireAuth(entryHandler.GetEntry)).Methods("GET")
	r.HandleFunc("/api/entries", middleware.RequireAuth(entryHandler.CreateEntry)).Methods("POST")
	r.HandleFunc("/api/entries/{id}", middleware.RequireAuth(entryHandler.UpdateEntry)).Methods("PUT")
	r.HandleFunc("/api/entries/{id}", middleware.RequireAuth(entryHandler.DeleteEntry)).Methods("DELETE")

	// Nutrition routes (auth required)
	r.HandleFunc("/api/nutrition/daily/{date}", middleware.RequireAuth(nutritionHandler.GetDailySummary)).Methods("GET")
	r.HandleFunc("/api/nutrition/weekly", middleware.RequireAuth(nutritionHandler.GetWeeklySummary)).Methods("GET")
	r.HandleFunc("/api/nutrition/goals", middleware.RequireAuth(nutritionHandler.GetGoals)).Methods("GET")
	r.HandleFunc("/api/nutrition/goals", middleware.RequireAuth(nutritionHandler.UpdateGoals)).Methods("PUT")

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
