package models

type Predictions struct {
	ID        *int     `json:"id"`
	ImagePath *string  `json:"image_path"`
	Label     *string  `json:"label"`
	Score     *float64 `json:"score"`
}