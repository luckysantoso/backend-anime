package responses

type AnimeResponse struct {
	ID       *int    `json:"id"`
	Title    *string `json:"title"`
	Genre    *string `json:"genre"`
	Review   *string `json:"review"`
	Episodes *int    `json:"episodes"`
}