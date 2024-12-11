package models

type Review struct {
	ID      *int    `json:"id" gorm:"primaryKey"`
	AnimeId *int    `json:"anime_id"`
	UserId  *int    `json:"user_id"`
	Message *string `json:"message"`
}