package models

import "time"

type Food struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"` // Empty for system foods, user_id for custom foods
	Name        string    `json:"name"`
	Calories    float64   `json:"calories"`
	Protein     float64   `json:"protein"`
	Carbs       float64   `json:"carbs"`
	Fats        float64   `json:"fats"`
	ServingSize float64   `json:"serving_size"`
	ServingUnit string    `json:"serving_unit"` // g, ml, cup, etc
	Category    string    `json:"category"`     // fruit, vegetable, protein, etc
	CreatedAt   time.Time `json:"created_at"`
}

type CreateFoodRequest struct {
	Name        string  `json:"name"`
	Calories    float64 `json:"calories"`
	Protein     float64 `json:"protein"`
	Carbs       float64 `json:"carbs"`
	Fats        float64 `json:"fats"`
	ServingSize float64 `json:"serving_size"`
	ServingUnit string  `json:"serving_unit"`
	Category    string  `json:"category"`
}

type UpdateFoodRequest struct {
	Name        string  `json:"name"`
	Calories    float64 `json:"calories"`
	Protein     float64 `json:"protein"`
	Carbs       float64 `json:"carbs"`
	Fats        float64 `json:"fats"`
	ServingSize float64 `json:"serving_size"`
	ServingUnit string  `json:"serving_unit"`
	Category    string  `json:"category"`
}
