package models

type Genre struct {
	ID   *int    `json:"id" gorm:"primaryKey"`
	Name *string `json:"name"`
}