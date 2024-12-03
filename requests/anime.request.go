package requests

import "time"

type AnimeRequest struct {
	Title     string `json:"title" binding:"required"`
	GenreId    int `json:"genreid" binding:"required"`
	Review    string `json:"review" binding:"required"`
	Episodes  int    `json:"episodes" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	UpdatedAt  time.Time `json:"updated_at" binding:"required"`
}