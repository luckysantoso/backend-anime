package requests

type ReviewRequest struct {
	AnimeId int    `json:"anime_id" binding:"required"`
	UserId  int    `json:"user_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}
