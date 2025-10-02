package models

import "time"

type Entry struct {
	ID       string    `json:"id"`
	UserID   string    `json:"user_id"`
	FoodID   string    `json:"food_id"`
	FoodName string    `json:"food_name"` // Denormalized for easy display
	Quantity float64   `json:"quantity"`  // Number of servings
	MealType string    `json:"meal_type"` // breakfast, lunch, dinner, snack
	EatenAt  time.Time `json:"eaten_at"`
	Calories float64   `json:"calories"` // Calculated: food.calories * quantity
	Protein  float64   `json:"protein"`  // Calculated: food.protein * quantity
	Carbs    float64   `json:"carbs"`    // Calculated: food.carbs * quantity
	Fats     float64   `json:"fats"`     // Calculated: food.fats * quantity
	CreatedAt time.Time `json:"created_at"`
}

type CreateEntryRequest struct {
	FoodID   string  `json:"food_id"`
	Quantity float64 `json:"quantity"`
	MealType string  `json:"meal_type"`
	EatenAt  string  `json:"eaten_at"` // ISO8601 format
}

type UpdateEntryRequest struct {
	Quantity float64 `json:"quantity"`
	MealType string  `json:"meal_type"`
	EatenAt  string  `json:"eaten_at"`
}

type NutritionSummary struct {
	Date     string  `json:"date"`
	Calories float64 `json:"calories"`
	Protein  float64 `json:"protein"`
	Carbs    float64 `json:"carbs"`
	Fats     float64 `json:"fats"`
	Entries  []Entry `json:"entries"`
}

type NutritionGoals struct {
	DailyCalorieGoal float64 `json:"daily_calorie_goal"`
	DailyProteinGoal float64 `json:"daily_protein_goal"`
	DailyCarbsGoal   float64 `json:"daily_carbs_goal"`
	DailyFatsGoal    float64 `json:"daily_fats_goal"`
}
