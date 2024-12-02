package models

import "time"

type Anime struct {
	ID          *int       `json:"id"`
	Title       *string    `json:"title"`
	Genre 	 	*string    `json:"genre"`
	Review		*string    `json:"review"`
	Episodes    *int       `json:"episodes"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}