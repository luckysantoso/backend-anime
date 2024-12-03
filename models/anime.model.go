package models

import "time"

type Anime struct {
    ID        *int       `json:"id" gorm:"primaryKey"`
    Title     *string    `json:"title"`
    GenreId   *int       `json:"genre_id"`
    Review    *string    `json:"review"`
    Episodes  *int       `json:"episodes"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
    Genre     Genre     `gorm:"foreignKey:GenreId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}