package responses

type AnimeResponse struct {
	ID       *int          `json:"id"`
	Title    *string       `json:"title"`
	GenreId  *int          `json:"genre"`
	Review   *string       `json:"review"`
	Episodes *int          `json:"episodes"`
	Genre    GenreResponse `gorm:"foreignKey:GenreId`
}