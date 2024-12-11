package responses

type AnimeResponse struct {
	ID       *int          `json:"id"`
	Title    *string       `json:"title"`
	GenreId  *int          `json:"genre"`
	Episodes *int          `json:"episodes"`
	Sinopsis *string       `json:"sinopsis"`
	Genre    GenreResponse `gorm:"foreignKey:GenreId`
}