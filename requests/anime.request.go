package requests

import "time"

type AnimeRequest struct {
	Title     string `json:"title" binding:"required"`
	GenreId    int `json:"genreid" binding:"required"`
	Episodes  int    `json:"episodes" binding:"required"`
	Sinopsis  string `json:"sinopsis" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}