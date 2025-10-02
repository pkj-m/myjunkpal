package models

import "time"

type User struct {
	ID                string    `json:"id"`
	Email             string    `json:"email"`
	Name              string    `json:"name"`
	Password          string    `json:"password"`
	DailyCalorieGoal  float64   `json:"daily_calorie_goal"`
	DailyProteinGoal  float64   `json:"daily_protein_goal"`
	DailyCarbsGoal    float64   `json:"daily_carbs_goal"`
	DailyFatsGoal     float64   `json:"daily_fats_goal"`
	CreatedAt         time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
