package responses

type ReviewsResponse struct {
	ID      *int    `json:"id"`
	AnimeID *int    `json:"anime_id"`
	UserID  *int    `json:"user_id"`
	Message *string `json:"message"`
}